package main

import (
	"database/sql"
	"github.com/eoinfarrell/photos.eoinfarrell.dev.api/pkg/db"
	"github.com/eoinfarrell/photos.eoinfarrell.dev.api/pkg/models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Use(cors.Default())

	dbCon := db.InitDb()

	// deleteUnusedImages()

	router.GET("/images", func(c *gin.Context) { renderIndexPage(c, getImages(dbCon, c.Query("tag")), db.GetTagsFromDb(dbCon)) })
	router.GET("/images/:id", func(c *gin.Context) { getImageByID(c, dbCon) })
	router.GET("/tags", func(c *gin.Context) { getTagsHtml(c, dbCon) })

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pictures.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run("0.0.0.0:8080")
}

func renderIndexPage(c *gin.Context, images []models.ImageInfo, tags []models.TagInfo) {
	c.HTML(http.StatusOK, "pictures.tmpl", gin.H{
		"images": images,
		"tags":   tags,
	})
}

func getImagesAsHtml(dbCon *sql.DB, c *gin.Context) {
	var images = db.GetImagesFromDbByTag(dbCon, c.Query("tag"))

	c.HTML(http.StatusOK, "pictures.tmpl", gin.H{
		"images": images,
	})
}

func getImages(dbCon *sql.DB, tag string) []models.ImageInfo {
	if tag != "" {
		return db.GetImagesFromDbByTag(dbCon, tag)
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
