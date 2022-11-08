package jwt

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"ldm/common/config"
	"time"
)

type UserInfo struct {
	Name   string `json:"name"`
	UserId int64  `json:"user_id"`
}

//生成token
func GenJwtToken(userInfo UserInfo) (token string, err error) {
	cfg := config.GlobalConfig.Jwt
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":    userInfo.Name,
		"user_id": userInfo.UserId,
		"exp":     time.Now().Add(time.Second * time.Duration(cfg.Expire)).Unix(), //过期时间
		"iss":     cfg.Issuer,
	})
	token, err = t.SignedString([]byte(cfg.SignKey))
	if err != nil {
		return "", err
	}
	return
}

//解释token
func ParseJwtToken(token string) (UserInfo, error) {
	t, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GlobalConfig.Jwt.SignKey), nil
	})
	if err != nil {
		return UserInfo{}, err
	}
	d, err := json.Marshal(t.Claims)
	if err != nil {
		return UserInfo{}, err
	}
	var userInfo UserInfo
	if err = json.Unmarshal(d, &userInfo); err != nil {
		return UserInfo{}, err
	}
	return userInfo, nil
}
