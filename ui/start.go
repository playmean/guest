package ui

import (
	"net/http"
	"time"
)

func WaitEndpoint(url string) bool {
	retries := 3

	for {
		if retries == 0 {
			retries--

			return false
		}

		time.Sleep(time.Second)

		resp, err := http.Get(url)
		if err != nil {
			continue
		}

		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			continue
		}

		break
	}

	return true
}
