package db

import (
	"database/sql"
	"github.com/eoinfarrell/photos.eoinfarrell.dev.api/pkg/aws/s3"
	"log"
)

func deleteUnusedImages(dbCon *sql.DB) {
	var notAv = 0
	error := 0
	success := 0

	for i := 0; i <= 2; i++ {
		var start = i * 50
		var end = start + 50

		for _, image := range GetImagesFromDbByIdLimit(dbCon, start, end) {
			resp, _ := s3.GetImage(image)

			if resp == nil {
				error++
				// log.Println(escapedUrl)
			} else {
				if resp.StatusCode == 403 {
					//log.Println("https://photos-efarrell.s3.eu-west-1.amazonaws.com/" + image.Name)
					notAv++
				} else {
					// log.Println(resp)
					success++
				}
			}
		}
	}

	log.Println(success)
	log.Println(notAv)
	log.Println(error)
}
