package main

import (
	"log"
	"os"

	ac "github.com/toddheslin/active-campaign-sdk-go"
)

func main() {

	baseURL := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")

	campaign, err := ac.NewClient(
		&ac.ClientOpts{
			BaseUrl: baseURL,
			Token:   token,
		},
	)
	if err != nil {
		panic(err)
	}

	lists, _, _ := campaign.Contacts.GetContactLists("287199")

	for i := range lists.ContactLists {
		log.Printf("List: %s\tStatus :%s", lists.ContactLists[i].List, lists.ContactLists[i].Status)
	}

}
