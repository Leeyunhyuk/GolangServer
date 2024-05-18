package types

// network, repository, service에서 사용되는 모든 타입을 정의 내림
type ApiResponse struct {
	Result      int64  `json:"result"`
	Description string `json:"description"`
}

func NewApiResponse(description string, result int64) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
	}
}
