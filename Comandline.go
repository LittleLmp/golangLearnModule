package test 

func Comandline(a int) bool{
	// "flag"
	// var mode = flag.String("mode","","process mode")
	// flag.Parse() // 解析命令行参数
	// fmt.Println(*mode) // 输出命令行参数
	if(a>0){
		return true
	}else{
		return false
	}
}