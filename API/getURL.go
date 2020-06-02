package API

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

func GetURLs(UUID string, server string, client *resty.Client) ([]File, error) {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(File{}).
		Get(server + "/build/getURL/" + UUID)

	if err != nil {
		panic(err)
	}
	var f []File
	err = json.Unmarshal(resp.Body(), &f)
	return f, err
}
