package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/urfave/cli"
)

var iv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func aesEncrypt(key []byte, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	finalMsg := make([]byte, len(text))

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(finalMsg, text)
	return finalMsg, nil
}

func aesDecrypt(key []byte, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	return text, nil
}
func passwordCheck() ([]byte, error) {
	//	s := c.String(Src)
	fmt.Printf("Password: ")
	passFirst, err := gopass.GetPasswdMasked()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Repeat Your Password: ")
	passSecond, err := gopass.GetPasswdMasked()
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(passFirst, passSecond) {
		return nil, errors.New("PassWord Does Not Match")
	}
	if len(passSecond) > 16 {
		return nil, errors.New("PassWord len must <=16")
	}
	p := append(passSecond, iv[len(passSecond):]...)
	return p, nil
}
func AesEncodeAction(c *cli.Context) error {
	s := c.String(Src)
	des := c.String(Des)
	//
	if !checkFileExist(s) {
		PrintRed(fmt.Sprintf("File %s does not exist", s))
		return nil
	}
	//
	p, err := passwordCheck()
	if err != nil {
		PrintRed(err.Error())
		return nil
	}

	//
	fileSuffix := path.Ext(s)
	fileDir := strings.TrimSuffix(s, fileSuffix)
	PrintYellow(" start  to encrypt: " + s)
	//
	bf, err := ioutil.ReadFile(s)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	df, err := aesEncrypt(p, bf)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	if des == "" {
		des = fileDir + ".enc"
	}
	//sprefix := strconv.FormatInt(time.Now().UnixNano(), 10)
	f, err := os.OpenFile(des, os.O_RDONLY|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	f.Write(df)
	f.Sync()
	f.Close()
	PrintYellow("success to encrypt: " + des)
	return nil
}
func AesDecodeAction(c *cli.Context) error {
	s := c.String(Src)
	des := c.String(Des)
	//
	if !checkFileExist(s) {
		PrintRed(fmt.Sprintf("File %s does not exist", s))
		return nil
	}
	//
	p, err := passwordCheck()
	if err != nil {
		PrintRed(err.Error())
		return nil
	}

	//
	fileSuffix := path.Ext(s)
	fileDir := strings.TrimSuffix(s, fileSuffix)
	PrintYellow(" start  to decrypt: " + s)
	//
	bf, err := ioutil.ReadFile(s)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	df, err := aesDecrypt(p, bf)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	if des == "" {
		des = fileDir + ".dec"
	}
	//sprefix := strconv.FormatInt(time.Now().UnixNano(), 10)
	f, err := os.OpenFile(des, os.O_RDONLY|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	f.Write(df)
	f.Sync()
	f.Close()
	PrintYellow("success to decrypt: " + des)
	return nil
}
