package encryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/urfave/cli"
)

//
func RsaGenKeyPairAction(c *cli.Context) error {
	var (
		bit = c.Int(Bit)
	)
	if bit <= 1024 {
		bit = 1024
	} else if bit <= 2048 {
		bit = 2048
	} else {
		bit = 1024
	}
	//
	reader := rand.Reader
	key, err := rsa.GenerateKey(reader, bit)
	checkError(err)

	publicKey := key.PublicKey
	PrintYellow(" start  to generate priv.pem pub.pem")
	savePrivitePEMKey("priv.pem", key)
	savePublicPEMKey("pub.pem", publicKey)
	PrintYellow("success to generate")
	return nil
}

//
func RsaEncodeAction(c *cli.Context) error {
	key := c.String(Key)
	src := c.String(Src)
	des := c.String(Des)
	//
	if !checkFileExist(key) {
		PrintRed(fmt.Sprintf("Key File %s does not exist", key))
		return nil
	}
	if !checkFileExist(src) {
		PrintRed(fmt.Sprintf("Src File %s does not exist", src))
		return nil
	}
	//
	keybyte, err := ioutil.ReadFile(key)
	checkError(err)
	srcbyte, err := ioutil.ReadFile(src)
	checkError(err)

	fileSuffix := path.Ext(src)
	fileDir := strings.TrimSuffix(src, fileSuffix)
	PrintYellow(" start  to encrypt: " + src)
	//
	desb, err := RsaEncrypt(srcbyte, keybyte)
	checkError(err)
	//
	if des == "" {
		des = fileDir + ".enc"
	}
	//sprefix := strconv.FormatInt(time.Now().UnixNano(), 10)
	f, err := os.OpenFile(des, os.O_RDONLY|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	f.Write(desb)
	f.Sync()
	f.Close()
	PrintYellow("success to encrypt: " + des)
	return nil
}

//
func RsaDecodeAction(c *cli.Context) error {
	key := c.String(Key)
	src := c.String(Src)
	des := c.String(Des)
	//
	if !checkFileExist(key) {
		err := errors.New(fmt.Sprintf("Key File %s does not exist", key))
		checkError(err)
		return nil
	}
	if !checkFileExist(src) {
		err := errors.New(fmt.Sprintf("Src File %s does not exist", src))
		checkError(err)
		return nil
	}
	//
	keybyte, err := ioutil.ReadFile(key)
	checkError(err)
	srcbyte, err := ioutil.ReadFile(src)
	checkError(err)

	fileSuffix := path.Ext(src)
	fileDir := strings.TrimSuffix(src, fileSuffix)
	PrintYellow(" start  to decrypt: " + src)
	//
	desb, err := RsaDecrypt(srcbyte, keybyte)
	checkError(err)
	//
	if des == "" {
		des = fileDir + ".dec"
	}
	//sprefix := strconv.FormatInt(time.Now().UnixNano(), 10)
	f, err := os.OpenFile(des, os.O_RDONLY|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	f.Write(desb)
	f.Sync()
	f.Close()
	PrintYellow("success to decrypt: " + des)
	return nil
}

//
func RsaSignAction(c *cli.Context) error {
	key := c.String(Key)
	src := c.String(Src)
	des := c.String(Des)
	//
	if !checkFileExist(key) {
		err := errors.New(fmt.Sprintf("Key File %s does not exist", key))
		checkError(err)
		return nil
	}
	if !checkFileExist(src) {
		err := errors.New(fmt.Sprintf("Src File %s does not exist", src))
		checkError(err)
		return nil
	}
	//
	keybyte, err := ioutil.ReadFile(key)
	checkError(err)
	srcbyte, err := ioutil.ReadFile(src)
	checkError(err)

	fileSuffix := path.Ext(src)
	fileDir := strings.TrimSuffix(src, fileSuffix)
	PrintYellow(" start  to signature: " + src)
	// 逻辑
	keyBlock, _ := pem.Decode(keybyte)
	if keyBlock == nil {
		checkError(errors.New("pem.Decode error"))
	}
	priv, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	checkError(err)
	hashed := sha256.Sum256(srcbyte)
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:])
	checkError(err)
	//
	if des == "" {
		des = fileDir + ".sign"
	}
	//sprefix := strconv.FormatInt(time.Now().UnixNano(), 10)
	f, err := os.OpenFile(des, os.O_RDONLY|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		PrintRed(err.Error())
		return nil
	}
	f.Write(signature)
	f.Sync()
	f.Close()
	PrintYellow("success to signature: " + des)
	return nil
}

//
func RsaUnSignAction(c *cli.Context) error {
	key := c.String(Key)
	src := c.String(Src)
	sign := c.String(Sign)
	//
	if !checkFileExist(key) {
		err := errors.New(fmt.Sprintf("Key File %s does not exist", key))
		checkError(err)
		return nil
	}
	if !checkFileExist(src) {
		err := errors.New(fmt.Sprintf("Src File %s does not exist", src))
		checkError(err)
		return nil
	}
	if !checkFileExist(sign) {
		err := errors.New(fmt.Sprintf("Signature File %s does not exist", sign))
		checkError(err)
		return nil
	}
	//
	keybyte, err := ioutil.ReadFile(key)
	checkError(err)
	srcbyte, err := ioutil.ReadFile(src)
	checkError(err)
	signbyte, err := ioutil.ReadFile(sign)
	checkError(err)
	PrintYellow(" start  to verify signature: " + src)
	// 逻辑
	keyBlock, _ := pem.Decode(keybyte)
	if keyBlock == nil {
		checkError(errors.New("pem.Decode error"))
	}
	pub, err := x509.ParsePKIXPublicKey(keyBlock.Bytes)
	checkError(err)
	if _, ok := pub.(*rsa.PublicKey); !ok {
		checkError(errors.New("x509.ParsePKIXPublicKey error"))
	}
	hashed := sha256.Sum256(srcbyte)
	err = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hashed[:], signbyte)
	if err != nil {
		PrintRed("  ===========================")
		PrintRed("   Error from verification")
		PrintRed("  ===========================")
		checkError(err)
		return nil
	}
	//
	PrintGreen("  ===========================")
	PrintGreen("   signature is verified")
	PrintGreen("  ===========================")
	return nil
}

//=========================================================
func savePrivitePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	x509Bytes, err := x509.MarshalPKIXPublicKey(&pubkey)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509Bytes,
	}

	pemfile, err := os.Create(fileName)
	checkError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	checkError(err)
}

//==============================================
// 加密
func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//==============================================
func checkError(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			file = "???"
			line = 0
		}
		s := fmt.Sprintf("Err: %s.%d :%s", file, line, err.Error())
		PrintRed(s)
		os.Exit(1)
	}
}
