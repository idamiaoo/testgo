package aes

import (
	"os"
	"testing"

	"github.com/Luzifer/go-openssl/v4"
)

func TestEnc(t *testing.T) {
	file, err := os.ReadFile("/Users/user/go/src/testgo/stdlib/crypto/aes/localhost.key")
	if err != nil {
		t.Fatal(err)
	}

	o := openssl.New()
	data, err := o.EncryptBytes("123457", file, openssl.BytesToKeyMD5)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("/Users/user/go/src/testgo/stdlib/crypto/aes/localhost.key.enc.2", data, 0o644)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDec(t *testing.T) {

	file, err := os.ReadFile("/Users/user/go/src/testgo/stdlib/crypto/aes/localhost.key.enc.1")
	if err != nil {
		t.Fatal(err)
	}

	o := openssl.New()
	data, err := o.DecryptBytes("123456", file, openssl.BytesToKeyMD5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
