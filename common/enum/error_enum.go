package enum

var (
	DataErr = NewErrorEnum("0001", "data error")
)

func NewErrorEnum(code string, msg string) *ErrorEnum {
	return &ErrorEnum{code, msg}
}

type ErrorEnum struct {
	Code string
	Msg  string
}
