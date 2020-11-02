package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_greet(t *testing.T) {

	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	curr := time.Now()

	timeMock := func() time.Time {
		return curr
	}

	t.Run("time test", func(t *testing.T) {
		timeHandler(timeMock)(w, r)
	})

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("code want %d got %d", http.StatusOK, res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error is not nil: %v", err)
	}

	want := curr.Format(time.RFC1123)

	if data := string(body); data != want {
		t.Errorf("response want %s got %s", want, data)
	}

}
