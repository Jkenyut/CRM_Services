package origin

type ResponseDefault struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode string `json:"statusCode"`
}
