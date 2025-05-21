package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/r3labs/sse"

	"github.com/gar-id/queued/internal/general/config/caches"
	"github.com/gar-id/queued/tools"
)

const jsonApp string = "application/json"
const userAgent string = "Golang QueueD Control CLI"

func PostAPIServer(jsonPayload []byte, endpoint string) []byte {
	// Register host into QueueD Server
	urlRequest := fmt.Sprintf(
		"http://%v/api/v1/queued/%v",
		tools.DefaultString(caches.MainConfig.QueueD.API.HTTPListen, "127.0.0.1:3000"),
		endpoint,
	)
	apiRequest, _ := http.NewRequest("POST", urlRequest, bytes.NewBuffer(jsonPayload))
	apiRequest.Header.Add("accept", jsonApp)
	apiRequest.Header.Add("User-Agent", userAgent)
	apiRequest.Header.Add("content-type", jsonApp)

	res, err := http.DefaultClient.Do(apiRequest)
	if err != nil {
		tools.ZapLogger("both").Fatal(fmt.Sprintf("Cannot connect to API. %v", err))
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

func GetAPIServer(endpoint string) []byte {
	// Register host into QueueD Server
	urlRequest := fmt.Sprintf(
		"http://%v/api/v1/queued/%v",
		tools.DefaultString(caches.MainConfig.QueueD.API.HTTPListen, "127.0.0.1:3000"),
		endpoint,
	)
	apiRequest, _ := http.NewRequest("GET", urlRequest, nil)
	apiRequest.Header.Add("User-Agent", userAgent)
	apiRequest.Header.Add("content-type", jsonApp)

	res, err := http.DefaultClient.Do(apiRequest)
	if err != nil {
		tools.ZapLogger("both").Fatal(fmt.Sprintf("Cannot connect to API. %v", err))
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

func SSEClient(endpoint string) {
	// Register host into QueueD Server
	urlRequest := fmt.Sprintf(
		"http://%v/api/v1/queued/%v",
		tools.DefaultString(caches.MainConfig.QueueD.API.HTTPListen, "127.0.0.1:3000"),
		endpoint,
	)

	client := sse.NewClient(urlRequest)
	client.Headers = make(map[string]string)
	client.Headers["User-Agent"] = userAgent

	client.Subscribe("", func(msg *sse.Event) {
		fmt.Println(msg.Data)
	})

}
