package img2img

import (
	"encoding/json"
	msg "livin-on-a-platter-api/internal/model/msg/types"
	"net/http"
	"os"
)

func GenerateImg(queuedImgId string) (*msg.BaseMsg, error) {
	apiUrl := os.Getenv("AI_API_URL")
	//body := &msg2.DataMsg{
	//	BaseMsg: msg2.DefaultBaseMsg(),
	//	Data: map[string]interface{}{
	//		imgId:    imgId,
	//		uploadId: uploadId,
	//	},
	//}
	response, err := http.Get(apiUrl + "generate/" + queuedImgId)
	if err != nil {
		return nil, err
	}
	msg := &msg.BaseMsg{}
	json.NewDecoder(response.Body).Decode(&msg)
	return msg, nil
}
