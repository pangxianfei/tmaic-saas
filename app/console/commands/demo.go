package commands

import (
	"fmt"
	"gitee.com/pangxianfei/framework/cmd"
	"io"
	"io/ioutil"
	"os"
)

func init() {
	cmd.Add(&Demo{})
}

type Demo struct {
}

func (hw *Demo) Command() string {
	return "demo:say {words}"
}

func (hw *Demo) Description() string {
	return "This is a demonstration command"
}

func (hw *Demo) Handler(arg *cmd.Arg) error {
	words, err := arg.Get("words")
	if err != nil {
		return err
	}

	if words == nil {
		fmt.Println("This is a demonstration command")
		return nil
	}

	fmt.Println(*words)

	var stirs = string(*words) + "\n\n"

	fmt.Println("golang 写文件")
	var (
		fileName = "G:\\go\\src\\saas\\commonApp\\storage\\logs\\tmaic.log"
		content  = stirs
		file     *os.File
	)
	//文件是否存在
	if Exists(fileName) {
		//使用追加模式打开文件
		file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("打开文件错误：", err)
			return nil
		}
	} else {
		//不存在创建文件
		file, err = os.Create(fileName)
		if err != nil {
			fmt.Println("创建失败", err)
			return nil
		}
	}

	defer file.Close()
	//写入文件
	n, _ := io.WriteString(file, content)
	if err != nil {
		fmt.Println("写入错误：", err)
		return nil
	}
	fmt.Println("写入成功：n=", n)
	//读取文件
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("读取错误：", err)
		return nil
	}
	fmt.Println("读取成功，文件内容：", string(fileContent))

	return nil
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
