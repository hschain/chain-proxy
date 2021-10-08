package utils

import (
	"crypto/tls"

	//"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type setRequest func(*http.Request)

//DoHTTP todo
//func DoHTTP(method string, headers map[string]string, cookies []*http.Cookie, body io.Reader, url string, obj interface{}) error {
func DoHTTP(method string, setFunc setRequest, body io.Reader, url string, obj interface{}) error {

	log.Printf("debuging...: url: %s.\n", url)

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	if setFunc != nil {
		setFunc(request)
	}

	// log.Printf("http url is %s, header is %+v", url, request.Header)
	// log.Printf("request.Body:%v.\n", request.Body)

	tr := &http.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 2 * time.Second}

	resp, err := client.Do(request)

	if err != nil {
		log.Printf("http request failed:%s", err)
		return err
	}

	defer resp.Body.Close()
	defer client.CloseIdleConnections()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body failed:%s", err)
		return err
	}

	// log.Printf("debuging...:\nbuf: %v.\n", string(buf))

	if obj != nil {
		if s, ok := obj.(*string); ok {
			*s = string(buf)
		} else {
			if err := json.Unmarshal(buf, obj); err != nil {
				log.Printf("unmarshal failed:%s", err)
				log.Printf("url is %s", url)
				log.Printf("buf is %s", string(buf))
				return err
			}
		}
	}

	return nil

}
