package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {

	api := anaconda.NewTwitterApiWithCredentials(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	v := url.Values{}
	v.Set("count", "1")
	v.Set("screen_name", "YavuzSubasi")
	tweets, err := api.GetUserTimeline(v)
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(tweets)
	fmt.Println(string(b))
	// for _, tweet := range tweets {
	// 	fmt.Println(tweet.Text)
	// }

}
