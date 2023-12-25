package img2img

import (
	msg2 "livin-on-a-platter-api/internal/model/msg/types"
	"net/http"
)

func QueueImg(imgId string) {
	body := &msg2.DataMsg{
		BaseMsg: msg2.DefaultBaseMsg(),
		Data: map[string]interface{}{
			"imgId": imgId,
		},
	}
	http.Get("")
}
