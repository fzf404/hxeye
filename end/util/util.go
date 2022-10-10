package util

import "regexp"

// VerifyEmailFormat 判断电子邮箱
func VerifyEmailFormat(email string) bool {
	// 匹配电子邮箱
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
