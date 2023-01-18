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
	user = user
	//PrettyPrintJSON(user)
	fmt.Println("testUserList PASS")
}

func testUserAdd(client *FluencyClient) {

	user := &model.User{
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
	fmt.Println("testUserAdd PASS")
}

func testGetUser(client *FluencyClient) {
	user, err := client.GetUser("test")
	if err != nil {
		panic(err.Error())
	}
	user = user
	//PrettyPrintJSON(user)
	fmt.Println("testGetUser PASS")
}

func testUserDelete(client *FluencyClient) {
	err := client.UserDelete("test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("testUserDelete PASS")
}

func testUserUpdateInfo(client *FluencyClient) {

	user := &model.User{
		APIPolicies:     []string{},
		DataPolicies:    []string{},
		Email:           "example@test.com",
		FirstName:       "NewTest",
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

	err := client.UserUpdateInfo("test", user)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("testUserUpdateInfo PASS")
}

func testUserSuspend(client *FluencyClient) {
	err := client.UserSuspend("test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("testUserSuspend PASS")
}

func testUserActivate(client *FluencyClient) {
	err := client.UserActivate("test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("testUserActivate PASS")
}

func testAuth(client *FluencyClient) {
	testUserList(client)
	testUserAdd(client)
	testUserUpdateInfo(client)
	testGetUser(client)
	testUserSuspend(client)
	testUserActivate(client)
	testUserDelete(client)
}
