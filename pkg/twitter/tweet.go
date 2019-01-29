package twitter

import (
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/eferhatg/hayri-irdal/pkg/models"
)

type TwTweet struct {
	Tweet *models.Tweet
}

func (twd *TwDeputy) GetUserTweets(api *anaconda.TwitterApi) error {
	user, err := api.GetUsersShow(twd.Deputy.TwLink, nil)
	if err != nil {

		return err
	}
	timeLayout := "Mon Jan 02 15:04:05 -0700 2006"

	twd.Deputy.TwDescription = user.Description
	twd.Deputy.TwFavouritesCount = user.FavouritesCount
	twd.Deputy.TwFollowersCount = user.FollowersCount
	twd.Deputy.TwFriendsCount = user.FriendsCount
	twd.Deputy.TwListedCount = user.ListedCount
	twd.Deputy.TwLocation = user.Location
	twd.Deputy.TwName = user.Name
	twd.Deputy.TwProfileBackgroundImage = user.ProfileBackgroundImageUrlHttps
	twd.Deputy.TwProfileImage = user.ProfileImageUrlHttps
	twd.Deputy.TwStatusCount = user.StatusesCount
	twd.Deputy.TwVerified = user.Verified
	twd.Deputy.TwCreatedAt, _ = time.Parse(timeLayout, user.CreatedAt)
	twd.Deputy.TwId = user.Id
	return nil

}
