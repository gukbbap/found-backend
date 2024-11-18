package exception

import "errors"

type Exception struct {
	Type    string
	Status  int
	Message string
	Data    any

	Err error
}

func Wrap(e error, t string, s int, m string, opts ...ExceptionOption) *Exception {
	exception := &Exception{
		Err:     e,
		Type:    t,
		Status:  s,
		Message: m,
	}

	// 옵션 적용
	for _, opt := range opts {
		opt(exception)
	}

	return exception
}

func New(t string, s int, m string, opts ...ExceptionOption) *Exception {
	exception := &Exception{
		Type:    t,
		Status:  s,
		Message: m,
		Err:     errors.New(t),
	}

	// 옵션 적용
	for _, opt := range opts {
		opt(exception)
	}

	return exception
}

func (e Exception) Error() string {
	return e.Err.Error()
}

func Is(err error, t string) bool {
	exception, ok := err.(*Exception)
	if !ok {
		return false
	}

	return exception.Type == t
}
