package speller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	serviceURL = "http://speller.yandex.net/services/spellservice.json/checkText"
)

type Response struct {
	Code        int      `json:"code"`
	Pos         int      `json:"pos"`
	Row         int      `json:"row"`
	Col         int      `json:"col"`
	Len         int      `json:"len"`
	Word        string   `json:"word"`
	Suggestions []string `json:"s"`
}

func CheckText(ctx context.Context, text string) (string, error) {
	values := url.Values{}
	values.Add("text", text)
	resp, err := http.PostForm(serviceURL, values)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprint("speller's response status: ", resp.Status))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response []Response
	if err = json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	ftext := []rune(text)
	for _, v := range response {
		if len(v.Suggestions) > 0 {
			for i, newChar := range []rune(v.Suggestions[0]) {
				ftext[v.Pos+i] = newChar
			}
		}
	}

	return string(ftext), nil
}
