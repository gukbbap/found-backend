package exception

const (
	// MySQL 관련 에러 (MYSQL_)
	ErrMySQLInternal    = "MYSQL_INTERNAL_ERROR"       // 내부 MySQL 에러
	ErrMySQLDuplicate   = "MYSQL_DUPLICATE_ENTRY"      // 중복 데이터
	ErrMySQLConstraint  = "MYSQL_CONSTRAINT_VIOLATION" // 제약조건 위반
	ErrMySQLNotFound    = "MYSQL_NOT_FOUND"            // 데이터 없음
	ErrMySQLNotLoaded   = "MYSQL_NOT_LOADED"           // 데이터 리로드 실패
	ErrMySQLNotSingular = "MYSQL_NOT_SINGULAR"         // 단일 데이터가 아님
	ErrMySQLValidation  = "MYSQL_VALIDATION_FAILED"    // 데이터 유효성 검사 실패

	// Service 관련 에러 (SERVICE_)
	ErrServiceFailedHashPassword = "SERVICE_FAILED_HASH_PASSWORD"

	// Presenter 관련 에러 (PRESENTER_)
	ErrPresenterFailedBind       = "PRESENTER_FAILED_BIND"
	ErrPresenterFailedValidation = "PRESENTER_FAILED_VALIDATION"
)
