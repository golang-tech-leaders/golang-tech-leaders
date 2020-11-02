package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// var _ server.API = (*CurrentTime)(nil)

const apiUrl = "http://worldtimeapi.org/api/timezone/Europe/Moscow"

type CurrentTime struct {
	apiAddr string
}

func New() *CurrentTime {
	return &CurrentTime{
		apiAddr: apiUrl,
	}
}

func (ct *CurrentTime) GetTime() (time.Time, error) {
	resp, err := http.Get(ct.apiAddr)
	if err != nil || resp.StatusCode != http.StatusOK {
		return time.Time{}, fmt.Errorf("something goes wrong")
	}
	var data TimeResponse
	json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return time.Time{}, fmt.Errorf("something goes wrong")
	}
	return data.Timestamp, nil
}

type TimeResponse struct {
	Timestamp time.Time `json:"datetime"`
}
