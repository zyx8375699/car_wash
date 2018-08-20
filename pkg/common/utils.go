package common

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	ERR_TYPE_ASSERTION = errors.New("type assertion error")
)

type CommonResponse struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Object  interface{} `json:"obj"`
}

type LoginConfig struct {
	UserName string `json:"user"`
	Password string `json:"password"`
}

func NewCommonResponse(ok bool, err error, obj interface{}) *CommonResponse {
	if ok {
		if err != nil {
			return &CommonResponse{
				Success: false,
				Msg:     err.Error(),
				Object:  nil,
			}
		} else {
			return &CommonResponse{
				Success: true,
				Msg:     "",
				Object:  obj,
			}
		}
	}
	return &CommonResponse{
		Success: false,
		Msg:     ERR_TYPE_ASSERTION.Error(),
		Object:  nil,
	}
}

func EncodeResp(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func GetResponseBodyByPost(url string, params []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
