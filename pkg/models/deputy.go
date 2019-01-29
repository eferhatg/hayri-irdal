package models

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

//Deputy - Milletvekili tablosu
type Deputy struct {
	ID                       int       `gorm:"primary_key;column:id"`
	Name                     string    `gorm:"column:name"`
	Surname                  string    `gorm:"column:surname"`
	DeputyNo                 int       `gorm:"column:deputy_no"`
	Party                    string    `gorm:"column:party"`
	TermNo                   int       `gorm:"column:term_no"`
	Region                   string    `gorm:"column:region"`
	Address                  string    `gorm:"column:address"`
	Tel                      string    `gorm:"column:tel"`
	Fax                      string    `gorm:"column:fax"`
	Email                    string    `gorm:"column:email"`
	Web                      string    `gorm:"column:web"`
	FbLink                   string    `gorm:"column:fb_link"`
	TwLink                   string    `gorm:"column:tw_link"`
	Position                 string    `gorm:"column:position"`
	Cv                       string    `gorm:"column:cv"`
	ProfilePic               string    `gorm:"column:profile_picture"`
	ActivityUpdateDate       time.Time `gorm:"column:activity_update_date"`
	IsActive                 bool      `gorm:"column:is_active"`
	BirthDay                 time.Time `gorm:"column:birth_day"`
	Profession               string    `gorm:"column:profession"`
	BirthPlace               string    `gorm:"column:birth_place"`
	University               string    `gorm:"column:university"`
	TwDescription            string    `gorm:"column:tw_description"`
	TwStatusCount            int64     `gorm:"column:tw_statuses_count"`
	TwVerified               bool      `gorm:"column:tw_verified"`
	TwFavouritesCount        int       `gorm:"column:tw_favourites_count"`
	TwFollowersCount         int       `gorm:"column:tw_followers_count"`
	TwFriendsCount           int       `gorm:"column:tw_friends_count"`
	TwListedCount            int64     `gorm:"column:tw_listed_count"`
	TwLocation               string    `gorm:"column:tw_location"`
	TwName                   string    `gorm:"column:tw_name"`
	TwProfileBackgroundImage string    `gorm:"column:tw_profile_background_image_url_https"`
	TwProfileImage           string    `gorm:"column:tw_profile_image_url_https"`
	TwCreatedAt              time.Time `gorm:"column:tw_created_at"`
	TwId                     int64     `gorm:"column:tw_id"`
}

//Upsert inserting deputy if not found in db
func (d Deputy) Upsert(db *gorm.DB) {
	logr := log.WithFields(log.Fields{"action": "upsert", "Deputy": d})
	logr.Debug("Checking if deputy is in db ")

	count := 0
	db.Model(&Deputy{}).Where("deputy_no = ?", d.DeputyNo).Count(&count)
	if count == 0 {
		logr.Debug("Not found in db. inserting deputy ")
		if err := db.Create(&d).Error; err != nil {
			logr.Error(err)
		}
		logr.Debug("Deputy inserted")
	}

}
