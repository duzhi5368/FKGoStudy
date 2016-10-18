package FKLib

//--------------------------------------------------------------
// string 去除头部
// 参数：s 准备被去除的字符串
// 		w 头部字符串
// 返回值：string 去除头部后的字符串
func TrimString(s, w string) string {
	ss := len(s)
	ws := len(w)

	if ss >= ws && s[:ws] == w {
		s = s[ws:]
		ss -= ws
	}

	if ss >= ws && s[ss-ws:] == w {
		s = s[:ss-ws]
	}

	return s
}

//--------------------------------------------------------------
// string切分
// 参数：s 准备被切分的字符串
// 		d 切割符
// 返回值：[]string 字符串切片
func SplitString(s, d string) []string {
	var i, c, n int

	l := len(s)

	for i < l {
		if s[i:i+1] == d {
			c++
		}
		i++
	}

	a := make([]string, c+1)
	c, i = 0, 0

	for i < l {
		if i == l-1 {
			a[c] = s[n : i+1]
			break
		}
		if s[i:i+1] == d {
			a[c] = s[n:i]
			n = i + 1
			c++
		}
		i++
	}

	return a
}

//--------------------------------------------------------------
