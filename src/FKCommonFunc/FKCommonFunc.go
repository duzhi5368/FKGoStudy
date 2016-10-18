package FKCommonFunc

func GetString() string {
	return "HelloWorld"
}

// 注意函数参数,类型和返回值的顺序
func Sqrt(p_fX float64) float64 {
	fRet := 0.0
	for i := 0; i < 1000; i++ {
		fRet -= (fRet*fRet - p_fX) / (2 * p_fX)
	}
	return fRet
}

// 注意，该函数是无法被导出的，原因是首字母是小写的
// 只有首字母为大写的函数才会被导出
func sqrt_inside(p_fX float64) float64 {
	fRet := 0.0
	for i := 0; i < 1000; i++ {
		fRet -= (fRet*fRet - p_fX) / (2 * p_fX)
	}
	return fRet
}
