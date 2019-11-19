package unionpay

import (
	"encoding/base64"
	"math/rand"
	"time"
)

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}

	rand_r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		result = append(result, bytes[rand_r.Intn(len(bytes))])
	}
	return string(result)
}

var Base64Base *base64.Encoding = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

func Base64Encode(data []byte) string {
	return Base64Base.EncodeToString(data)
}

// Base64Decode 仅仅用于测试提交2
func Base64Decode(data string) []byte {
	v, err := Base64Base.DecodeString(data)
	if err != nil {
		return []byte("")
	}
	return v
}
