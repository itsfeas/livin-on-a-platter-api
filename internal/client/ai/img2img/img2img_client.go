package img2img

import (
	"net/http"
	"os"
)

func GenerateImg(queuedImgId string) error {
	apiUrl := os.Getenv("AI_API_URL")
	//body := &msg2.DataMsg{
	//	BaseMsg: msg2.DefaultBaseMsg(),
	//	Data: map[string]interface{}{
	//		imgId:    imgId,
	//		uploadId: uploadId,
	//	},
	//}
	if _, err := http.Get(apiUrl + "/queue/" + queuedImgId); err != nil {
		return err
	}
	return nil
}
