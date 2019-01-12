package models

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

//Deputy - Milletvekili tablosu
type Deputy struct {
	ID                 int       `gorm:"primary_key;column:id"`
	Name               string    `gorm:"column:name"`
	Surname            string    `gorm:"column:surname"`
	DeputyNo           int       `gorm:"column:deputy_no"`
	Party              string    `gorm:"column:party"`
	TermNo             int       `gorm:"column:term_no"`
	Region             string    `gorm:"column:region"`
	Address            string    `gorm:"column:address"`
	Tel                string    `gorm:"column:tel"`
	Fax                string    `gorm:"column:fax"`
	Email              string    `gorm:"column:email"`
	Web                string    `gorm:"column:web"`
	FbLink             string    `gorm:"column:fb_link"`
	TwLink             string    `gorm:"column:tw_link"`
	Position           string    `gorm:"column:position"`
	Cv                 string    `gorm:"column:cv"`
	ProfilePic         string    `gorm:"column:profile_picture"`
	ActivityUpdateDate time.Time `gorm:"column:activity_update_date"`
	IsActive           bool      `gorm:"column:is_active"`
	BirthDay           time.Time `gorm:"column:birth_day"`
	Profession         string    `gorm:"column:profession"`
	BirthPlace         string    `gorm:"column:birth_place"`
	University         string    `gorm:"column:university"`
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
