package alienvault

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestJobRetrieval(t *testing.T) {

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "XSRF-TOKEN=abc123")
		w.Header().Add("Set-Cookie", "SESSION=mysession")

		if strings.HasSuffix(r.RequestURI, "/scheduler") {

			data, err := json.Marshal([]Job{
				{UUID: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", Name: "First Job"},
				{UUID: "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb", Name: "Second Job"},
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write(data)
		}
	}))
	defer ts.Close()

	client := New(strings.Replace(ts.URL, "https://", "", -1), Credentials{
		Username: "something",
		Password: "something",
	})

	err := client.Authenticate()
	if err != nil {
		t.Fatal(err)
	}

	job, err := client.GetJob("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, job.Name, "First Job")

	job, err = client.GetJob("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, job.Name, "Second Job")

}
