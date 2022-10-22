package grpc_err

import (
	"encoding/json"
)
type SelfDefineErr struct {
	Code int `json:"code"`
	Message  string `json:"message"`
}
func (s *SelfDefineErr)Error() string{
	d := SelfDefineErr{
		Code: s.Code,
		Message: s.Message,
	}
	b,_ := json.Marshal(d)
	return string(b)
}
func GrpcErr(status int,msg string)error{
	return &SelfDefineErr{
		Code: status,
		Message: msg,
	}
}
