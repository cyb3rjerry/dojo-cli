package dojo

type CredentialsList struct {
	Count       int          `json:"count"`
	Next        string       `json:"next"`
	Previous    string       `json:"previous"`
	Credentials []Credential `json:"results"`
}

type Credential struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	Role           string `json:"role"`
	Authentication string `json:"authentication"`      // "Form" or "SSO"
	HttpAuth       string `json:"http_authentication"` // "Basic" or "NTLM"
	Description    string `json:"description"`
	Url            string `json:"url"`
	LoginRegex     string `json:"login_regex"`
	LogoutRegex    string `json:"logout_regex"`
	IsValid        bool   `json:"is_valid"`
	Environment    int    `json:"environment"`
	Notes          []int  `json:"notes"`
}
