package pkg

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "github.com/sirupsen/logrus"
)

const (
    REQ_KEY = "AKGFBN2KU89EWU10"
    REQ_IV  = "J56HGMLWH28HDLQR"
)

var (
    CodeSaltLen = int(3)
)

// ReqDecode decrypt
func DecodeSecurity(src string) ([]byte, error) {
    return decodeAesBase64(src)
}

//RespEncrypt encrypt
func EncryptSecurity(src string) (string, error) {
    return encodeAesBase64(src, Key.Key, Key.Iv)
}

//decodeAesBase64 base64+aes128 decrypt
func decodeAesBase64(src string) ([]byte, error) {
    //base64
    s1, err := base64.URLEncoding.DecodeString(src)
    if err != nil {
        logrus.Error("decode base64 error : ", err)
        return nil, err
    }

    //aes128
    out, err := decodeAesCbc(s1)
    if err != nil {
        logrus.Error("decode aes cbc error : ", err)
        return nil, err
    }

    return out, nil
}

//encodeAesBase64 base64+aes128 encrypt
func encodeAesBase64(src, key, iv string) (string, error) {
    //aes128
    s1, err := encodeAesCbc(src, key, iv)
    if err != nil {
        logrus.Error("encode aes cbc error : ", err)
        return "", err
    }

    //base64
    out := base64.URLEncoding.EncodeToString(s1)

    return out, nil
}

// encodeAesCbc aes cbc encrypt
func encodeAesCbc(src, key, iv string) ([]byte, error) {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        logrus.Error("aes decode err : ", err)
        return nil, err
    }

    blockSize := block.BlockSize()
    oldData := PKCS5Padding([]byte(src), blockSize)
    blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
    newData := make([]byte, len(oldData))
    blockMode.CryptBlocks(newData, oldData)
    return newData, nil
}

// decodeAesCbc aes cbc decrypt
func decodeAesCbc(src []byte) ([]byte, error) {
    block, err := aes.NewCipher([]byte(REQ_KEY))
    if err != nil {
        logrus.Error("aes decode err : ", err)
        return nil, err
    }

    blockMode := cipher.NewCBCDecrypter(block, []byte(REQ_IV))
    dst := make([]byte, len(src))
    blockMode.CryptBlocks(dst, src)
    out := KCS5UnPadding(dst)

    if len(out) > CodeSaltLen {
        out = out[CodeSaltLen:]
    }

    return out, nil
}

func KCS5UnPadding(src []byte) []byte {
    length := len(src)
    if length <= 0 {
        logrus.Error("PKCS5UnPadding error:", string(src))
        return src
    }
    return src[:(length - int(src[length-1]))]
}

func PKCS5Padding(in []byte, blockSize int) []byte {
    padding := blockSize - len(in) % blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(in, padtext...)
}