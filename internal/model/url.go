package urlModel

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"

	"gorm.io/gorm"
)

type Url struct {
	OldURL       string `json:"url"`
	ShortenedURL string `gorm:"primaryKey" json:"shortened_url"`
}

func (Url) TableName() string {
	return "urls"
}

func ShortenUrl(oldUrl string, db *gorm.DB) (Url, error) {
	shortUrl := hashTheUrl(oldUrl)
	url := Url{OldURL: oldUrl, ShortenedURL: shortUrl}
	result := db.Create(&url)
	if result.Error != nil {
		log.Println("Error While inserting")
		return Url{}, result.Error
	}
	if result.RowsAffected == 0 {
		log.Println("error occured while inserting")
		return Url{}, errors.New("Nothing got inserted in the db")
	}

	return url, nil
}

func GetOldUrl(newUrl string, db *gorm.DB) (string, error) {
	url := Url{}
	result := db.Find(&url, "shortened_url = ?", newUrl)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("No Old url found with that perticular url")
			return "", gorm.ErrRecordNotFound
		}
		log.Println("Error Occured while finding")
		return "", result.Error
	}
	return url.OldURL, nil
}

func hashTheUrl(url string) string {
	hashed := sha256.Sum256([]byte(url))

	return hex.EncodeToString(hashed[:][:5])
}
