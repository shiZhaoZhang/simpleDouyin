package main

import (
	"douyin/src/service"
	"testing"
)

var passwd string = "123456"
var salt string = "zsz"

func BenchmarkEncruption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		service.Encryption(salt + passwd)
	}
}

func BenchmarkEncruptionPBKDF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		service.Encryption_PBKDF2(passwd, salt, 100)
	}
}
