package main

type ContactRequest struct {
	Name          string `json:"name,omitempty"`
	Email         string `json:"email,omitempty"`
	Organization  string `json:"organization,omitempty"`
	Subject       string `json:"subject,omitempty"`
	Message       string `json:"message,omitempty"`
	RequestTarget string `json:"requestType,omitempty"`
}

type ContactResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
