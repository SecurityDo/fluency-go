package model

import "time"

type User struct {
	Username        string          `json:"username"`
	Email           string          `json:"email"`
	FirstName       string          `json:"firstName"`
	LastName        string          `json:"lastName"`
	Organization    string          `json:"organization"`
	Rights          string          `json:"rights"`
	RadiusFlag      bool            `json:"radiusFlag"`
	RadiusServer    string          `json:"radiusServer"`
	RadiusGroup     bool            `json:"radiusGroup"`
	OauthProvider   string          `json:"oauthProvider"`
	OauthFlag       bool            `json:"oauthFlag"`
	MfaProvider     string          `json:"mfaProvider"`
	MfaFlag         bool            `json:"mfaFlag"`
	Customer        string          `json:"customer"`
	Restricted      bool            `json:"restricted"`
	Disabled        bool            `json:"disabled"`
	LastLogin       time.Time       `json:"lastLogin"`
	PasswordDate    time.Time       `json:"passwordDate"`
	Preference      *UserPreference `json:"preference"`
	Roles           []string        `json:"roles"`
	DataPolicies    []string        `json:"dataPolicies"`
	APIPolicies     []string        `json:"APIPolicies"`
	ConfirmPassword interface{}     `json:"confirm_password"`
	Password        string          `json:"password"`
	Profile         *UserProfile    `json:"profile"`
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

type GetUserRequest struct {
	Username string `json:"username"`
}

type GetUserResponse struct {
	Users *User `json:"user"`
}

type UserListResponse struct {
	Users []*User `json:"users"`
}

type UserAddRequest struct {
	User *User `json:"user"`
}

type UserDeleteRequest struct {
	Username string `json:"username"`
}

type UserUpdateInfoRequest struct {
	Username string `json:"username"`
	User     *User  `json:"user"`
}

type UserSuspendRequest struct {
	Username string `json:"username"`
}

type UserActivateRequest struct {
	Username string `json:"username"`
}
