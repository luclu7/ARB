package API

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

func LaunchBuild(pkg string, server string, client *resty.Client) (RequestResponse, error) {
	body, err := json.Marshal(&BuildRequest{
		PackageName: pkg,
	})
	if err != nil {
		panic(err)
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(body).
		SetResult(RequestResponse{}).
		Post(server + "/build/launch")

	if err != nil {
		panic(err)
	}
	var f RequestResponse
	err = json.Unmarshal(resp.Body(), &f)
	return f, err
}
