package utils

import (
    "fmt"
    "github.com/satori/go.uuid"
)

// GenerateRandomString 固定长度的随机字符串
func GenerateRandomString(len int) string{
    u, _ := uuid.NewV4()
    var s = fmt.Sprintf("%s", u)
    return s[0:len]
}