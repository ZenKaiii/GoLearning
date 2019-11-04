package reall

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgner string
	Timeout time.Duration
}

func (r *Retriever) Get(url string) string {
	resp,err := http.Get(url)
	if err != nil{
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil{
		panic(err)
	}
	return string(result)
}