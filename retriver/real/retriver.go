package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retreiver struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retreiver) Get(url string) string {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(res, true)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
