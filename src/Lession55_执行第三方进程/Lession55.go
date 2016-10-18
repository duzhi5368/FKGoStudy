package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {
	//------------------------------------
	// 启动第三方进程，使用exec.Command
	//------------------------------------
	// 不需要参数的第三方进程
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
	// 给予输入和输出的第三方进程
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	// 显式的接管Pipe，接下来指定其输入输出
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 注意这里的参数给予
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

	fmt.Println("==================")
	//------------------------------------
	// 启动第三方进程，并关闭自身进程 使用syscall.Exec
	//------------------------------------
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	// 一些参数，注意，第一个参数应该是进程名字
	args := []string{"ls", "-a", "-l", "-h"}
	// 获取当前 的环境变量
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
