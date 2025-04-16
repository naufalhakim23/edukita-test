package payload

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error,omitempty"`
}

type BaseResponseWithMeta struct {
	BaseResponse
	Meta MetaResponse `json:"meta"`
}

type MetaResponse struct {
	TotalCount int `json:"total_count"`
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
}
