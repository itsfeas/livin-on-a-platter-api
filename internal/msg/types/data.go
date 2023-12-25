package msg

import (
	"encoding/json"
)

type DataMsg struct {
	*BaseMsg
	Data map[string]interface{} `json:"data"`
}

func (r *DataMsg) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func DefaultDataMsg() *DataMsg {
	return &DataMsg{
		BaseMsg: DefaultBaseMsg(),
		Data:    map[string]interface{}{},
	}
}
