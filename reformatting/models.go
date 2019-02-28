package reformatting

import (
	"encoding/xml"
)

// OpenCmsNewsBlocks тип, создающий структуру типа Новость XML-файла, для загрузки в OpenCMS
type OpenCmsNewsBlocks struct {
	XMLName           xml.Name         `xml:"NewsBlocks"`
	XMLAttr           string           `xml:"xmlns:xsi,attr"`
	XMLSchemaLocation string           `xml:"xsi:noNamespaceSchemaLocation,attr"`
	NewsBlock         OpenCmsNewsBlock `xml:"NewsBlock"`
}

// OpenCmsNewsBlock внутренний тип структуры OpenCmsNewsBlocks
type OpenCmsNewsBlock struct {
	Lang             string              `xml:"language,attr"`
	Title            Cdata               `xml:"Title"`
	ShortDescription Cdata               `xml:"ShortDescription"`
	ImagePreview     ImagePreviewType    `xml:"ImagePreview"`
	Date             string              `xml:"Date"`
	FullDescription  FullDescriptionType `xml:"FullDescription"`
	Category         CategoryType        `xml:"Category"`
	Images           ImagesType          `xml:"Images"`
	Counter          Cdata               `xml:"Counter"`
}

// OpenCmsPersonalInfos тип, создающий структуру типа Сотрудник XML-файла, для загрузки в OpenCMS
type OpenCmsPersonalInfos struct {
	XMLName           xml.Name            `xml:"PersonalInfos"`
	XMLAttr           string              `xml:"xmlns:xsi,attr"`
	XMLSchemaLocation string              `xml:"xsi:noNamespaceSchemaLocation,attr"`
	PersonalInfo      OpenCmsPersonalInfo `xml:"PersonalInfo"`
}

// OpenCmsPersonalInfo внутренний тип структуры OpenCmsPersonalInfos
type OpenCmsPersonalInfo struct {
	Lang        string          `xml:"language,attr"`
	Surname     Cdata           `xml:"Surname"`
	FirstName   Cdata           `xml:"FirstName"`
	Patronimic  Cdata           `xml:"Patronimic"`
	Photography PhotographyType `xml:"Photography"`
	Contacts    ContactsType    `xml:"Contacts"`
	Position    Cdata           `xml:"Position"`
	Institute   Cdata           `xml:"Institute"`
	Department  Cdata           `xml:"Department"`
	MainInfo    MainInfoType    `xml:"MainInfo"`
	OtherInfo   OtherInfoType   `xml:"OtherInfo"`
}

// OpenCmsAnnouncements тип, создающий структуру типа Анонс XML-файла, для загрузки в OpenCMS
type OpenCmsAnnouncements struct {
	XMLName           xml.Name            `xml:"Announcements"`
	XMLAttr           string              `xml:"xmlns:xsi,attr"`
	XMLSchemaLocation string              `xml:"xsi:noNamespaceSchemaLocation,attr"`
	Announcement      OpenCmsAnnouncement `xml:"Announcement"`
}

// OpenCmsAnnouncement внутренний тип структуры OpenCmsAnnouncements
type OpenCmsAnnouncement struct {
	Lang             string              `xml:"language,attr"`
	Title            Cdata               `xml:"Title"`
	ShortDescription Cdata               `xml:"ShortDescription"`
	ImagePreview     ImagePreviewType    `xml:"ImagePreview"`
	Date             string              `xml:"Date"`
	FullDescription  FullDescriptionType `xml:"FullDescription"`
}
