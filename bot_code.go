package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/segmentio/encoding/json"
)

type GetMeT struct {
	Ok     bool         `json:"ok"`
	Result GetMeResultT `json: "result"`
}

type GetMeResultT struct {
	Id        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type SendMessageT struct {
	Ok     bool     `json:"ok"`
	Result MessageT `json:"result"`
}

type MessageT struct {
	MessageID int                          `json:"message_id"`
	From      GetUpdatesResultMessageFromT `json:"from"`
	Chat      GetUpdatesResultMessageChatT `json:"chat"`
	Date      int                          `json:"date"`
	Text      string                       `json:"text"`
}

type GetUpdatesResultMessageFromT struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type GetUpdatesResultMessageChatT struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type GetUpdatesT struct {
	Ok     bool                `json:"ok"`
	Result []GetUpdatesResultT `json:"result"`
}

type GetUpdatesResultT struct {
	UpdateID int                `json:"update_id"`
	Message  GetUpdatesMessageT `json:"message,omitempty"`
}

type GetUpdatesMessageT struct {
	MessageID int `json:"message_id"`
	From      struct {
		ID           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	} `json:"from"`
	Chat struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	} `json:"chat"`
	Date int    `json:"date"`
	Text string `json:"text"`
}

const telegramBaseUrl = "xxxxxxxxxxxxxxxxxx"
const telegramToken = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

const methodGetMe = "getMe"
const methodGetUpdates = "getUpdates"
const methodSendMessage = "sendMessage"

func main() {
	getUpdates := GetUpdatesT{}
	body := getBodyByUrl(getUrlByMethod(methodGetUpdates))

	err := json.Unmarshal(body, &getUpdates)
	if err != nil {
		fmt.Printf(" Error in unmarshal getUpdates: %s", err.Error())

		return
	}

	sendMessageUrl := getUrlByMethod(methodSendMessage)
	for _, item := range getUpdates.Result {
		formattedText := strings.ToLower(item.Message.Text)
		if strings.Contains(formattedText, "go") {
			fmt.Println(item.Message.From.FirstName + ": " + item.Message.Text)

			body := getBodyByUrl(sendMessageUrl + "?chat_id=" + strconv.Itoa(item.Message.Chat.ID) + "&text=" + item.Message.From.FirstName)
			fmt.Println(string(body))
		}
	}
}

func getUrlByMethod(methodName string) string {
	return telegramBaseUrl + telegramToken + "/" + methodName
}

func getBodyByUrl(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	return body
}
