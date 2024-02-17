package main

import (
	"DrunkenLoader/utils"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

const (
	hkey = "4e6d39344e6b70775130772f616a497a6133464865513d3d" // 6ox6JpCL?j23kqGy
	hiv  = "5a305a6f4f4574424d585a775a315134553155684e673d3d" // gFh8KA1vpgT8SU!6
)

// AESEncode AES Encoding Mechanism
func AESEncode(encryptData, key, iv []byte) string {
	blockEnc, err := aes.NewCipher(key)
	utils.HandleError("Failed to initiate cipher", err)
	padding := aes.BlockSize - len(encryptData)%aes.BlockSize
	if padding == aes.BlockSize {
		padding = 0
	}
	paddedData := make([]byte, len(encryptData)+padding)
	copy(paddedData, encryptData)

	mode := cipher.NewCBCEncrypter(blockEnc, iv)
	ciphertext := make([]byte, aes.BlockSize+len(paddedData))
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedData)
	return hex.EncodeToString(ciphertext)
}

func main() {
	cmdArgs := utils.ParseFlags()
	if cmdArgs.Encode {
		baseData := utils.BaseEncode([]byte(cmdArgs.URL))
		hexData := utils.HexEncode(baseData)
		encodedData := AESEncode([]byte(hexData), []byte(cmdArgs.KEY), []byte(cmdArgs.IV))
		fmt.Println(encodedData)
	} else {
		utils.Usage()
	}
}
