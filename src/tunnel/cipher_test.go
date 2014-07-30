package tunnel

import (
    "testing"
)

func TestRC4(t *testing.T) {
    secret := []byte("testsecret")
    clearText := "thisISaCLEARtext"
    c := NewCipher("rc4", secret)
    c2 := NewCipher("rc4", secret)
    dst := make([]byte, len(clearText))
    dst2 := make([]byte, len(clearText))
    c.encrypt(dst, []byte(clearText))
    c2.decrypt(dst2, dst)
    if (clearText != string(dst2)) {
        t.Error(string(dst2))
    }
}
