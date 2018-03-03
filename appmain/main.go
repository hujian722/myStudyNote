package main

import (
	"os"
	"tmain/encrypt/encryption"

	"github.com/urfave/cli"
)

//base64 编码
var (
	Base64Command = cli.Command{
		Name:  "base64",
		Usage: "Base64编码",
		//ArgsUsage:   "加密或者解密",
		Subcommands: []cli.Command{Base64EncodeCommand, Base64DecodeCommand},
		Description: "加密解密方法：Base64",
	}
	Base64EncodeCommand = cli.Command{
		Action: encryption.Base64EncodeAction,
		Name:   "encode",
		Usage:  "使用base64加密",
		//ArgsUsage: "",
		Flags: []cli.Flag{
			encryption.SrcName,
			encryption.DesName,
		},
		//Category: "BLOCKCHAIN COMMANDS",
		Description: "使用base64加密文本内容",
	}
	Base64DecodeCommand = cli.Command{
		Action: encryption.Base64DecodeAction,
		Name:   "decode",
		Usage:  "使用base64解密",
		//ArgsUsage: "加密或者解密",
		Flags: []cli.Flag{
			encryption.SrcName,
			encryption.DesName,
		},
		Description: `使用base64解密文本内容`,
	}
)

//aes加密
var (
	AesCommand = cli.Command{
		Name:        "aes",
		Usage:       "Aes加密解密",
		Subcommands: []cli.Command{AesEncodeCommand, AesDecodeCommand},
		Description: "加密解密方法：Aes",
	}
	AesEncodeCommand = cli.Command{
		Action: encryption.AesEncodeAction,
		Name:   "encode",
		Usage:  "使用Aes加密",
		Flags: []cli.Flag{
			encryption.SrcName,
			encryption.DesName,
		},
		//Category: "BLOCKCHAIN COMMANDS",
		Description: "使用Aes加密文本内容",
	}
	AesDecodeCommand = cli.Command{
		Action: encryption.AesDecodeAction,
		Name:   "decode",
		Usage:  "使用Aes解密",
		//ArgsUsage: "加密或者解密",
		Flags: []cli.Flag{
			encryption.SrcName,
			encryption.DesName,
		},
		Description: "使用Aes解密文本内容",
	}
)

//rsa加密
var (
	RsaCommand = cli.Command{
		Name:        "rsa",
		Usage:       "Rsa加密解密,密钥生成",
		Subcommands: []cli.Command{RsaGenKeyCommand, RsaEncodeCommand, RsaDecodeCommand, RsaSignCommand, RsaUnSignCommand},
		Description: "加密解密方法：Rsa",
	}
	RsaGenKeyCommand = cli.Command{
		Action: encryption.RsaGenKeyPairAction,
		Name:   "genkey",
		Usage:  "生成公钥私钥对",
		Flags: []cli.Flag{
			encryption.RsaBit,
		},
		Description: "生成公钥私钥对 priv.pem pub.pem",
	}
	RsaEncodeCommand = cli.Command{
		Action: encryption.RsaEncodeAction,
		Name:   "encode",
		Usage:  "使用Rsa加密",
		Flags: []cli.Flag{
			encryption.KeyName,
			encryption.SrcName,
			encryption.DesName,
		},
		//Category: "BLOCKCHAIN COMMANDS",
		Description: "使用Rsa加密文本内容",
	}
	RsaDecodeCommand = cli.Command{
		Action: encryption.RsaDecodeAction,
		Name:   "decode",
		Usage:  "使用Rsa解密",
		//ArgsUsage: "加密或者解密",
		Flags: []cli.Flag{
			encryption.KeyName,
			encryption.SrcName,
			encryption.DesName,
		},
		Description: "使用Rsa解密文本内容",
	}
	RsaSignCommand = cli.Command{
		Action: encryption.RsaSignAction,
		Name:   "sign",
		Usage:  "数字签名",
		Flags: []cli.Flag{
			encryption.KeyName,
			encryption.SrcName,
			encryption.DesName,
		},
		//Category: "BLOCKCHAIN COMMANDS",
		Description: "使用Rsa加密文本内容",
	}
	RsaUnSignCommand = cli.Command{
		Action: encryption.RsaUnSignAction,
		Name:   "unsign",
		Usage:  "解析数字签名",
		//ArgsUsage: "加密或者解密",
		Flags: []cli.Flag{
			encryption.KeyName,
			encryption.SrcName,
			encryption.SignName,
		},
		Description: "使用Rsa解密文本内容",
	}
)

//MD5 计算
var (
	Md5Command = cli.Command{
		Action:      encryption.Md5EncodeAction,
		Name:        "md5",
		Usage:       "计算md5",
		Description: "计算md5值",
		Flags: []cli.Flag{
			encryption.SrcName,
		},
	}
	Crc32Command = cli.Command{
		Action:      encryption.Crc32EncodeAction,
		Name:        "crc32",
		Usage:       "计算crc32",
		Description: "计算crc32值",
		Flags: []cli.Flag{
			encryption.SrcName,
		},
	}
	Sha1Command = cli.Command{
		Action:      encryption.Sha1EncodeAction,
		Name:        "sha1",
		Usage:       "计算sha1",
		Description: "计算sha1值",
		Flags: []cli.Flag{
			encryption.SrcName,
		},
	}
	Sha256Command = cli.Command{
		Action:      encryption.Sha256EncodeAction,
		Name:        "sha256",
		Usage:       "计算sha256",
		Description: "计算sha256值",
		Flags: []cli.Flag{
			encryption.SrcName,
		},
	}
	Sha512Command = cli.Command{
		Action:      encryption.Sha512EncodeAction,
		Name:        "sha512",
		Usage:       "计算sha512",
		Description: "计算sha512值",
		Flags: []cli.Flag{
			encryption.SrcName,
		},
	}
)

func main() {
	//实例化cli
	app := cli.NewApp()

	//Name可以设定应用的名字
	//app.Name = "encrypt"
	app.Usage = "一款小巧的加密解密工具--by hujian"
	app.Version = "1.0.0"
	app.Copyright = "Copyright 2018-2022 hujian"
	// Commands用于创建命令
	app.Commands = []cli.Command{
		Base64Command,
		AesCommand,
		RsaCommand,
		Md5Command,
		Crc32Command,
		Sha1Command,
		Sha256Command,
		Sha512Command,
	}

	// 接受os.Args启动程序
	app.Run(os.Args)
}
