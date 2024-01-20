package db

import (
	"fmt"
	"github.com/eoinfarrell/photos.eoinfarrell.dev.api/pkg/models"
	"log"

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

func GetImagesFromDbByTagName(dbCon *sql.DB, tag string) []models.ImageInfo {
	rows, err := dbCon.Query(fmt.Sprintf("select i.id, i.name from digikam.Tags t join digikam.ImageTags it on t.id = it.tagid join digikam.Images i on it.imageid = i.id where t.name = '%s' limit 8;", tag))

	if err != nil {
		panic(err)
	}

	return handleImageDbRows(rows)
}

func GetImagesFromDbByTagId(dbCon *sql.DB, tag int) []models.ImageInfo {
	rows, err := dbCon.Query(fmt.Sprintf("select i.id, i.name from digikam.Tags t join digikam.ImageTags it on t.id = it.tagid join digikam.Images i on it.imageid = i.id where t.id = '%d';", tag))

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

func GetRandomImageForAllTopLevelTags(dbCon *sql.DB) []models.ImageAndTag {
	query := "SELECT\n    CASE WHEN Images.id IS NULL THEN 0 ELSE Images.id END,\n    CASE WHEN Images.Name IS NULL THEN '' ELSE Images.Name END,\n    Tags.id,\n    Tags.pid,\n    Tags.name,\n    CASE WHEN t.tagCount IS NULL THEN 0 ELSE t.tagCount END\nFROM (\n  SELECT\n    id AS tagsId,\n    (\n        SELECT imageid FROM ImageTags WHERE ImageTags.tagid = tagsId ORDER BY RAND() LIMIT 1\n    ) AS imageId,\n    (\n        SELECT count(imageid) as count FROM ImageTags WHERE ImageTags.tagid = tagsId GROUP BY tagid\n    ) AS tagCount\n  FROM Tags\n) t\nJOIN Tags on t.tagsId = Tags.id\nLEFT JOIN Images on t.imageId = Images.id\nWHERE Tags.pid = 0\nORDER BY t.tagCount desc;"
	log.Println(query)
	rows, err := dbCon.Query(query)

	if err != nil {
		panic(err)
	}

	return handleImageAndTagDbRows(rows)
}

func GetRandomImageForAllTagsInCollection(dbCon *sql.DB, collection string) []models.ImageAndTag {
	query := fmt.Sprintf("SELECT Images.id, Images.Name, Tags.id, Tags.pid, Tags.name, t.tagCount\nFROM (\n  SELECT\n    id AS tagsId,\n    (\n        SELECT imageid FROM ImageTags WHERE ImageTags.tagid = tagsId ORDER BY RAND() LIMIT 1\n    ) AS imageId,\n    (\n        SELECT count(imageid) as count FROM ImageTags WHERE ImageTags.tagid = tagsId GROUP BY tagid\n    ) AS tagCount\n  FROM Tags\n) t\nJOIN Tags on t.tagsId = Tags.id\nJOIN Images on t.imageId = Images.id\nWHERE Tags.pid = %s \nORDER BY t.tagCount desc;", collection)
	log.Println(query)
	rows, err := dbCon.Query(query)

	if err != nil {
		panic(err)
	}

	return handleImageAndTagDbRows(rows)
}
