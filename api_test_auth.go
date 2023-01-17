package fluency

import (
	"fmt"
	"github.com/SecurityDo/fluency-go/model"
)

func testUserList(client *FluencyClient) {
	user, err := client.UserList()
	if err != nil {
		panic(err.Error())
	}
	PrettyPrintJSON(user)
}

func testUserAdd(client *FluencyClient) {

	user := &model.UserAddPayload{
		APIPolicies:     []string{},
		DataPolicies:    []string{},
		Email:           "test@test.com",
		FirstName:       "Test",
		LastName:        "Test",
		OauthFlag:       false,
		OauthProvider:   "",
		Organization:    "test",
		RadiusFlag:      false,
		RadiusGroup:     false,
		MfaFlag:         false,
		MfaProvider:     "",
		RadiusServer:    "",
		Rights:          "analyst",
		Roles:           []string{},
		Username:        "test",
		Restricted:      false,
		Customer:        "",
		ConfirmPassword: nil,
		Password:        "",
		Profile:         &model.UserProfile{},
	}

	err := client.UserAdd(user)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("OK")
}

func testUserDelete(client *FluencyClient) {
	err := client.UserDelete("test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("OK")
}
