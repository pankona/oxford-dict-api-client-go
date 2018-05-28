package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/pankona/oxford-dict-api-client-go/client"
)

func main() {
	cfg := client.NewConfiguration()
	cfg.BasePath = "https://od-api.oxforddictionaries.com/api/v1"
	c := client.NewAPIClient(cfg)

	var (
		appID  = os.Getenv("OD_API_APP_ID")
		appKey = os.Getenv("OD_API_APP_KEY")
	)
	if appID == "" || appKey == "" {
		fmt.Println("error. please specify both OD_API_APP_ID and OD_API_APP_KEY")
		return
	}

	wl, _, err := c.WordlistApi.WordlistSourceLangFiltersBasicGet(
		context.Background(),
		"en",
		"registers=Rare",
		appID,
		appKey,
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m := (*wl.Metadata).(map[string]interface{})
	fmt.Println(m)
}
