package apicalls

import (
	"encoding/json"
	"net/http"

	"github.com/lancer2672/DandelionServer_Go/constants"
)

type AuthApi struct {
}
type Role struct {
	Name        string `json:"name"`
	Permissions string `json:"permissions"`
	Child       string `json:"child"`
}

type Permission struct {
	Read   []string `json:"read"`
	Write  []string `json:"write"`
	Delete []string `json:"delete"`
}

func CheckApiKey(apikey string) (*Role, *Permission, error) {
	resp, err := http.Get(constants.AUTH_PATH + "checkapikey?apikey=" + apikey)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Role       Role
		Permission Permission
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, nil, err
	}

	return &result.Role, &result.Permission, nil
}
