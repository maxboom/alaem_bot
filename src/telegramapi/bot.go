package telegramapi

import (
	"encoding/json"
	"fmt"
	"httpclient"
)

// BotI Interface
type BotI interface {
	GetMe(bot BotSettingsT) GetMeT
	GetUpdates(bot BotSettingsT) GetUpdatesT
	SendMessage(bot BotSettingsT, sendMessage SendMessageT)
}

// BotT Type
type BotT struct {
}

const baseURL = "https://api.telegram.org/%s"
const methodURL = baseURL + "/%s"

const getMeMethod = "getMe"
const getUpdates = "getUpdates"
const methodSendMessage = "sendMessage"

// GetMe method
func (BotT) GetMe(bot BotSettingsT) GetMeT {
	url := fmt.Sprintf(methodURL, bot.Token, getMeMethod)

	params := make(map[string]string, 0)
	body := httpclient.RequestT{}.GetByURLWithParams(url, params)

	meData := GetMeT{}

	error := json.Unmarshal(body, &meData)

	if error != nil {
		fmt.Println(error.Error())
	}

	return meData
}

// GetUpdates method
func (s BotT) GetUpdates(bot BotSettingsT, getUpdatesRequest GetUpdatesRequestT) GetUpdatesT {
	params := map[string]string{"offset": fmt.Sprintf("%d", getUpdatesRequest.Offset)}

	url := fmt.Sprintf(methodURL, bot.Token, getUpdates)

	body := httpclient.RequestT{}.GetByURLWithParams(url, params)

	updatesData := GetUpdatesT{}

	error := json.Unmarshal(body, &updatesData)

	if error != nil {
		fmt.Println(error.Error())
	}

	return updatesData
}

// SendMessage method
func (s BotT) SendMessage(bot BotSettingsT, sendMessage SendMessageT) {
	params := map[string]string{"chat_id": fmt.Sprintf("%d", sendMessage.ChatID), "text": sendMessage.Message}

	url := fmt.Sprintf(methodURL, bot.Token, methodSendMessage)

	httpclient.RequestT{}.GetByURLWithParams(url, params)
}
