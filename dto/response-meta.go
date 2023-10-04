package dto

type ResponseMeta struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
	StatusCode   string `json:"StatusCode"`
}
