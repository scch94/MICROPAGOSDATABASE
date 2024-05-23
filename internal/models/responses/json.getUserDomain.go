package responses

type DomainResponse struct {
	Response
	DomainName string `json:"domain"`
}
