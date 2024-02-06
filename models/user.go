package models

type User struct {
	Guid          string   `json:"guid"`
	RefreshTokens []string `json:"reftokens"`
}
