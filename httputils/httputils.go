package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// PostRequest simplifies building post requests that take JSON, JSON must be
// pre-packaged into struct. Header map must be in the same format as when
// using *http.Header.Set, key is the name of the header, and value is the
// value.
func PostRequest(postStruct interface{}, URL string, headers map[string]string) (b []byte, err error) {
	postData, err := json.Marshal(postStruct)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(postData))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// send json with new secrets to vault
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return
}
