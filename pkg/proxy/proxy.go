package proxy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	tracingmodel "github.com/devMiguelFerrer/EasyMonitoring/internal/models/tracing_model"
)

type dbProxy interface {
	Save(i interface{})
}

var db dbProxy
var start time.Time
var reqBody []byte

func Create(rawUrl string, rawPort int, dbproxy dbProxy) {
	db = dbproxy
	port := fmt.Sprintf(":%d", rawPort)
	remote, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start = time.Now()
		reqBody, _ = ioutil.ReadAll(r.Body)
		if len(reqBody) == 0 {
			reqBody = []byte(`{message: "Request has NOT body"}`)
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		p.ModifyResponse = handleResponse
		log.Println(r.URL)
		p.ServeHTTP(w, r)
	}
}

func handleResponse(r *http.Response) error {
	resBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		resBodyBytes = []byte(`{message: "Response has NOT body"}`)
	}

	elpased := int(time.Since(start).Milliseconds())
	rawTracing := tracingmodel.TracingModel{
		Method:       r.Request.Method,
		Url:          r.Request.URL.Path,
		CreatedAt:    time.Now().String(),
		StatusCode:   r.StatusCode,
		ResponseBody: string(resBodyBytes),
		RequestBody:  string(reqBody),
		ResponseTime: elpased,
	}

	fmt.Printf("%s %s %dms\n", r.Request.Method, r.Request.URL, elpased)
	db.Save(rawTracing)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(resBodyBytes))
	return nil
}
