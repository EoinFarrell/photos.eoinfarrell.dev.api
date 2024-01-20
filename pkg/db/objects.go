package db

import (
	"github.com/eoinfarrell/photos.eoinfarrell.dev.api/pkg/models"
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func handleImageDbRows(rows *sql.Rows) []models.ImageInfo {
	var image models.ImageInfo
	var images []models.ImageInfo

	// the result object has a method called Next,
	// which is used to iterate through all returned rows.
	for rows.Next() {
		err := rows.Scan(&image.ID, &image.Name)
		if err != nil {
			log.Fatal(err)
		}

		images = append(images, image)
	}

	return images
}

func handleTagDbRows(rows *sql.Rows) []models.TagInfo {
	var tag models.TagInfo
	var tags []models.TagInfo

	// the result object has a method called Next,
	// which is used to iterate through all returned rows.
	for rows.Next() {
		err := rows.Scan(&tag.ID, &tag.Pid, &tag.Name)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(tag.ID, tag.Pid, tag.Name)
		tags = append(tags, tag)
	}

	return tags
}

func handleImageAndTagDbRows(rows *sql.Rows) []models.ImageAndTag {
	var tag models.TagInfo
	var image models.ImageInfo
	var tagCount int

	var imageAndTag []models.ImageAndTag

	// the result object has a method called Next,
	// which is used to iterate through all returned rows.
	for rows.Next() {
		err := rows.Scan(&image.ID, &image.Name, &tag.ID, &tag.Pid, &tag.Name, &tagCount)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(tag.ID, tag.Pid, tag.Name)
		imageAndTag = append(imageAndTag, models.ImageAndTag{image, tag, tagCount})
	}

	return imageAndTag
}
