package active_campaign

import (
	"fmt"
	"net/http"
)

type GetContactListsResponse struct {
	ContactLists []struct {
		Contact               string      `json:"contact"`
		List                  string      `json:"list"`
		Form                  string      `json:"form"`
		Seriesid              string      `json:"seriesid"`
		Sdate                 string      `json:"sdate"`
		Udate                 interface{} `json:"udate"`
		Status                string      `json:"status"`
		Responder             string      `json:"responder"`
		Sync                  string      `json:"sync"`
		Unsubreason           string      `json:"unsubreason"`
		Campaign              interface{} `json:"campaign"`
		Message               interface{} `json:"message"`
		FirstName             string      `json:"first_name"`
		LastName              string      `json:"last_name"`
		IP4Sub                string      `json:"ip4Sub"`
		Sourceid              string      `json:"sourceid"`
		AutosyncLog           interface{} `json:"autosyncLog"`
		IP4Last               string      `json:"ip4_last"`
		IP4Unsub              string      `json:"ip4Unsub"`
		CreatedTimestamp      string      `json:"created_timestamp"`
		UpdatedTimestamp      string      `json:"updated_timestamp"`
		CreatedBy             string      `json:"created_by"`
		UpdatedBy             string      `json:"updated_by"`
		UnsubscribeAutomation interface{} `json:"unsubscribeAutomation"`
		Links                 struct {
			Automation            string `json:"automation"`
			List                  string `json:"list"`
			Contact               string `json:"contact"`
			Form                  string `json:"form"`
			AutosyncLog           string `json:"autosyncLog"`
			Campaign              string `json:"campaign"`
			UnsubscribeAutomation string `json:"unsubscribeAutomation"`
			Message               string `json:"message"`
		} `json:"links"`
		ID         string      `json:"id"`
		Automation interface{} `json:"automation"`
	} `json:"contactLists"`
}

// GetContactLists Retrieving a contact's lists's
func (s *ContactsService) GetContactLists(contactId string) (*GetContactListsResponse, *Response, error) {
	u := fmt.Sprintf("contacts/%s/contactLists", contactId)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := &GetContactListsResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}
