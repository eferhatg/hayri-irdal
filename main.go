package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {

	api := anaconda.NewTwitterApiWithCredentials(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	user, err := api.GetUsersShow(os.Args[1], nil)
	if err != nil {
		fmt.Println(err)
	}
	userjson, err := json.Marshal(user)
	fmt.Println(string(userjson))
	fmt.Println("--------------------")
}

/*
description
statuses_count
verified
favourites_count
followers_count
friends_count
listed_count
location
name
profile_background_image_url_https
profile_image_url_https
*/
