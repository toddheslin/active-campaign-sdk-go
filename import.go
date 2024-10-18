package active_campaign

import "net/http"

type ImportsService service

type SubscribeValues struct {
	ListId uint64 `json:"listid"`
}

type ImportFieldValues struct {
	Id    uint64      `json:"id"`
	Value interface{} `json:"value"`
}

type ImportContact struct {
	Email       string              `json:"email"`
	FirstName   string              `json:"first_name,omitempty"`
	LastName    string              `json:"last_name,omitempty"`
	Phone       string              `json:"phone,omitempty"`
	Tags        []string            `json:"tags"`
	Fields      []ImportFieldValues `json:"fields"`
	Subscribe   []SubscribeValues   `json:"subscribe"`
	Unsubscribe []SubscribeValues   `json:"unsubscribe"`
}

type BulkImportRequest struct {
	Contacts           []*ImportContact `json:"contacts"`
	ExcludeAutomations bool             `json:"exclude_automations"`
}

type BulkImportResponse struct {
	Success        uint8  `json:"Success"`
	QueuedContacts uint16 `json:"queued_contacts"`
	BatchId        string `json:"batchId"`
	Message        string `json:"message"`
}

func (s *ImportsService) BulkImport(contacts *BulkImportRequest) (*BulkImportResponse, *Response, error) {
	url := "import/bulk_import"
	req, err := s.client.NewRequest(http.MethodPost, url, contacts)
	if err != nil {
		return nil, nil, err
	}
	i := &BulkImportResponse{}
	resp, err := s.client.Do(req, i)
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	return i, resp, nil
}
