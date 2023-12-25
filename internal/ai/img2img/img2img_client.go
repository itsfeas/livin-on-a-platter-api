package img2img

import (
	msg "livin-on-a-platter-api/internal/msg/types"
	"net/http"
)

func QueueImg(imgId string) {
	body := &msg.DataMsg{
		BaseMsg: msg.DefaultBaseMsg(),
		Data: map[string]interface{}{
			"imgId": imgId,
		},
	}
	http.Get("")
}
