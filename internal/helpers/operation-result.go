package helpers

type OperationResult struct {
	Error     string      `json:"error,omitempty"`
	IsSuccess bool        `json:"isSuccess"`
	Data      interface{} `json:"data,omitempty"`
}

// Các hàm khởi tạo nhanh
func NewOperationResultSuccess(data interface{}) OperationResult {
	return OperationResult{
		IsSuccess: true,
		Data:      data,
	}
}

func NewOperationResultError(err string) OperationResult {
	return OperationResult{
		IsSuccess: false,
		Error:     err,
	}
}

func NewOperationResult(isSuccess bool, err string, data interface{}) OperationResult {
	return OperationResult{
		IsSuccess: isSuccess,
		Error:     err,
		Data:      data,
	}
}
