package img2img

// APIRequest represents the parameters for an API request.
type APIRequest struct {
	Key               string  `json:"key"`
	Prompt            string  `json:"prompt"`
	NegativePrompt    string  `json:"negative_prompt"`
	InitImage         string  `json:"init_image"`
	Width             string  `json:"width"`
	Height            string  `json:"height"`
	Samples           int     `json:"samples"`
	NumInferenceSteps int     `json:"num_inference_steps"`
	SafetyChecker     string  `json:"safety_checker"`
	EnhancePrompt     string  `json:"enhance_prompt"`
	GuidanceScale     int     `json:"guidance_scale"`
	Strength          float64 `json:"strength"`
	Seed              int32   `json:"seed"`
	Base64            string  `json:"base64"`
	Webhook           string  `json:"webhook"`
	TrackID           string  `json:"track_id"`
}
