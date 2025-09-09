package error

type ErrorCode int

const (
	SUCCESS ErrorCode = 1 // 成功固定代碼

	UNKNOWN_ERROR       ErrorCode = 10000 + iota // 未知錯誤
	INVALID_PARAMS                               // 參數錯誤
	NOT_FOUND                                    // 資源未找到
	UNAUTHORIZED                                 // 權限不足
	TIMEOUT                                      // 請求超時
	SERVICE_UNAVAILABLE                          // 服務不可用
)

func (e ErrorCode) String() string {
	switch e {
	case SUCCESS:
		return "Success"
	case UNKNOWN_ERROR:
		return "Unknown error"
	case INVALID_PARAMS:
		return "Invalid parameters"
	case NOT_FOUND:
		return "Resource not found"
	case UNAUTHORIZED:
		return "Unauthorized"
	case TIMEOUT:
		return "Timeout"
	case SERVICE_UNAVAILABLE:
		return "Service unavailable"
	default:
		return "Undefined error"
	}
}
