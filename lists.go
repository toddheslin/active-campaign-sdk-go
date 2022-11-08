package active_campaign

import "net/http"

// TagsService handles communication with tag related
// methods of the Active Campaign API.
//
// Active Campaign API docs: https://developers.activecampaign.com/reference#tags
type ListsService service

// List are labels that you can apply to contacts to help you organize them.
// The API enables you to add, view, update, and delete tags.
type List struct {
	Name           string `json:"name"`
	StringID       string `json:"stringid"`
	SenderURL      string `json:"sender_url,omitempty"`
	SenderReminder string `json:"sender_reminder,omitempty"`
	User           string `json:"user,omitempty"`
	Links          *Links `json:"links"`
	ID             string `json:"id"`
}

// AddListToContacRequest is the request body used for adding a tag to a contact.
type AddListToContacRequest struct {
	List *ListToContact `json:"contactList"`
}

// ListToConta is used to add a list to a contact.
type ListToContact struct {
	List    string `json:"list"`
	Contact string `json:"contact"`
	Status  string `json:"status"`
}

// ListResponse is the response body returned from creating or retrieving a tag.
type ListResponse struct {
	List *ListToContact `json:"contactList"`
}

// ListAllResponse is the response body returned from listing all tags.
type ListsAllResponse struct {
	List []*List `json:"lists"`
	Meta *Meta   `json:"meta"`
}

// Lists all tags.
func (s *ListsService) ListAll() (*ListsAllResponse, *Response, error) {
	u := "lists"
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := &ListsAllResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}

// AddContactToList adds a contact to a List.
func (s *ListsService) AddContactToList(contact *AddListToContacRequest) (*AddListToContacRequest, *Response, error) {
	u := "contactLists"
	req, err := s.client.NewRequest(http.MethodPost, u, contact)
	if err != nil {
		return nil, nil, err
	}

	c := &AddListToContacRequest{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}
