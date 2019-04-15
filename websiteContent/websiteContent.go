package websiteContent

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WebContent struct {
	Name1      string `json:"fname"`
	Name2      string `json:"sname"`
	Persentage int    `json:"persentage"`
	Result     string `json:"result"`
}

func GetResult(content string) (WebContent, error) {
	webRes := WebContent{}
	err := json.Unmarshal([]byte(content), &webRes)
	return webRes, err

}

func DoRequest(content *WebContent) (*http.Response, error) {
	url := "https://love-calculator.p.rapidapi.com/getPercentage?fname=" + content.Name1 + "&sname=" + content.Name2

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	request.Header.Set("X-RapidAPI-Host", "love-calculator.p.rapidapi.com")
	request.Header.Set("X-RapidAPI-Key", "3e0c85aed6mshc3673dbbc78ec1bp1f62f9jsn3cba8e8a0184")
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return resp, nil
}
