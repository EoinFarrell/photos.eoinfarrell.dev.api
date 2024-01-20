package main

import (
	"database/sql"
	"github.com/eoinfarrell/photos.eoinfarrell.dev.api/pkg/db"
	"github.com/eoinfarrell/photos.eoinfarrell.dev.api/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Use(cors.Default())

	dbCon := db.InitDb()

	router.GET("/images", func(c *gin.Context) { renderIndexPage(c, getImages(dbCon, c.Query("tag")), db.GetTagsFromDb(dbCon)) })
	router.GET("/images/:id", func(c *gin.Context) { getImageByID(c, dbCon) })
	router.GET("/tags", func(c *gin.Context) { getTagsHtml(c, dbCon) })

	router.GET("/index", func(c *gin.Context) { getIndexPage(c, dbCon) })
	router.GET("/collection", func(c *gin.Context) { getCollectionPage(c, dbCon, c.Query("id")) })

	router.Run("0.0.0.0:8080")
}

func getIndexPage(c *gin.Context, dbCon *sql.DB) {
	imageAndTag := db.GetRandomImageForAllTopLevelTags(dbCon)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"imageAndTag": imageAndTag,
	})
}

func getCollectionPage(c *gin.Context, dbCon *sql.DB, col string) {
	imageAndTag := db.GetRandomImageForAllTagsInCollection(dbCon, col)

	if imageAndTag == nil {
		if colId, err := strconv.Atoi(col); err == nil {
			images := db.GetImagesFromDbByTagId(dbCon, colId)

			c.HTML(http.StatusOK, "collection.tmpl", gin.H{
				"images": images,
			})
		}
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"imageAndTag": imageAndTag,
		})
	}
}

func renderIndexPage(c *gin.Context, images []models.ImageInfo, tags []models.TagInfo) {
	c.HTML(http.StatusOK, "pictures.tmpl", gin.H{
		"images": images,
		"tags":   tags,
	})
}

func getImages(dbCon *sql.DB, tag string) []models.ImageInfo {
	if tag != "" {
		return db.GetImagesFromDbByTagName(dbCon, tag)
	} else {
		return db.GetImagesFromDb(dbCon)
	}
}

// getImagesAsJson responds with the list of all albums as JSON.
func getImagesAsJson(c *gin.Context, dbCon *sql.DB) {
	c.IndentedJSON(http.StatusOK, getImages(dbCon, c.Query("tag")))
}

// getAlbums responds with the list of all albums as JSON.
func getImageByID(c *gin.Context, dbCon *sql.DB) {
	c.IndentedJSON(http.StatusOK, db.GetImageFromDbById(dbCon, c.Param("id")))
}

func getTagsHtml(c *gin.Context, dbCon *sql.DB) {
	var tags = db.GetTagsFromDb(dbCon)

	c.HTML(http.StatusOK, "pictures.tmpl", gin.H{
		"images": tags,
	})
}

func getTags(c *gin.Context, dbCon *sql.DB) {
	c.IndentedJSON(http.StatusOK, db.GetTagsFromDb(dbCon))
}
