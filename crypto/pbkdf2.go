package crypto

import (
	"golang.org/x/crypto/pbkdf2"
	"encoding/base64"
	"strconv"
	"crypto/sha256"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/5/7.
// GeneratePasswd 生成与Django兼容的加密字符串
func GeneratePasswd(password, userSalt string, iterations int) (cryptoPasswd string) {
	pwd := []byte(password)  // 用户设置的原始密码
	salt := []byte(userSalt) // 盐，是一个随机字符串，每一个用户都不一样，在这里我们随机选择 "I1lrI7wqJOJZ" 作为盐
	digest := sha256.New     // digest 算法，使用 sha256

	// 第一步：使用 pbkdf2 算法加密
	dk := pbkdf2.Key(pwd, salt, iterations, 32, digest)

	// 第二步：Base64 编码
	str := base64.StdEncoding.EncodeToString(dk)

	// 第三步：组合加密算法、迭代次数、盐、密码和分割符号 "$"
	return "pbkdf2_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + string(salt) + "$" + str
}
