package active_campaign

import (
	"fmt"
	"net/http"
)

// ContactTag is used to add a tag to a contact.
type ContactTag struct {
	CDate   string `json:"cdate,omitempty"`
	Contact string `json:"contact"`
	ID      string `json:"id,omitempty"`
	Links   *struct {
		Contact string `json:"contact,omitempty"`
		Tag     string `json:"tag,omitempty"`
	} `json:"links,omitempty"`
	Tag string `json:"tag"`
}

// AddTagToContactRequest is the request body used for adding a tag to a contact.
type AddTagToContactRequest struct {
	ContactTag *ContactTag `json:"contactTag"`
}

// AddTagToContactResponse is the response body from adding a tag to a contact.
type AddTagToContactResponse struct {
	ContactTag *ContactTag `json:"contactTag"`
}

// GetContactTagResponse is the response body from getting contact tags.
type GetContactTagResponse struct {
	ContactTag *[]ContactTag `json:"contactTags"`
}

// AddTagToContact adds a tag to a contact.
func (s *ContactsService) AddTagToContact(contact *AddTagToContactRequest) (*AddTagToContactResponse, *Response, error) {
	u := "contactTags"
	req, err := s.client.NewRequest(http.MethodPost, u, contact)
	if err != nil {
		return nil, nil, err
	}

	c := &AddTagToContactResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}

// GetContactTags Retrieving a contact's tag id's
func (s *ContactsService) GetContactTags(contactId string) (*GetContactTagResponse, *Response, error) {
	u := fmt.Sprintf("contacts/%s/contactTags", contactId)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := &GetContactTagResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}

// RemoveContactTag removed a tage from a contact ising the resault of GetContactTags
func (s *ContactsService) RemoveContactTag(tagid string) (*Response, error) {
	u := fmt.Sprintf("contactTags/%s", tagid)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return resp, nil
}
