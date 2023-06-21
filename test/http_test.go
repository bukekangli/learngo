package test

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"
)

const defaultIdleConns = 160

func TestHttpContextCanceledRetry(t *testing.T) {
	//ctx := context.Background()
	ctx, cancelCtxFunc := context.WithCancel(context.Background())
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.baidu.com", nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		IdleConnTimeout:     60 * time.Second,
		MaxIdleConnsPerHost: defaultIdleConns,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	cancelCtxFunc()
	var response *http.Response
	for i := 0; i < 2; i++ {
		response, err = client.Do(request)
		if err == nil {
			break
		}
		t.Logf(err.Error())
		request = request.WithContext(context.Background())
	}
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf(string(body))
}
