package response

type BaseResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type DataResponse struct {
	*BaseResponse
	Data interface{} `json:"data"`
}
