package code

import "ldm/common/constant"

//允许哪些自定义头信息
func allowHeader(s string) (string, bool) {
	if _, ok := constant.MAP_ALLOW_ENDPOINT_HEADER[s]; ok {
		return s, true
	}
	return "", false
}
