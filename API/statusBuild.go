package API

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

func GetStatusOfBuild(UUID string, server string, client *resty.Client) (Pkg, error) {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(Pkg{}).
		Get(server + "/build/check/" + UUID)

	if err != nil {
		panic(err)
	}
	var f Pkg
	err = json.Unmarshal(resp.Body(), &f)
	return f, err
}
