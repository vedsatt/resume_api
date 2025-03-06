package vk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserData struct {
	Id     int    `json:"id"`
	Domain string `json:"domain"`
	Bdate  string `json:"bdate"`
	Fname  string `json:"first_name"`
	Lname  string `json:"last_name"`
}

type vkResp struct {
	Response []UserData `json:"response"`
}

func GetData(token string) (*UserData, error) {
	url := fmt.Sprintf("https://api.vk.com/method/users.get?access_token=%s&v=5.199&fields=bdate,domain",
		token)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var vkResponse vkResp
	err = json.NewDecoder(resp.Body).Decode(&vkResponse)
	if err != nil {
		return nil, err
	}

	return &vkResponse.Response[0], nil
}
