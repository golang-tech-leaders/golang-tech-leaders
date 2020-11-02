package server_test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	"week5/internal/server"
)

type MockAPI struct {
	now time.Time
}

func (m *MockAPI) GetTime() (time.Time, error) {
	return m.now, nil
}

func TestServer(t *testing.T) {
	now := time.Now()
	api := MockAPI{now: now}
	addr := ":18080"
	s := server.New(addr, &api)
	s.Start()
	time.Sleep(time.Second)
	defer s.Stop()
	t.Log(now.Format(time.RFC822))

	resp, err := http.Get("http://localhost:18080/")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	want := "current time is " + now.Format(time.RFC822)
	if txt := string(data); txt != want {
		t.Errorf("want %s got %s", want, txt)
	}
}
