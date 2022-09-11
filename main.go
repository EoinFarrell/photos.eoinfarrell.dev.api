package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type imageInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var dbCon *sql.DB

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	initDb()

	router.GET("/images", getImages)
	router.GET("/images/:id", getImageByID)

	router.Run("0.0.0.0:8080")
}

func initDb() {
	host := os.Getenv("DIGIKAM_DB_HOST")
	port := os.Getenv("DIGIKAM_DB_PORT")
	username := os.Getenv("DIGIKAM_DB_USER")
	password := os.Getenv("DIGIKAM_DB_USER_PASSWORD")
	dbName := os.Getenv("DIGIKAM_DB_NAME")

	log.Println("asdf-1")
	log.Println(host)
	log.Println(port)
	log.Println(username)
	log.Println(dbName)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	log.Println("asdf-2")

	var err error

	// create a database object which can be used
	// to connect with database.
	dbCon, err = sql.Open("mysql", connectionString)

	log.Println("asdf-3")

	// handle error, if any.
	if err != nil {
		log.Println("asdf-4")
		panic(err)
	}

	log.Println("asdf-5")

	testDb()
}

func testDb() {
	log.Println("asdf-6")
	err := dbCon.Ping()
	if err != nil {
		log.Println("asdf-7")
		panic(err)
	}

	log.Println("asdf-8")

	// Here a SQL query is used to return all
	// the data from the table user.
	rows, err := dbCon.Query("SELECT id, name FROM digikam.Images limit 3")
	log.Println("asdf-9")

	if err != nil {
		log.Println("asdf-10")
		panic(err)
	}

	log.Println("asdf-11")

	// the result object has a method called Next,
	// which is used to iterate through all returned rows.
	for _, image := range handleImageDbRows(rows) {
		fmt.Printf("Id: %d Name: %s\n", image.ID, image.Name)
	}
}

func getImagesFromDb() []imageInfo {
	rows, err := dbCon.Query("SELECT id, name FROM digikam.Images limit 3")

	if err != nil {
		panic(err)
	}

	return handleImageDbRows(rows)
}

func getImagesFromDbByTag(tag string) []imageInfo {
	rows, err := dbCon.Query(fmt.Sprintf("select i.id, i.name from digikam.Tags t join digikam.ImageTags it on t.id = it.tagid join digikam.Images i on it.imageid = i.id where t.name = '%s' limit 3;", tag))

	if err != nil {
		panic(err)
	}

	return handleImageDbRows(rows)
}

func getImageFromDbById(id string) []imageInfo {
	rows, err := dbCon.Query(fmt.Sprintf("select id, name from digikam.Images where id = '%s';", id))

	if err != nil {
		panic(err)
	}

	return handleImageDbRows(rows)
}

func handleImageDbRows(rows *sql.Rows) []imageInfo {
	var image imageInfo
	var images []imageInfo

	// the result object has a method called Next,
	// which is used to iterate through all returned rows.
	for rows.Next() {
		err := rows.Scan(&image.ID, &image.Name)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(image.ID, image.Name)
		images = append(images, image)
	}

	return images
}

// getAlbums responds with the list of all albums as JSON.
func getImages(c *gin.Context) {
	tag := c.Query("tag")

	if tag != "" {
		c.IndentedJSON(http.StatusOK, getImagesFromDbByTag(tag))
	} else {
		c.IndentedJSON(http.StatusOK, getImagesFromDb())
	}
}

// getAlbums responds with the list of all albums as JSON.
func getImageByID(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getImageFromDbById(c.Param("id")))
}
