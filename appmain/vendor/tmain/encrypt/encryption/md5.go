package encryption

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io/ioutil"

	"github.com/urfave/cli"
)

func Md5EncodeAction(c *cli.Context) error {
	s := c.String(Src)

	if !checkFileExist(s) {
		PrintRed(fmt.Sprintf("File %s does not exist", s))
		return nil
	}
	//
	bf, err := ioutil.ReadFile(s)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	h := md5.New()
	h.Write(bf)
	cipherText := h.Sum(nil)
	ss := hex.EncodeToString(cipherText)
	PrintGreen("  ===========================")
	PrintGreen("  " + ss)
	PrintGreen("  ===========================")
	return nil
}
func Crc32EncodeAction(c *cli.Context) error {
	s := c.String(Src)

	if !checkFileExist(s) {
		PrintRed(fmt.Sprintf("File %s does not exist", s))
		return nil
	}
	//
	bf, err := ioutil.ReadFile(s)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	h := crc32.NewIEEE()
	h.Write(bf)
	cipherText := h.Sum(nil)
	ss := hex.EncodeToString(cipherText)
	PrintGreen("  ===========================")
	PrintGreen("  " + ss)
	PrintGreen("  ===========================")
	return nil
}
func Sha1EncodeAction(c *cli.Context) error {
	s := c.String(Src)

	if !checkFileExist(s) {
		PrintRed(fmt.Sprintf("File %s does not exist", s))
		return nil
	}
	//
	bf, err := ioutil.ReadFile(s)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	h := sha1.New()
	h.Write(bf)
	cipherText := h.Sum(nil)
	ss := hex.EncodeToString(cipherText)
	PrintGreen("  ===========================")
	PrintGreen("  " + ss)
	PrintGreen("  ===========================")
	return nil
}
func Sha256EncodeAction(c *cli.Context) error {
	s := c.String(Src)

	if !checkFileExist(s) {
		PrintRed(fmt.Sprintf("File %s does not exist", s))
		return nil
	}
	//
	bf, err := ioutil.ReadFile(s)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	h := sha256.New()
	h.Write(bf)
	cipherText := h.Sum(nil)
	ss := hex.EncodeToString(cipherText)
	PrintGreen("  ===========================")
	PrintGreen("  " + ss)
	PrintGreen("  ===========================")
	return nil
}
func Sha512EncodeAction(c *cli.Context) error {
	s := c.String(Src)

	if !checkFileExist(s) {
		PrintRed(fmt.Sprintf("File %s does not exist", s))
		return nil
	}
	//
	bf, err := ioutil.ReadFile(s)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	h := sha512.New()
	h.Write(bf)
	cipherText := h.Sum(nil)
	ss := hex.EncodeToString(cipherText)
	PrintGreen("  ===========================")
	PrintGreen("  " + ss)
	PrintGreen("  ===========================")
	return nil
}
