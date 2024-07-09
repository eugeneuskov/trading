package services

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpClientService struct {
	client *resty.Client
}

func NewHttpClientService() *HttpClientService {
	return &HttpClientService{
		client: resty.New(),
	}
}

type errorMessage struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (h *HttpClientService) Get(url string, header http.Header, query url.Values) ([]byte, error) {
	resp, err := h.client.R().
		SetQueryString(query.Encode()).
		SetHeaders(headers(header)).
		Get(url)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.RawBody())

	if resp.IsError() {
		var errMessage errorMessage
		err = json.Unmarshal(resp.Body(), &errMessage)
		if err != nil {
			return nil, errors.New(resp.Status())
		}
		return nil, errors.New(errMessage.Msg)
	}

	return resp.Body(), nil
}

func (h *HttpClientService) Post(url string, header http.Header, params []byte) ([]byte, error) {
	resp, err := h.client.R().
		SetBody(params).
		SetHeaders(headers(header)).
		Post(url)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.RawBody())

	if resp.IsError() {
		var errMessage errorMessage
		err = json.Unmarshal(resp.Body(), &errMessage)
		if err != nil {
			return nil, errors.New(resp.Status())
		}
		return nil, errors.New(errMessage.Msg)
	}

	return resp.Body(), nil
}

func (h *HttpClientService) Delete(url string, header http.Header, query url.Values) ([]byte, error) {
	resp, err := h.client.R().
		SetQueryString(query.Encode()).
		SetHeaders(headers(header)).
		Delete(url)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.RawBody())

	if resp.IsError() {
		var errMessage errorMessage
		err = json.Unmarshal(resp.Body(), &errMessage)
		if err != nil {
			return nil, errors.New(resp.Status())
		}
		return nil, errors.New(errMessage.Msg)
	}

	return resp.Body(), nil
}

func headers(header http.Header) map[string]string {
	h := make(map[string]string)

	for name, values := range header {
		h[name] = strings.Join(values, ";")
	}

	return h
}
