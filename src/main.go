package main

import (
	"awesomeProject/src/accumulator/rsaAccumulator"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main()  {
	//rsa生成密钥测试
	//privateKey, err := rsa.GenerateKey(rand.Reader, 12)
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(privateKey)

	//贝祖系数计算测试
	//a := big.NewInt(1105)
	//b := big.NewInt(7)
	//x,y := util.Bezoute_Coefficients(*a,*b)
	//fmt.Printf("x[%s],y[%s]\n",x.String(),y.String())

	//以文本形式读取文件取
	file, err := os.Open("test_server.crt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	crtStr := ""
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		crtStr += str
		//fmt.Print(str)
	}
	//fmt.Println(crtStr)
	//fmt.Println("文件读取结束...")
	//var test = new(big.Int)
	//test.Mod(a,b)
	//fmt.Println(test)

	//贝祖
	//x,y := util.Bezoute_coefficients(1105,7)
	//fmt.Printf("x[%d],y[%d]\n",x,y)

	//累加器添加成员测试
	test := rsaAccumulator.New()
	//fmt.Println(test)
	fmt.Println("P:"+test.GetP().String())
	fmt.Println("Q:"+test.GetQ().String())
	fmt.Println("N:"+test.GetN().String())
	fmt.Println("初始累加器状态A0:",test.GetA0().String())
	fmt.Println("当前累加器状态A:",test.GetA().String())
	test.AddMember("5")
	fmt.Println("当前累加器状态A:",test.GetA().String())
	test.AddMember("7")
	fmt.Println("当前累加器状态A:",test.GetA().String())
	test.AddMember("11")
	fmt.Println("当前累加器状态A:",test.GetA().String())
	test.AddMember("13")
	fmt.Println("当前累加器状态A:",test.GetA().String())
	test.AddMember(crtStr)
	witness:=test.ProveMembership(crtStr)
	fmt.Println("witnss",witness)
	fmt.Println("A:",test.GetA())
	start := time.Now()
	if test.VerifyMembership(crtStr,witness){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	fmt.Println(time.Since(start))
	//start := time.Now()
	//fmt.Println(time.Since(start))
}
