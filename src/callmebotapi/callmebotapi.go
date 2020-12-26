package callmebotapi

import (
	"fmt"
	"httpclient"
	"strings"
)

// CallMeBotI Interface
type CallMeBotI interface {
	CallUser(username string, text string) bool
}

// CallMeBotT Tyep
type CallMeBotT struct {
}

const baseCallURL = "http://api.callmebot.com/start.php"

// CallUser Method
func (CallMeBotT) CallUser(username string, text string) bool {
	params := map[string]string{"user": "@" + username, "text": text, "lang": "ru-RU-Standard-D", "rpt": fmt.Sprintf("%d", 1)}
	body := httpclient.RequestT{}.GetByURLWithParams(baseCallURL, params)

	responseContent := string(body)

	return !strings.Contains(responseContent, "Warning! User not authorized. Click")
}
