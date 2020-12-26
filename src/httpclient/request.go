package httpclient

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// RequestI Interface
type RequestI interface {
	GetByURLWithParams(url string, params map[string]string) []byte
}

// RequestT Type
type RequestT struct {
}

// GetByURLWithParams Method
func (RequestT) GetByURLWithParams(url string, params map[string]string) []byte {
	request, error := http.NewRequest("GET", url, nil)

	if error != nil {
		log.Print(error.Error())
		os.Exit(1)
	}

	query := request.URL.Query()

	for key, value := range params {
		query.Add(key, value)
	}

	request.URL.RawQuery = query.Encode()

	client := &http.Client{}

	response, error := client.Do(request)

	if error != nil {
		log.Print(error.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		log.Print(error.Error())
		os.Exit(1)
	}

	return body
}
