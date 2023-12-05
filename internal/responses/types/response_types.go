package response

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type DataResponse struct {
	*BaseResponse
	Data map[string]interface{} `json:"data"`
}

func (r *BaseResponse) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func DefaultBaseResponse() *BaseResponse {
	return &BaseResponse{
		Status: http.StatusOK,
		Msg:    "ok",
	}
}

func DefaultDataResponse() *DataResponse {
	return &DataResponse{
		BaseResponse: DefaultBaseResponse(),
		Data:         map[string]interface{}{},
	}
}
