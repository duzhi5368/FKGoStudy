package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func main() {
	strCur := getCurrentDirectory()
	strDst := getParentDirectory(getParentDirectory(strCur))
	strDstExe := strDst + "/github.com/akavel/rsrc/rsrc.exe"

	arg := "rsrc -manifest " + strDst + "/LibTest_WalkUILib/main.manifest -o " + strDst + "/LibTest_WalkUILib/main.syso"
	fmt.Println("Arg = " + arg)

	cmd := exec.Command(strDstExe, arg)

	//会向 cmd.Stdout和cmd.Stderr写入信息,其实cmd.Stdout==cmd.Stderr,具体可见源码
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("调用第三方进程结果:\n%v\n\n%v\n\n%v", string(output), cmd.Stdout, cmd.Stderr)
}
