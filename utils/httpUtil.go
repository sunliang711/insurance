package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// PostJson TODO
// 2019/08/16 10:45:16
func PostJson(url string, headers map[string]string, body interface{}) (*http.Response, error) {
	bs, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}
	log.Debugf("PostJson body: %v", string(bs))

	request, err := http.NewRequest("POST", url, bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		request.Header.Add(k, v)
	}

	client := &http.Client{}
	return client.Do(request)
}
