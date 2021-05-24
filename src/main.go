package main

import (
	"awesomeProject/src/accumulator/rsaAccumulator"
	"awesomeProject/src/utils/util"
	"fmt"
)

func main()  {
	test()
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
	//file, err := os.Open("test_server.crt")
	//if err != nil {
	//	fmt.Println("文件打开失败 = ", err)
	//}
	////及时关闭 file 句柄，否则会有内存泄漏
	//defer file.Close()
	//crtStr := ""
	////创建一个 *Reader ， 是带缓冲的
	//reader := bufio.NewReader(file)
	//for {
	//	str, err := reader.ReadString('\n') //读到一个换行就结束
	//	if err == io.EOF {                  //io.EOF 表示文件的末尾
	//		break
	//	}
	//	crtStr += str
	//	//fmt.Print(str)
	//}
	//fmt.Println(crtStr)
	//fmt.Println("文件读取结束...")
	//var test = new(big.Int)
	//test.Mod(a,b)
	//fmt.Println(test)

	//贝祖
	//x,y := util.Bezoute_coefficients(1105,7)
	//fmt.Printf("x[%d],y[%d]\n",x,y)


	//start := time.Now()
	//fmt.Println(time.Since(start))

}

func test(){
	//var list []*big.Int
	//a := big.NewInt(3)
	//list = append(list,a)
	//fmt.Println(list)
	test := rsaAccumulator.New()

	N := test.GetN()
	dict := make(map[string]int)
	test.AddMember("37")
	test.AddMember("59")
	test.AddMember("73")
	test.AddMember("117")
	dict["37"]=test.GetVal("37")
	dict["59"]=test.GetVal("59")
	dict["73"]=test.GetVal("73")
	dict["117"]=test.GetVal("117")

	fmt.Println(dict)
	//var primes []*big.Int
	//primes = append(primes, big.NewInt(3))
	//primes = append(primes, big.NewInt(5))
	//primes = append(primes, big.NewInt(7))
	//primes = append(primes, big.NewInt(11))
	//result :=util.Root_factor(big.NewInt(2),primes,N)
	//fmt.Println(result)

	witnesses := util.Create_all_membership_witness(test.GetA0(),dict,N)
	fmt.Printf("累加器状态:%s\n",test.GetA().String())
	fmt.Println(witnesses)
	fmt.Println("37的素数hash值：")
	fmt.Println(util.HashToPrime("37"))
	fmt.Println(test.VerifyMembership("59",witnesses[1]))

}

//func root_factor(g big.Int,primes []big.Int,N big.Int)[]big.Int{
//	n := len(primes)
//	if n==1{
//		return primes
//	}
//	n_tag := n/2
//	primes_L := primes[n_tag:n]
//	product_L := calculate_product(primes_L)
//	g_L := g.Exp(&g,&product_L,&N)
//}
//
//func calculate_product(list []big.Int)big.Int{
//	base := big.NewInt(1)
//	for _,i := range list{
//		base.Mul(base,i)
//	}
//	return base
//}