// 声明本包名
package main

// 加载第三方包，fmt也是一个系统包
// 注意，这里是使用的 圆括号() 不是花括号
// import导入包，也可以写多行 import "a"  import "b"，但是更建议用圆括号的方式
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 常量
const sConst string = "正在等待输入"

// 注意  "{" 不允许在下一行
//  这样是不可以的...
// 	if (A)
//	{
// 	}
//  必须这样
// 	if (A){
//	}
func main() {
	fmt.Println("Hello World")

	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println(sConst)

	// 输入等待
	var input string
	fmt.Scanln(&input)
	fmt.Println("================================")

	// 重置输入缓冲区
	scanner := bufio.NewScanner(os.Stdin)

	// 处理三次
	nTimes := 0
	for scanner.Scan() {
		// 将输入数据进行大写处理
		ucl := strings.ToUpper(scanner.Text())
		// 输出
		fmt.Println(ucl)

		nTimes++
		if nTimes >= 3 {
			break
		}
	}

	// 如果有错误，则输出错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
