package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/eferhatg/hayri-irdal/pkg/models"
	"github.com/eferhatg/hayri-irdal/pkg/scrapers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	deputyListPage    = "https://www.tbmm.gov.tr/develop/owa/milletvekillerimiz_sd.liste"
	deputyProfileRoot = "https://www.tbmm.gov.tr/develop/owa/milletvekillerimiz_sd.bilgi"
)

func main() {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		fmt.Print("can't initialize db", err)
	}

	db.AutoMigrate(&models.Deputy{})
	log.SetLevel(log.DebugLevel)

	//fetchTBMMDeputyList(db, deputyListPage)
	fetchDeputyProfiles(db)
}

func fetchTBMMDeputyList(db *gorm.DB, deputyListPage string) {
	deplist := scrapers.ScrapeTBMMDeputyList(deputyListPage)
	for _, d := range deplist {
		d.Upsert(db)
	}
	log.Info(len(deplist), "deputy is added")
}

func fetchDeputyProfiles(db *gorm.DB) {
	deputyList := []models.Deputy{}
	db.Find(&deputyList)
	for index := 0; index < len(deputyList); index++ {
		d := deputyList[index]
		logr := log.WithFields(log.Fields{"action": "update profile", "Deputy": d})
		updatedDeputy := scrapers.ScrapeDeputyProfilePage(d, deputyProfileRoot)
		db.Save(&updatedDeputy)
		logr.Debug("Deputy profile updated", updatedDeputy)
	}

}
