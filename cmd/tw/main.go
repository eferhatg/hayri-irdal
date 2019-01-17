package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	log "github.com/Sirupsen/logrus"
	"github.com/eferhatg/hayri-irdal/pkg/models"
	"github.com/eferhatg/hayri-irdal/pkg/twitter"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	api := anaconda.NewTwitterApiWithCredentials(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		fmt.Print("can't initialize db", err)
	}

	db.AutoMigrate(&models.Deputy{})
	log.SetLevel(log.DebugLevel)
	fetchUserProfiles(api, db)
	fmt.Println("--------------------")
}

func fetchUserProfiles(api *anaconda.TwitterApi, db *gorm.DB) {
	deputyList := []models.Deputy{}
	db.Where("tw_link != ?", "").Find(&deputyList)
	if len(deputyList) == 0 {
		log.Info("Not found available deputy")
	}
	for index := 0; index < len(deputyList); index++ {
		log.Info("Getting twitter profile of ", deputyList[index].TwLink)
		twd := twitter.TwDeputy{Deputy: &deputyList[index]}
		err := twd.GetUserProfile(api)
		if err != nil {
			log.Error(err)
			continue
		}
		db.Save(&twd.Deputy)
		log.Info("Saved ", twd.Deputy.TwLink)

	}
}
