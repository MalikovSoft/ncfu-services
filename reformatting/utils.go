package reformatting

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

// DownloadFile загружает файл по ссылке в указанную директорию
func DownloadFile(targetDir, link string) error {

	tmpURL, err := url.Parse(link)
	if err != nil {
		return err
	}

	var validURL string
	if validURL = tmpURL.String(); !tmpURL.IsAbs() {
		validURL = "http://www.ncfu.ru" + tmpURL.String()
	}

	out, err := os.Create(targetDir + validURL[strings.LastIndex(validURL, "/"):len(validURL)])
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(validURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// unixMilli преобразование полной даты в миллисекунды (timestamp)
func unixMilli(t time.Time) int64 {
	return t.UnixNano() / 1000000
	//return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func deleteEmptyHTMLTags(content string) string {
	emptyTagsRegexp := regexp.MustCompile(`(?m)<[div>]*?>\s*<\/[div>]*?>`)
	//emptyTagsRegexp := regexp.MustCompile(`(?m)<[^>]*?>\s*<\/[^>]*?>`)
	result := emptyTagsRegexp.ReplaceAllString(content, "")
	return result
}
