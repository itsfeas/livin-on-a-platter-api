package Msg

import (
	"encoding/json"
	"net/http"
)

type BaseMsg struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type DataMsg struct {
	*BaseMsg
	Data map[string]interface{} `json:"data"`
}

func (r *BaseMsg) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func (r *DataMsg) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func DefaultBaseMsg() *BaseMsg {
	return &BaseMsg{
		Status: http.StatusOK,
		Msg:    "ok",
	}
}

func DefaultDataMsg() *DataMsg {
	return &DataMsg{
		BaseMsg: DefaultBaseMsg(),
		Data:    map[string]interface{}{},
	}
}
