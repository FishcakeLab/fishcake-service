package enum

var (
	DataErr  = NewErrorEnum("0001", "data error")
	GrpcErr  = NewErrorEnum("0002", "grpc error")
	ParamErr = NewErrorEnum("0003", "query param error")
)

func NewErrorEnum(code string, msg string) *ErrorEnum {
	return &ErrorEnum{code, msg}
}

type ErrorEnum struct {
	Code string
	Msg  string
}
