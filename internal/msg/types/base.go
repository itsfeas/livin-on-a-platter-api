package msg

import (
	"encoding/json"
	"net/http"
)

type BaseMsg struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func (r *BaseMsg) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func DefaultBaseMsg() *BaseMsg {
	return &BaseMsg{
		Status: http.StatusOK,
		Msg:    "ok",
	}
}
