package main

import (
	"encoding/json"
	"fmt"
)

// GetMeT struct
type GetMeT struct {
	Ok     bool         `json:"ok"`
	Result GetMeResultT `json:"result"`
}

// GetMeResultT struct
type GetMeResultT struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

func main() {
	getMeJSON := `{
		"ok": true,
		"result": {
			"id": XXXXXXXXXX,
			"is_bot": true,
			"first_name": "XXXXXXXX",
			"username": "XXXXXXX",
			"can_join_groups": true,
			"can_read_all_group_messages": false,
			"supports_inline_queries": false
		}
	}`

	getMe := GetMeT{}

	err := json.Unmarshal([]byte(getMeJSON), &getMe)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(getMe)