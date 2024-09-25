package result

type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Ok() Result {
	return Result{
		Success: true,
	}
}
func Data(data interface{}) Result {
	return Result{
		Success: true,
		Data:    data,
	}
}
func Error(err error) Result {
	return Result{
		Success: false,
		Message: err.Error(),
	}
}

func ErrorWithMessage(message string) Result {
	return Result{
		Success: false,
		Message: message, // 直接使用传入的字符串
	}
}
