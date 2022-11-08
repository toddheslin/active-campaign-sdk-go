package active_campaign

import "net/http"

// CustomFieldsService handles communication with custom field related
// methods of the Active Campaign API.
//
// Active Campaign API docs: https://developers.activecampaign.com/reference/retrieve-fields
type CustomFieldsService service

// Tags are labels that you can apply to contacts to help you organize them.
// The API enables you to add, view, update, and delete tags.
type CustomField struct {
	Type         string `json:"type,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"descript,omitempty"`
	Required     string `json:"isrequired,omitempty"`
	PersTag      string `json:"perstag,omitempty"`
	DefaultValue string `json:"defval,omitempty"`
	Visible      string `json:"visible,omitempty"`
	OrderNum     string `json:"ordernum,omitempty"`
	Links        *Links `json:"links,omitempty"`
	ID           string `json:"id"`
}

// CreateCustomFieldRequest is the request body used for creating a custom field.
type CreateCustomFieldRequest struct {
	Field *CreatedCustomField `json:"field"`
}

// CreatedCustomField is a struct embedded to creating or retrieving a custom field.
type CreatedCustomField struct {
	Type         string `json:"type"`
	Title        string `json:"title"`
	Description  string `json:"descript"`
	Required     string `json:"isrequired"`
	PersTag      string `json:"perstag"`
	DefaultValue string `json:"defval"`
	Visible      string `json:"visible"`
	OrderNum     string `json:"ordernum"`
	Links        *Links `json:"links"`
	ID           string `json:"id"`
}

// CustomFieldResponse is the response body returned from creating or retrieving a tag.
type CustomFieldResponse struct {
	Field *CustomField `json:"field"`
}

// CustomFieldsListAllResponse is the response body returned from listing all tags.
type CustomFieldsListAllResponse struct {
	CustomFields []*CustomField `json:"fields"`
	Meta         *Meta          `json:"meta"`
}

// Create a tag.
func (s *CustomFieldsService) Create(tag *CreateTagRequest) (*CustomFieldResponse, *Response, error) {
	u := "fields"
	req, err := s.client.NewRequest(http.MethodPost, u, tag)
	if err != nil {
		return nil, nil, err
	}

	c := &CustomFieldResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}

// Retrieve a tag.
func (s *CustomFieldsService) Retrieve(id string) (*CustomFieldResponse, *Response, error) {
	u := "fields/" + id
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := &CustomFieldResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}

// Lists all tags.
func (s *CustomFieldsService) ListAll() (*CustomFieldsListAllResponse, *Response, error) {
	u := "fields"
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := &CustomFieldsListAllResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}
