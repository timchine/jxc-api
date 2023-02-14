package http

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			DualStack: true,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   30 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   5,
		MaxConnsPerHost:       5,
		IdleConnTimeout:       30 * time.Second,
		ResponseHeaderTimeout: 30 * time.Second,
		ExpectContinueTimeout: 30 * time.Second,
	},
	Timeout: time.Minute,
}

func Get(url string, scan ...interface{}) (*http.Response, error) {
	res, err := client.Get(url)
	if scan != nil && err == nil {
		if res != nil && res.StatusCode == http.StatusOK {
			defer res.Body.Close()
			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return res, err
			}
			err = scanV(b, scan...)
			if err != nil {
				return res, err
			} else {
				return nil, nil
			}
		}
	}
	return res, err
}

func scanV(b []byte, scan ...interface{}) error {
	for _, v := range scan {
		err := json.Unmarshal(b, v)
		if err != nil {
			return err
		}
	}
	return nil
}
