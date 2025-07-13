package presenter

import "encoding/json"

type AuthenticationResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

func (r AuthenticationResponse) String() string {
	o, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(o)
}
