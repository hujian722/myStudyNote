package encryption

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func Base64EncodeAction(c *cli.Context) error {
	s := c.String(Src)
	des := c.String(Des)
	if !checkFileExist(s) {
		PrintRed(fmt.Sprintf("File %s does not exist", s))
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
	df := base64Encode(bf)
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
func Base64DecodeAction(c *cli.Context) error {
	s := c.String(Src)
	des := c.String(Des)
	if !checkFileExist(s) {
		PrintRed(fmt.Sprintf("File %s does not exist", s))
		return nil
	}
	//
	fileSuffix := path.Ext(s)
	fileDir := strings.TrimSuffix(s, fileSuffix)
	PrintYellow(" start  to decrypt: " + s)
	bf, err := ioutil.ReadFile(s)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	df, err := base64Decode(bf)
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
