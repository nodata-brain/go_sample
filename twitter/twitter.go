package twitter

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"

)


func main() {

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

    v := url.Values{}
    v.Set("count", "10")
    searchResult, _ := api.GetSearch("#golang", v)
    for _, tweet := range searchResult.Statuses {
        fmt.Printf("[%s] %s\n", tweet.CreatedAt, tweet.Text)
    }

}
