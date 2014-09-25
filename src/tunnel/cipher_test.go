package tunnel

import (
    "testing"
)

func TestRC4(t *testing.T) {
    secret := []byte("testsecret")
    clearText := "thisISaCLEARtext"
    c := NewCipher("rc4", secret)
    dst := make([]byte, len(clearText))
    dst2 := make([]byte, len(clearText))
    c.encrypt(dst, []byte(clearText))
    c.decrypt(dst2, dst)
    if (clearText != string(dst2)) {
        t.Error(string(dst2))
    }
}

func TestAES256CFB(t *testing.T) {
    secret := []byte("testsecret")
    clearText := "thisISaCLEARtext"
    c := NewCipher("aes256cfb", secret)
    dst := make([]byte, len(clearText))
    dst2 := make([]byte, len(clearText))
    c.encrypt(dst, []byte(clearText))
    c.decrypt(dst2, dst)
    if (clearText != string(dst2)) {
        t.Error(string(dst2))
    }
}
