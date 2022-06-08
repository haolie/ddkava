package model

type HResponse struct {
	ErrorStr string
	Content  map[string]interface{}
}

func NewHResponse(errStr string) *HResponse {
	return &HResponse{
		ErrorStr: "",
		Content:  make(map[string]interface{}, 2),
	}
}
