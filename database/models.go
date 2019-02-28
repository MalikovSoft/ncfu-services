package database

import "time"

// DleDynamicRec запись из таблицы динамического контента DLE
type DleDynamicRec struct {
	ID              uint       `json:"id"`
	DateOfRelease   *time.Time `gorm:"column:date;type:datetime;NOT NULL" json:"date"`
	Title           string     `gorm:"column:title;type:varchar;size:255;NOT NULL" json:"title"`
	Description     string     `gorm:"column:short_story;type:text;NOT NULL" json:"short_story"`
	DescriptionText string     `gorm:"column:descr;type:type:varchar;size:200;NOT NULL" json:"descr"`
	Content         string     `gorm:"column:full_story;type:text;NOT NULL" json:"full_story"`
	Category        string     `gorm:"column:category;type:varchar;size:200;NOT NULL" json:"category"`
	AllowMainFlag   bool       `gorm:"column:allow_main;type:tinyint;size:1;NOT NULL" json:"allow_main"`
	AltName         string     `gorm:"column:alt_name;type:varchar;size:200;NOT NULL" json:"alt_name"`
	XFields         string     `gorm:"column:xfields;type:text;NOT NULL" json:"xfields"`
}

// DleDynamicRecCategory запись из таблицы категорий динамического контента DLE
type DleDynamicRecCategory struct {
	ID   uint8  `json:"id"`
	Name string `gorm:"column:name;type:varchar;size:50;NOT NULL" json:"name"`
}

// DleStaticRec запись из таблицы статического контента DLE
type DleStaticRec struct {
	ID            uint   `json:"id"`
	Name          string `gorm:"column:name;type:varchar;size:100;NOT NULL" json:"name"`
	Description   string `gorm:"column:descr;type:varchar;size:255;NOT NULL" json:"descr"`
	Content       string `gorm:"column:template;type:longtext;NOT NULL" json:"template"`
	DateOfRelease uint   `gorm:"column:date;type:varchar;size:15;NOT NULL" json:"date"`
}

// TableName возвращает имя таблицы динамического контента в БД DLE
func (DleDynamicRec) TableName() string {
	return "dle.dle_post"
}

// TableName возвращает имя таблицы категорий динамического контента в БД DLE
func (DleDynamicRecCategory) TableName() string {
	return "dle.dle_category"
}

// TableName возвращает имя таблицы статического контента в БД DLE
func (DleStaticRec) TableName() string {
	return "dle.dle_static"
}
