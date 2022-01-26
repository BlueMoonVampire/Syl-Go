package Vanitas-Go

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type USER struct {
	Blacklisted string                 `json:"blacklisted"`
        Enforcer    string                 `json:"enforcer"`
        Message     string                 `json:"message"`
        Reason      string                 `json:"reason"`
        User        string                 `json:"user"`
}


func vanitas(UserId string) (*USER, error) {
	resp, err := http.Get(fmt.Sprintf("https://vanitas-api.up.railway.app/user/%v", UserId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r = new(Response)
	json.NewDecoder(resp.Body).Decode(&r)
	return &r.Data.USER, nil
