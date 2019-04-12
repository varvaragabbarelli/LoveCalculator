package web

import (
	"encoding/json"
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
