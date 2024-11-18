package exception

const (
	// 웹 라이브라리 에러(WEBSERVER_)
	ErrWebServerInternal = "WEBSERVER_INTERNAL_ERROR" // 웹 서버 내부 에러
	ErrWebServerValidate = "WEBSERVER_VALIDATION"     // 웹 서버 유효성 검사 실패
)
