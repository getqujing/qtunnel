package tunnel

import (
    "log"
    "crypto/rc4"
    "crypto/sha256"
    "crypto/cipher"
)

type Cipher struct {
    enc cipher.Stream
    dec cipher.Stream
}

type chiperCreator func(key []byte) (*Cipher, error)

var cipherMap = map[string]chiperCreator {
    "rc4": newRC4Cipher,
}

func secretToKey(secret []byte) []byte {
    h := sha256.New()
    return h.Sum(secret)
}

func newRC4Cipher(key []byte) (*Cipher, error) {
    c, err := rc4.NewCipher(key)
    if err != nil {
        return nil, err
    }
    c2 := *c

    return &Cipher{c, &c2}, nil
}

func NewCipher(cryptoMethod string, secret []byte) *Cipher {
    cc := cipherMap[cryptoMethod]
    if cc == nil {
        log.Fatalf("unsupported crypto method %s", cryptoMethod)
    }
    c, err := cc(secretToKey(secret))
    if err != nil {
        log.Fatal(err)
    }
    return c
}

func (c *Cipher) encrypt(dst, src []byte) {
    c.enc.XORKeyStream(dst, src)
}

func (c *Cipher) decrypt(dst, src []byte) {
    c.dec.XORKeyStream(dst, src)
}
