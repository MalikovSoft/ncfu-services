package reformatting

// Cdata служебный тип, обрамляющий содержимое в CDATA
type Cdata struct {
	Value string `xml:",cdata"`
}

// LinkToResourceType служебный тип для серилизации ссылки на ресурс
type LinkToResourceType struct {
	Type   string `xml:"type,attr"`
	Target Cdata  `xml:"target"`
}

// ImagePreviewType служебный тип, отражающий поле изображения предварительного просмотра
type ImagePreviewType struct {
	Link LinkToResourceType `xml:"link"`
}

// FullDescriptionType служебный тип, серилизующий полное описание
type FullDescriptionType struct {
	AttrName string `xml:"name,attr"`
	Links    string `xml:"links"`
	Content  Cdata  `xml:"content"`
}

// CategoryType служебный тип, серилизующий категорию
type CategoryType struct {
	Link []LinkToResourceType `xml:"link"`
}

// PhotoType служебный тип, отражающий Фото для галереи изображений
type PhotoType struct {
	Link LinkToResourceType `xml:"link"`
}

// ImagesType служебный тип, серилизующий галерею изображений
type ImagesType struct {
	Photo []PhotoType `xml:"Photo"`
}

// PhotographyType служебный тип, отражающий Фото сотрудника
type PhotographyType struct {
	Link LinkToResourceType `xml:"link"`
}

// ContactsType служебный тип, отражающий контактную информацию
type ContactsType struct {
	Address     Cdata `xml:"Address"`
	PhoneNumber Cdata `xml:"PhoneNumber"`
}

// MainInfoType служебный тип, отражающий основную информацию
type MainInfoType struct {
	AttrName string `xml:"name,attr"`
	Links    string `xml:"links"`
	Content  Cdata  `xml:"content"`
}

// OtherInfoType служебный тип, отражающий дополнительную информацию
type OtherInfoType struct {
	AttrName string `xml:"name,attr"`
	Links    string `xml:"links"`
	Content  Cdata  `xml:"content"`
}
