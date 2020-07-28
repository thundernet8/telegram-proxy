package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type requestPayloadStruct struct {
	Text string `json:"text"`
	ChatId string `json:"chat_id"`
}

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host
	proxy.ServeHTTP(res, req)
}

func requestBodyDecoder(request *http.Request) *json.Decoder {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		panic(err)
	}

	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

func parseRequestBody(request *http.Request) requestPayloadStruct {
	decoder := requestBodyDecoder(request)
	var requestPayload requestPayloadStruct
	err := decoder.Decode(&requestPayload)

	if err != nil {
		panic(err)
	}

	return requestPayload
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	requestPayload := parseRequestBody(req)
	url := "https://api.telegram.org"

	log.Printf("proxy_url: %s\n", req.URL.Path)
	log.Printf("send msg: %s\n", requestPayload.Text)

	serveReverseProxy(url, res, req)
}

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}