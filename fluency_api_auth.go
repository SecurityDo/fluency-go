package fluency

import (
	"encoding/json"
	"fmt"
	"github.com/SecurityDo/fluency-go/model"
)

func (r *FluencyClient) UserList() (result *model.UserListResponse, err error) {

	functionName := "userList"

	res, err := r.serviceClient.Call("api/auth", functionName, nil)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Printf("fail to parse response of %s: %s", functionName, err.Error())
		return nil, err
	}

	return result, nil
}

func (r *FluencyClient) UserAdd(user *model.User) (err error) {

	input := &model.UserAddRequest{User: user}

	functionName := "userAdd"

	_, err = r.serviceClient.Call("api/auth", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

func (r *FluencyClient) GetUser(username string) (result *model.GetUserResponse, err error) {

	input := model.GetUserRequest{
		Username: username,
	}

	functionName := "getUser"

	res, err := r.serviceClient.Call("api/auth", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return nil, err
	}

	err = json.Unmarshal(res.GetBytes(), &result)
	if err != nil {
		fmt.Printf("fail to parse response of %s: %s", functionName, err.Error())
		return nil, err
	}

	return result, nil
}

func (r *FluencyClient) UserDelete(username string) (err error) {

	input := model.UserDeleteRequest{
		Username: username,
	}

	functionName := "userDelete"

	_, err = r.serviceClient.Call("api/auth", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

func (r *FluencyClient) UserUpdateInfo(username string, user *model.User) (err error) {

	input := &model.UserUpdateInfoRequest{
		Username: username,
		User:     user,
	}

	functionName := "userUpdateInfo"

	_, err = r.serviceClient.Call("api/auth", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

func (r *FluencyClient) UserSuspend(username string) (err error) {

	input := model.UserSuspendRequest{
		Username: username,
	}

	functionName := "userSuspend"

	_, err = r.serviceClient.Call("api/auth", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

func (r *FluencyClient) UserActivate(username string) (err error) {

	input := model.UserActivateRequest{
		Username: username,
	}

	functionName := "userActivate"

	_, err = r.serviceClient.Call("api/auth", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}
