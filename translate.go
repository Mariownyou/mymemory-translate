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

type Config struct {
	Text  string
	From  string
	To    string
	Email string // 10x more translations if you provide an email
}

func Translate(c Config) (string, error) {
	params := url.Values{}
	params.Add("q", c.Text)
	if c.Email != "" {
		params.Add("de", c.Email)
	}
	params.Add("langpair", fmt.Sprintf("%s|%s", c.From, c.To))

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
