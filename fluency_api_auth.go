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

func (r *FluencyClient) UserAdd(user *model.UserAddPayload) (err error) {

	input := user

	functionName := "userAdd"

	_, err = r.serviceClient.Call("api/auth", functionName, input)
	if err != nil {
		fmt.Printf("fail to call %s: %s", functionName, err.Error())
		return err
	}

	return nil
}

func (r *FluencyClient) UserDelete(username string) (err error) {

	input := model.UserDeletePayload{
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
