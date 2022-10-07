package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	lggr "../logger"

	"golang.org/x/crypto/bcrypt"
	rc4 "gopkg.in/goyy/goyy.v0/util/crypto/rc4"
)

var logger = lggr.Log()

func ToInt(str string) int {
	if str == "" {
		return 0
	}

	intFromStr, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}

	return intFromStr
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ToBool(str string) bool {
	if str == "" {
		return false
	}

	strBooled, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println(err)
	}

	return strBooled
}

func GetTimeStamp() int64 {
	return time.Now().Unix()
}

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

// инициализация ГСЧ
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenRandStr - генерация строки из n символов
func GenRandStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func EncryptData(data string) string {
	encKey := GenRandStr(16)
	data_rc4, err := rc4.Encrypt([]byte(data), []byte(encKey))
	if err != nil {
		logger.Println("utils.EncryptData: ", err)
	}

	encrypted_data := base64.StdEncoding.EncodeToString(data_rc4)

	return encKey + encrypted_data
}

func DecodeData(data_enc string) string {
	dataLen := len(data_enc)
	if dataLen <= 16 {
		return ""
	}

	encKey := data_enc[0:16] // ключ длинной 16 символов

	// заменяем все пробелы +,обрезаем ключ
	data_base64 := strings.ReplaceAll(data_enc[16:dataLen], " ", "+")
	data_rc4, err := base64.StdEncoding.DecodeString(data_base64)
	if err != nil {
		logger.Println("utils.DecodeData: ", err)
	}

	data_decrypted, err := rc4.Decrypt(data_rc4, []byte(encKey))
	if err != nil {
		logger.Println("utils.DecodeData: ", err)
	}

	return string(data_decrypted)
}

func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		logger.Println("utils.HashAndSalt:", err)
	}
	return string(hash)
}

func HashCompare(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		return false
	}

	return true
}

func DelSpace(str string) string {
	return strings.ReplaceAll(str, " ", "")
}
