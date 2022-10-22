package grpc_err

import (
	"encoding/json"
)
type selfDefineErr struct {
	Code int `json:"code"`
	Message  string `json:"message"`
}
func (s *selfDefineErr)Error() string{
	d := selfDefineErr{
		Code: s.Code,
		Message: s.Message,
	}
	b,_ := json.Marshal(d)
	return string(b)
}
func GrpcErr(status int,msg string)error{
	return &selfDefineErr{
		Code: status,
		Message: msg,
	}
}
