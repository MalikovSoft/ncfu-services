package reformatting

import (
	"encoding/xml"
	"fmt"
	"ncfu-services/database"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/microcosm-cc/bluemonday"
)

// GenerateXMLContent генерирует XML-контенет для последующей загрузки в OpenCMS
func GenerateXMLContent(dbRecord *database.DleDynamicRec, homeCatDir, targetCatDir, elementType string) []byte {
	result := make([]byte, 0)
	switch elementType {
	case "newsblock":
		outXMLContent := &OpenCmsNewsBlocks{
			XMLAttr:           "http://www.w3.org/2001/XMLSchema-instance",
			XMLSchemaLocation: "opencms://system/modules/ru.soft.malikov.web/schemas/NewsBlock.xsd",
		}
		innerContent := &outXMLContent.NewsBlock
		innerContent.Lang = "ru"
		innerContent.Title.Value = dbRecord.Title
		innerContent.ShortDescription.Value = removeAllTagsFromContent(removeAllImgTagsFromContent(dbRecord.Description))
		innerContent.ImagePreview.Link.Type = "WEAK"
		innerContent.ImagePreview.Link.Target.Value = getFirstImgLink(strings.Replace(dbRecord.Description, "\\\"", "\"", -1))
		innerContent.Date = strconv.FormatInt(unixMilli(*dbRecord.DateOfRelease), 10)
		innerContent.FullDescription.AttrName = "FullDescription0"
		content := strings.Replace(dbRecord.Content, "\\\"", "\"", -1)
		innerContent.FullDescription.Content.Value = removeAllImgTagsFromContent(content)
		linksToCategories := make([]LinkToResourceType, 0)
		linksToCategories = append(linksToCategories, LinkToResourceType{
			Type: "WEAK",
			Target: Cdata{
				Value: targetCatDir,
			},
		})
		if dbRecord.AllowMainFlag {
			linksToCategories = append(linksToCategories, LinkToResourceType{
				Type: "WEAK",
				Target: Cdata{
					Value: homeCatDir,
				},
			})
		}
		innerContent.Category.Link = linksToCategories
		images := make([]PhotoType, 0)
		for _, img := range getValidLinks(getImgsLinks(strings.Replace(dbRecord.Content, "\\\"", "\"", -1))) {
			images = append(images, PhotoType{
				Link: LinkToResourceType{
					Type: "WEAK",
					Target: Cdata{
						Value: img,
					},
				},
			})
		}
		innerContent.Images.Photo = images
		innerContent.Counter.Value = "0"
		out, err := xml.MarshalIndent(outXMLContent, " ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		result = []byte(xml.Header + string(out))
	}
	return result
}

// getFirstImgLink возвращает первую ссылку на картинку (обычно используется для парсинга адреса картинки предварительного просмотра)
func getFirstImgLink(content string) string {
	result := ""
	if len(getImgsLinks(content)) > 0 {
		imgLink := getImgsLinks(content)[0]

		err := DownloadFile("./export/files/imported-from-dle/prev", imgLink)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		result = "/sites/ncfu/uploads/imported-from-dle/prev" + imgLink[strings.LastIndex(imgLink, "/"):len(imgLink)]
	}

	return result
}

// getImgsLinks возвращает слайс ссылок на картинки в контенте
func getImgsLinks(content string) []string {
	formattedContent := getFormattedContent(content)
	links := make([]string, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(formattedContent))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	doc.Find("img").Each(func(loop int, s *goquery.Selection) {

		if s.ParentsFiltered(`a`).Length() > 0 {
			s.ParentFiltered(`a`).Each(func(index int, parent *goquery.Selection) {
				href, _ := parent.Attr("href")
				links = append(links, href)
			})
		} else {
			href, _ := s.Attr("src")
			links = append(links, href)
		}
	})

	/*
		doc.Find("a").Each(func(loop int, s *goquery.Selection) {
			if s.Find("img").Length() > 0 {
				href, _ := s.Attr("href")
				links = append(links, href)
			}
		})
	*/

	return links
}

// removeAllImgTagsFromContent удаляет все изображения(включая ссылки-обертки) из контента
func removeAllImgTagsFromContent(content string) string {
	serviceTagsRegexp := regexp.MustCompile(`(?m)<!--TBegin.*?-->(.*?)<!--TEnd-->`)
	content = serviceTagsRegexp.ReplaceAllString(content, "")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	doc.Find(`div:empty`).Each(func(loop int, s *goquery.Selection) {
		s.Remove()
	})
	doc.Find(`span:empty`).Each(func(loop int, s *goquery.Selection) {
		s.Remove()
	})
	doc.Find(`p:empty`).Each(func(loop int, s *goquery.Selection) {
		s.Remove()
	})
	doc.Find(`img`).Each(func(loop int, s *goquery.Selection) {
		s.Remove()
	})
	/*
		doc.Find(`p`).Each(func(loop int, s *goquery.Selection) {
			str, _ := s.Html()
			reg := regexp.MustCompile(`\s|&nbsp;|\s\n|[ ]`)
			if len(reg.ReplaceAllString(str, ``)) == 0 {
				fmt.Println("remove: ", str)
				//s.Remove()
			}

		})
	*/
	doc.Find(`:empty`).Each(func(loop int, s *goquery.Selection) {
		s.Remove()
	})

	doc.Find(`a`).Each(func (loop int, s *goquery.Selection)  {
		
	})

	result, err := doc.Html()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return rmvEmptyTags(getFormattedContent(result))
}

// getFormattedContent возвращает форматированный контент
func getFormattedContent(content string) string {
	htmlPolicy := bluemonday.UGCPolicy()
	htmlPolicy.AllowAttrs("style").OnElements("p", "span")
	return htmlPolicy.Sanitize(content)
}

// getValidLinks преобразует пути ссылок на соответствующие новому расположению
func getValidLinks(links []string) []string {
	result := make([]string, 0)
	for _, link := range links {
		if isFileLink(link) {
			tmp := link[strings.LastIndex(link, "/"):len(link)]
			if isImgFile(tmp) {
				tmp = "/sites/ncfu/uploads/imported-from-dle/img" + tmp
				err := DownloadFile("./export/files/imported-from-dle/img", link)
				if err != nil {
					fmt.Printf("error: %v\n", err)
				}
			} else {
				tmp = "/sites/ncfu/uploads/imported-from-dle/doc" + tmp
				err := DownloadFile("./export/files/imported-from-dle/doc", link)
				if err != nil {
					fmt.Printf("error: %v\n", err)
				}
			}
			result = append(result, tmp)
		}
	}
	return result
}

// isFileLink проверяет является ли url ссылкой на файл или на страницу
func isFileLink(link string) bool {
	tmpURL, err := url.Parse(link)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return false
	}

	var validURL string
	if validURL = tmpURL.String(); !tmpURL.IsAbs() {
		validURL = "http://www.ncfu.ru" + tmpURL.String()
	}

	resp, err := http.Get(validURL)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return false
	}

	if resp.Header.Get("content-type") == "text/html" {
		return false
	}
	return true
}

// isImgFile определяет является ли файл изображением
func isImgFile(link string) bool {
	types := [...]string{
		".jpg",
		".jpeg",
		".png",
		".gif",
		".bmp",
		".psd",
		".ico",
	}
	for _, imgType := range types {
		if strings.HasSuffix(link, imgType) {
			return true
		}
	}
	return false
}

// removeAllTagsFromContent возвращает текст из контента(без html-тегов)
func removeAllTagsFromContent(content string) string {
	htmlPolicy := bluemonday.StrictPolicy()
	return htmlPolicy.Sanitize(content)
}

func rmvEmptyTags(content string) string {
	re := regexp.MustCompile(`<p>([\r\n\t\f\v]|&nbsp;|[\r\n\t\f\v]$|[ ])</p>`)
	result := re.ReplaceAllLiteralString(content, "")
	return result
}
