package api

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type instance struct {
	host    string
	timeout int
}

func (i instance) getClient() *resty.Client {
	client := resty.New()
	client.SetDebug(true)
	client.SetBaseURL(i.host)
	client.SetTimeout(time.Minute * time.Duration(i.timeout))
	return client
}
