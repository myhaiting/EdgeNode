// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package requests

import (
	"bytes"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

type TestRequest struct {
	req      *http.Request
	BodyData []byte
}

func NewTestRequest(raw *http.Request) *TestRequest {
	return &TestRequest{
		req: raw,
	}
}

func (this *TestRequest) WAFSetCacheBody(bodyData []byte) {
	this.BodyData = bodyData
}

func (this *TestRequest) WAFGetCacheBody() []byte {
	return this.BodyData
}

func (this *TestRequest) WAFRaw() *http.Request {
	return this.req
}

func (this *TestRequest) WAFRemoteAddr() string {
	return this.req.RemoteAddr
}

func (this *TestRequest) WAFRemoteIP() string {
	host, _, err := net.SplitHostPort(this.req.RemoteAddr)
	if err != nil {
		return this.req.RemoteAddr
	} else {
		return host
	}
}

func (this *TestRequest) WAFReadBody(max int64) (data []byte, err error) {
	if this.req.ContentLength > 0 {
		data, err = ioutil.ReadAll(io.LimitReader(this.req.Body, max))
	}
	return
}

func (this *TestRequest) WAFRestoreBody(data []byte) {
	if len(data) > 0 {
		rawReader := bytes.NewBuffer(data)
		buf := make([]byte, 1024)
		_, _ = io.CopyBuffer(rawReader, this.req.Body, buf)
		this.req.Body = ioutil.NopCloser(rawReader)
	}
}

func (this *TestRequest) WAFServerId() int64 {
	return 0
}

func (this *TestRequest) Format(s string) string {
	return s
}