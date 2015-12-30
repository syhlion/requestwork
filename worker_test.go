package requestwork

import (
	"net/http"
	"testing"
)

func TestExecute(t *testing.T) {
	req, err := http.NewRequest("GET", "http://tw.yahoo.com", nil)
	if err != nil {
		t.Error("request error: ", err)
	}
	a := New(&http.Client{}, 5)
	resp, err := a.Execute(req)
	if err != nil {
		t.Error("response error: ", err)
	}
	if resp.StatusCode != 200 {
		t.Error("Status error:", resp.StatusCode)
	}

}
