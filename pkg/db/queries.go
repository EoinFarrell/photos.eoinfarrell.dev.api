package db

import (
	"fmt"
	"github.com/eoinfarrell/photos.eoinfarrell.dev.api/pkg/models"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetImagesFromDb(dbCon *sql.DB) []models.ImageInfo {
	rows, err := dbCon.Query("SELECT id, name FROM digikam.Images limit 50")

	if err != nil {
		panic(err)
	}

	return handleImageDbRows(rows)
}

func GetImagesFromDbByIdLimit(dbCon *sql.DB, start int, end int) []models.ImageInfo {
	rows, err := dbCon.Query("SELECT id, name FROM digikam.Images limit 50")

	if err != nil {
		panic(err)
	}

	return handleImageDbRows(rows)
}

func GetImagesFromDbByTag(dbCon *sql.DB, tag string) []models.ImageInfo {
	rows, err := dbCon.Query(fmt.Sprintf("select i.id, i.name from digikam.Tags t join digikam.ImageTags it on t.id = it.tagid join digikam.Images i on it.imageid = i.id where t.name = '%s' limit 8;", tag))

	if err != nil {
		panic(err)
	}

	return handleImageDbRows(rows)
}

func GetImageFromDbById(dbCon *sql.DB, id string) []models.ImageInfo {
	rows, err := dbCon.Query(fmt.Sprintf("select id, name from digikam.Images where id = '%d';", id))

	if err != nil {
		panic(err)
	}

	return handleImageDbRows(rows)
}

func GetTagsFromDb(dbCon *sql.DB) []models.TagInfo {
	rows, err := dbCon.Query("SELECT id, pid, name FROM digikam.Tags where pid = 0")

	if err != nil {
		panic(err)
	}

	return handleTagDbRows(rows)
}
