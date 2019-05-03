package main

import (
	"cloud.google.com/go/logging"
	"encoding/json"
	"google.golang.org/appengine"
	"net/http"
)

type MyEntry struct {
	Name  string
	Count int
}

func init() {
	http.HandleFunc("/test", logcheck)
}

func logcheck(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client, err := logging.NewClient(ctx, "gcp-projectname")
	if err != nil {
		return
	}
	defer client.Close()
	logger := client.Logger("output-logname")

	// JSON Payload
	j := []byte(`{"Name": "Bob", "Count": 3}`)
	logger.Log(logging.Entry{Payload: json.RawMessage(j)})

}
