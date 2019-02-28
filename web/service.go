package web

import (
	"fmt"
	"io/ioutil"
	"ncfu-services/database"
	"ncfu-services/reformatting"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	dbConnection      *gorm.DB
	findRecordsResult []*database.DleDynamicRec
)

// StartService старт сервиса
func StartService() {
	var err error
	dbConnection, err = database.InitDB("root@/dle?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.GET("/", startPage)
	router.Static("/css", "web/static/css")
	router.Static("/js", "web/static/js")
	router.Static("/img", "web/static/img")

	v1 := router.Group("/api/v1/")
	{
		v1.POST("dynamic_content/", restAllDynamicRecords)
		v1.POST("convert_dynamic_content/", convertDynamicContent)
	}

	router.Run()
}

func startPage(ctx *gin.Context) {
	os.MkdirAll("./export/xml", os.ModePerm)
	files, _ := ioutil.ReadDir("./export/xml")
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":           "Добро пожаловать в службу сервисов СКФУ",
			"categories":      database.GetAllDleDynamicRecCategories(dbConnection),
			"export_dir_list": files,
		})
}

func restAllDynamicRecords(ctx *gin.Context) {
	categoryID := ctx.PostForm("category")
	startDate := ctx.PostForm("startDate")
	endDate := ctx.PostForm("endDate")
	switch {
	case categoryID != "" && (startDate == "" && endDate == ""):
		findRecordsResult = database.GetAllDleDynamicRecordsByCategory(dbConnection, categoryID)
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"data":   findRecordsResult,
			})
	case categoryID != "" && (startDate != "" && endDate != ""):
		findRecordsResult = database.GetAllDleDynamicRecordsByCategoryAndDateRange(dbConnection, categoryID, startDate, endDate)
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"data":   findRecordsResult,
			})
	case categoryID == "" && (startDate != "" && endDate != ""):
		findRecordsResult = database.GetAllDleDynamicRecordsByDateRange(dbConnection, startDate, endDate)
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"data":   findRecordsResult,
			})
	default:
		findRecordsResult = database.GetAllDleDynamicRecords(dbConnection)
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"data":   findRecordsResult,
			})

	}

}

func convertDynamicContent(ctx *gin.Context) {
	openCmsMainCatDir := ctx.PostForm("output-main-category")
	openCmsTargetCatDir := ctx.PostForm("output-category")
	openCmsOutputFormat := ctx.PostForm("select-opencms-type")

	var outputFormat string
	switch openCmsOutputFormat {
	case "0":
		outputFormat = "newsblock"
	case "1":
		outputFormat = "personalinfo"
	case "2":
		outputFormat = "announcement"
	default:
		outputFormat = ""
	}

	for _, dynamicRec := range findRecordsResult {
		xmlFile, err := os.Create("./export/xml/" + dynamicRec.AltName + ".xml")
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer xmlFile.Close()
		xmlFile.Write(reformatting.GenerateXMLContent(dynamicRec, openCmsMainCatDir, openCmsTargetCatDir, outputFormat))
	}
}
