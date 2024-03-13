package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	ResponseData struct {
		TranslatedText string `json:"translatedText"`
	} `json:"responseData"`

	QuotaFinished bool `json:"quotaFinished"`
}

type TranslationConfig struct {
	text  string
	from  string
	to    string
	email string // 10x more translations if you provide an email
}

func Translate(tc TranslationConfig) (string, error) {
	params := url.Values{}
	params.Add("q", tc.text)
	if tc.email != "" {
		params.Add("de", tc.email)
	}
	params.Add("langpair", fmt.Sprintf("%s|%s", tc.from, tc.to))

	resp, err := http.Get("https://api.mymemory.translated.net/get?" + params.Encode())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if data.QuotaFinished {
		return "", fmt.Errorf("Quota finished")
	}

	return data.ResponseData.TranslatedText, nil
}
