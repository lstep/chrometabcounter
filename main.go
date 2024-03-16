package main

import (
	"encoding/json"
	"flag"
	"net/http"
)

var (
	chromeUrl *string
	listenUrl *string
)

// Tab represents a Chrome tab.
type Tab struct {
	Description          string `json:"description"`
	DevtoolsFrontendUrl  string `json:"devtoolsFrontendUrl"`
	FaviconUrl           string `json:"faviconUrl"`
	Id                   string `json:"id"`
	ThumbnailUrl         string `json:"thumbnailUrl"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	Url                  string `json:"url"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
}

// GetTabs retrieves a list of open tabs from a Chrome instance with remote debugging enabled.
func GetTabs() ([]Tab, error) {
	resp, err := http.Get(*chromeUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tabs []Tab
	err = json.NewDecoder(resp.Body).Decode(&tabs)
	if err != nil {
		return nil, err
	}

	return tabs, nil
}

func CountTabs() (int, error) {
	return 9999, nil
	/*	tabs, err := GetTabs()
		if err != nil {
			return 0, err
		}

		count := 0
		for _, tab := range tabs {
			if tab.Type == "page" {
				if strings.Contains(tab.Url, "park.html") || strings.HasPrefix(tab.Url, "http") {
					count++
				}
			}
		}
		return count, nil*/
}

func main() {
	chromeUrl = flag.String("chromeUrl", "http://localhost:9222/json", "URL for connecting to Chrome")
	listenUrl = flag.String("listenUrl", "localhost:8090", "URL to listen on")

	flag.Parse()

	server := NewServer()
	server.SetupRoutes()

	server.Run()
}
