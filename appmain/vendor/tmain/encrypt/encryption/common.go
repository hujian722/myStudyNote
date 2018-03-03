package encryption

import (
	"bytes"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const (
	Src  = "src"
	Des  = "des"
	Bit  = "bit"
	Key  = "key"
	Sign = "sign"
)

var SrcName = cli.StringFlag{
	Name:  Src,
	Usage: "源文本",
}
var DesName = cli.StringFlag{
	Name:  Des,
	Usage: "目标文本",
}
var RsaBit = cli.IntFlag{
	Name:  Bit,
	Usage: "Rsa生成位数选择",
}
var KeyName = cli.StringFlag{
	Name:  Key,
	Usage: "使用的密钥",
}
var SignName = cli.StringFlag{
	Name:  Sign,
	Usage: "数字签名",
}

func checkFileExist(f string) bool {
	_, err := os.Stat(f)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

const (
	ForeBlack  = iota + 30 //30
	ForeRed                //31
	ForeGreen              //32
	ForeYellow             //33
	ForeBlue               //34
	ForePurple             //35
	ForeCyan               //36
	ForeWhite              //37
)

func PrintColor(s string, color int) {
	var b = new(bytes.Buffer)
	sc := fmt.Sprintf("\033[1;%dm", color)
	b.Write([]byte(sc))
	b.Write([]byte(s))
	b.Write([]byte("\033[0m"))
	fmt.Println(b.String())
}
func PrintRed(s string) {
	PrintColor(s, ForeRed)
}
func PrintGreen(s string) {
	PrintColor(s, ForeGreen)
}
func PrintYellow(s string) {
	PrintColor(s, ForeYellow)
}
