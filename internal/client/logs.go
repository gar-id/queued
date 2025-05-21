package client

import (
	"fmt"

	"github.com/gar-id/queued/internal/general/config/caches"
	"github.com/gar-id/queued/tools"
	"github.com/r3labs/sse"
)

func QueuedLogs(processName string) {
	// Hit endpoint
	urlRequest := fmt.Sprintf(
		"http://%v/api/v1/queued/program/logs?processName=%v",
		tools.DefaultString(caches.MainConfig.QueueD.API.HTTPListen, "127.0.0.1:3000"),
		processName,
	)

	client := sse.NewClient(urlRequest)
	client.URL = urlRequest
	client.Headers = make(map[string]string)
	client.Headers["User-Agent"] = userAgent

	client.Subscribe("", func(msg *sse.Event) {
		fmt.Println(string(msg.Data))
	})
}
