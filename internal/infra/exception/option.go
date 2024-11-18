package exception

type Map map[string]any

type ExceptionOption func(*Exception)

func WithData(data any) ExceptionOption {
	return func(e *Exception) {
		e.Data = data
	}
}
