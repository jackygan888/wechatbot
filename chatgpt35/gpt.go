package gpt35

import (
	"bytes"
	"encoding/json"
	"github.com/jackygan888/wechatbot/config"
	"io/ioutil"
	"log"
	"net/http"
)

var Cache CacheInterface

func init() {
	Cache = GetSessionCache()
}

func Completions(session, msg string) (string, error) {

	ms := Cache.GetMsg(session)
	message := &Message{
		Role:    "user",
		Content: msg,
	}
	ms = append(ms, *message)

	requestBody := ChatGPTRequestBody{
		Model:       "gpt-3.5-turbo",
		Messages:    ms,
		MaxTokens:   2048,
		Temperature: 0.7,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	log.Printf("request chatgpt35 json string : %v", string(requestData))
	req, err := http.NewRequest("POST", DefaultUrl, bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	apiKey := config.LoadConfig().ApiKey
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}

	ms = append(ms, gptResponseBody.Choices[0].Message)
	Cache.SetMsg(session, ms)

	var reply string
	reply = gptResponseBody.Choices[0].Message.Content

	log.Printf("gpt response text: %s \n", reply)

	return reply, nil
}
