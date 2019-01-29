package models

import (
	"time"
)

type Tweet struct {
	ID                int       `gorm:"primary_key;column:id"`
	TwID              int64     `gorm:"column:tw_id"`
	FavoriteCount     int64     `gorm:"column:favorite_count"`
	RetweetCount      int64     `gorm:"column:retweet_count"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	Text              string    `gorm:"column:text"`
	Source            string    `gorm:"column:source"`
	InReplyToStatusID int64     `gorm:"column:in_reply_to_status_id"`
	UserID            int64     `gorm:"column:user_id"`
	Country           string    `gorm:"column:country"`
	City              string    `gorm:"column:city"`
	IsQoutedStatus    string    `gorm:"column:is_quote_status"`
	QoutedStatusText  string    `gorm:"column:quoted_status_text"`
	Hashtags          []string  `gorm:"column:hashtags"`
	Mentions          []string  `gorm:"column:mentions"`
	Medias            []string  `gorm:"column:medias"`
	Urls              []string  `gorm:"column:Urls"`
}

// func (d Deputy) Upsert(db *gorm.DB) {
// 	logr := log.WithFields(log.Fields{"action": "upsert", "Deputy": d})
// 	logr.Debug("Checking if deputy is in db ")

// 	count := 0
// 	db.Model(&Deputy{}).Where("deputy_no = ?", d.DeputyNo).Count(&count)
// 	if count == 0 {
// 		logr.Debug("Not found in db. inserting deputy ")
// 		if err := db.Create(&d).Error; err != nil {
// 			logr.Error(err)
// 		}
// 		logr.Debug("Deputy inserted")
// 	}

// }
