package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type imageInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()

	dbTest()

	router.GET("/images", getImages)
	router.GET("/images/:id", getImageByID)

	router.Run("localhost:8080")
}

func getDbConnection() *sql.DB {
	host := os.Getenv("DIGIKAM_DB_HOST")
	port := os.Getenv("DIGIKAM_DB_PORT")
	username := os.Getenv("DIGIKAM_DB_USER")
	password := os.Getenv("DIGIKAM_DB_USER_PASSWORD")
	dbName := os.Getenv("DIGIKAM_DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	// create a database object which can be used
	// to connect with database.
	db, err := sql.Open("mysql", connectionString)

	// handle error, if any.
	if err != nil {
		panic(err)
	}

	return db
}

func dbTest() {
	db := getDbConnection()

	// Here a SQL query is used to return all
	// the data from the table user.
	result, err := db.Query("SELECT id, name FROM digikam.Images limit 3")

	// handle error
	if err != nil {
		panic(err)
	}

	// the result object has a method called Next,
	// which is used to iterate through all returned rows.
	for result.Next() {

		var id int
		var name string

		// The result object provided Scan  method
		// to read row data, Scan returns error,
		// if any. Here we read id and name returned.
		err = result.Scan(&id, &name)

		// handle error
		if err != nil {
			panic(err)
		}

		fmt.Printf("Id: %d Name: %s\n", id, name)
	}

	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
}

func getImagesFromDb() []imageInfo {
	db := getDbConnection()

	// Here a SQL query is used to return all
	// the data from the table user.
	rows, err := db.Query("SELECT id, name FROM digikam.Images limit 3")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	return handleImageDbRows(rows)
}

func getImagesFromDbByTag(tag string) []imageInfo {
	db := getDbConnection()

	// Here a SQL query is used to return all
	// the data from the table user.
	rows, err := db.Query(fmt.Sprintf("select i.id, i.name from digikam.Tags t join digikam.ImageTags it on t.id = it.tagid join digikam.Images i on it.imageid = i.id where t.name = '%s' limit 3;", tag))
	defer db.Close()

	if err != nil {
		panic(err)
	}

	return handleImageDbRows(rows)
}

func getImageFromDbById(id string) []imageInfo {
	db := getDbConnection()

	// Here a SQL query is used to return all
	// the data from the table user.
	rows, err := db.Query(fmt.Sprintf("select id, name from digikam.Images where id = '%s';", id))
	defer db.Close()

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
