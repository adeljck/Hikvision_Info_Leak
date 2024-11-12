package modules

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"net"
	"net/url"
	"os"
	"strconv"
)

func LoadFile(Path string) ([]string, error) {
	var lines []string

	// 打开文件
	file, err := os.Open(Path)
	if err != nil {
		return nil, fmt.Errorf("Can Not Open File: %s, ERROR: %v", Path, err)
	}
	defer file.Close() // 确保在函数结束时关闭文件

	// 创建Scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // 将每一行添加到切片中
	}

	// 检查扫描过程中是否发生错误
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Read File Error: %v", err)
	}

	return lines, nil // 返回包含所有行的切片
}
func UrlChecker(target string) (string, string, bool) {
	Schema, err := url.ParseRequestURI(target)
	if err != nil {
		return "", "", false
	}
	return Schema.Scheme + "://" + Schema.Host, Schema.Hostname(), true
}
func IPChecker(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	} else {
		return true
	}
}
func PortChecker(port string) bool {
	p, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	if p <= 0 || p >= 65535 {
		return false
	}
	return true
}
func DecryptData(base string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		return "", err
	}

	// Check magic number
	if decode[0] != 17 {
		return "", fmt.Errorf("Invalid magic number")
	}

	// Extract size and salt
	size := int(decode[1])<<8 | int(decode[2])
	salt := decode[4 : 4+size]

	// Extract IV and data
	iv := decode[20 : 20+16]
	data := decode[36 : 36+size]

	// Derive AES key from password and salt using PBKDF2
	key := pbkdf2.Key([]byte("Abc123@&$++Hik45"), salt, 10000, 32, sha256.New)

	// Create AES cipher in CBC mode with PKCS5Padding
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt data
	mode.CryptBlocks(data, data)

	// Remove padding
	unpad := func(src []byte) []byte {
		length := len(src)
		unpadding := int(src[length-1])
		return src[:(length - unpadding)]
	}

	decryptedData := unpad(data)

	// Convert decrypted data to UTF-8 string
	return string(decryptedData), nil
}
