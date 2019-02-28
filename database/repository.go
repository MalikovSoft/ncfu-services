package database

import (
	"github.com/jinzhu/gorm"
)

// GetAllDleDynamicRecCategories возвращает все записи из таблицы категорий динамического контента DLE
func GetAllDleDynamicRecCategories(db *gorm.DB) []*DleDynamicRecCategory {
	dleCategories := make([]*DleDynamicRecCategory, 0)
	db.AutoMigrate(&DleDynamicRecCategory{})
	db.Find(&dleCategories)
	return dleCategories
}

// GetAllDleDynamicRecords возвращает все записи из таблицы динамического контента DLE
func GetAllDleDynamicRecords(db *gorm.DB) []*DleDynamicRec {
	dleDynamicRecords := make([]*DleDynamicRec, 0)
	db.AutoMigrate(&DleDynamicRec{})
	db.Find(&dleDynamicRecords)
	return dleDynamicRecords
}

// GetAllDleDynamicRecordsByCategory возвращает все записи, заданной категории, из таблицы динамического контента DLE
func GetAllDleDynamicRecordsByCategory(db *gorm.DB, categoryID string) []*DleDynamicRec {
	dleDynamicRecords := make([]*DleDynamicRec, 0)
	regexp := "(^" + categoryID + "$|," + categoryID + ",|," + categoryID + "$)"
	db.AutoMigrate(&DleDynamicRec{})
	db.Where("category RLIKE ?", regexp).Find(&dleDynamicRecords)
	return dleDynamicRecords
}

// GetAllDleDynamicRecordsByDateRange возвращает все записи, заданного дипазона дат, из таблицы динамического контента DLE
func GetAllDleDynamicRecordsByDateRange(db *gorm.DB, startDate string, endDate string) []*DleDynamicRec {
	dleDynamicRecords := make([]*DleDynamicRec, 0)
	db.AutoMigrate(&DleDynamicRec{})
	db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&dleDynamicRecords)
	return dleDynamicRecords
}

// GetAllDleDynamicRecordsByCategoryAndDateRange возвращает все записи, заданной категории и дипазона дат, из таблицы динамического контента DLE
func GetAllDleDynamicRecordsByCategoryAndDateRange(db *gorm.DB, categoryID string, startDate string, endDate string) []*DleDynamicRec {
	dleDynamicRecords := make([]*DleDynamicRec, 0)
	regexp := "(^" + categoryID + "$|," + categoryID + ",|," + categoryID + "$)"
	db.AutoMigrate(&DleDynamicRec{})
	db.Where("(category RLIKE ?) AND (date BETWEEN ? AND ?)", regexp, startDate, endDate).Find(&dleDynamicRecords)
	return dleDynamicRecords
}

// GetAllDleStaticRecords аозвращает все записи из таблицы статического контента DLE
func GetAllDleStaticRecords(db *gorm.DB) []*DleStaticRec {
	dleStaticRecords := make([]*DleStaticRec, 0)
	db.AutoMigrate(&DleStaticRec{})
	db.Find(&dleStaticRecords)
	return dleStaticRecords
}
