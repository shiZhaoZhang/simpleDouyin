package service

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

//密码加密，生成哈希值，用转化为十六进制，并用string类型进行保存
func Encryption(passwd string) string {
	h := sha256.New()
	h.Write([]byte(passwd))
	return fmt.Sprintf("%x", h.Sum(nil))

}

func Encryption_PBKDF2(passwd, salt string, iter int) string {
	return fmt.Sprintf("%x", pbkdf2.Key([]byte(passwd), []byte(salt), iter, 32, sha256.New))
}
