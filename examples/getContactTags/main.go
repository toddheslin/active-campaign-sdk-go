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

	tags, _, _ := campaign.Contacts.GetContactTags("5")

	for i := range tags.ContactTag {
		log.Printf("Tag: %s\tID :%s", tags.ContactTag[i].Tag, tags.ContactTag[i].ID)
		if tags.ContactTag[i].Tag == "21" {
			campaign.Contacts.RemoveContactTag(tags.ContactTag[i].ID)
		}
	}

}
