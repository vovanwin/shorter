package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/vovanwin/shorter/internal/app/config"
)

type Token struct {
	config.Config
}

func (t Token) CreateUserId() ([]byte, error) {
	return uuid.New().MarshalText()
}

func (t Token) Encode(userId []byte) (string, error) {
	// будем использовать AES256, создав ключ длиной 32 байта
	key := t.GetConfig().Key // ключ шифрования

	aesblock, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	// создаём вектор инициализации
	nonce := []byte(key[len(key)-aesgcm.NonceSize():])

	dst := aesgcm.Seal(nil, nonce, userId, nil) // зашифровываем
	str := hex.EncodeToString(dst)
	return str, nil
}

func (t Token) Decode(cookie string) ([]byte, error) {
	// будем использовать AES256, создав ключ длиной 32 байта
	key := t.GetConfig().Key // ключ шифрования

	aesblock, err := aes.NewCipher([]byte(key))
	if err != nil {
		return []byte(""), err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return []byte(""), err
	}

	// создаём вектор инициализации
	nonce := []byte(key[len(key)-aesgcm.NonceSize():])

	user, _ := hex.DecodeString(cookie)

	src2, err := aesgcm.Open(nil, nonce, user, nil) // расшифровываем
	if err != nil {
		return []byte(""), err
	}
	return src2, nil
}
