package controllers

import (
	"io/ioutil"
	"os"
	"crypto/des"
	"crypto/cipher"
	"bytes"
	"github.com/levigross/grequests"
	"fmt"
	"encoding/json"
	"encoding/base64"
)

const url = "http://disk.bjsasc.com:8180/NetDisk/rest/mobile"

func ListDir(dirPth string) (files []string, err error) {
	files = make([]string, 0, 5)

	dir, err := ioutil.ReadDir(dirPth)

	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}

		files = append(files, dirPth+PthSep+fi.Name())

	}
	return files, nil
}

func TripleDesEncrypt(origData []byte) ([]byte, error) {
	key := []byte("sr$*)(ruan$@lx100$#365#$")
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte("01234567"))
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func GetToken(username, pass string) (string,  error)  {

	//url := "http://127.0.0.1:8080/list"
	paras := &grequests.RequestOptions{Params: map[string]string{"userName": username, "passWord": EncryptPass(pass), "method": "login"}}
	res, err := grequests.Get(url, paras)
	if err != nil {
		fmt.Println(err)
		var a=""
		return a,err
	}
	var token Token

	json.Unmarshal(res.Bytes(), &token)

	return token.Token,nil
}

func EncryptPass(orig string) string {
	s, _ := TripleDesEncrypt([]byte(orig))
	encStr := base64.StdEncoding.EncodeToString(s)
	return encStr
}
