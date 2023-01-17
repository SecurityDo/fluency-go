package model

import "time"

type UserListResponse struct {
	Users []*User `json:"users"`
}

type User struct {
	Username      string          `json:"username"`
	Email         string          `json:"email"`
	FirstName     string          `json:"firstName"`
	LastName      string          `json:"lastName"`
	Organization  string          `json:"organization"`
	Rights        string          `json:"rights"`
	RadiusFlag    bool            `json:"radiusFlag"`
	RadiusServer  string          `json:"radiusServer"`
	RadiusGroup   bool            `json:"radiusGroup"`
	OauthProvider string          `json:"oauthProvider"`
	OauthFlag     bool            `json:"oauthFlag"`
	MfaProvider   string          `json:"mfaProvider"`
	MfaFlag       bool            `json:"mfaFlag"`
	Customer      string          `json:"customer"`
	Restricted    bool            `json:"restricted"`
	Disabled      bool            `json:"disabled"`
	LastLogin     time.Time       `json:"lastLogin"`
	PasswordDate  time.Time       `json:"passwordDate"`
	Preference    *UserPreference `json:"preference"`
	Roles         []string        `json:"roles"`
	DataPolicies  []string        `json:"dataPolicies"`
	APIPolicies   []string        `json:"APIPolicies"`
	Profile       *UserProfile    `json:"profile"`
}

type UserPreference struct {
	Timezone *Timezone `json:"timezone"`
}

type Timezone struct {
	Value  string `json:"value"`
	Policy string `json:"policy"`
}

type UserProfile struct {
	Homepage string `json:"homepage"`
}

type UserAddPayload struct {
	APIPolicies     []string     `json:"APIPolicies"`
	DataPolicies    []string     `json:"dataPolicies"`
	Email           string       `json:"email"`
	FirstName       string       `json:"firstName"`
	LastName        string       `json:"lastName"`
	OauthFlag       bool         `json:"oauthFlag"`
	OauthProvider   string       `json:"oauthProvider"`
	Organization    string       `json:"organization"`
	RadiusFlag      bool         `json:"radiusFlag"`
	RadiusGroup     bool         `json:"radiusGroup"`
	MfaFlag         bool         `json:"mfaFlag"`
	MfaProvider     string       `json:"mfaProvider"`
	RadiusServer    string       `json:"radiusServer"`
	Rights          string       `json:"rights"`
	Roles           []string     `json:"roles"`
	Username        string       `json:"username"`
	Restricted      bool         `json:"restricted"`
	Customer        string       `json:"customer"`
	ConfirmPassword interface{}  `json:"confirm_password"`
	Password        string       `json:"password"`
	Profile         *UserProfile `json:"profile"`
}

type UserDeletePayload struct {
	Username string `json:"username"`
}
