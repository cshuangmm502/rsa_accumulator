package main

import (
	"awesomeProject/src/accumulator/rsaAccumulator"
	"awesomeProject/src/utils/util"
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"sort"
	"time"
)

func main()  {
	start := time.Now()
	//test()
	//rsa生成密钥测试
	//privateKey, err := rsa.GenerateKey(rand.Reader, 12)
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(privateKey)

	//贝祖系数计算测试
	a := big.NewInt(165)
	b := big.NewInt(7)
	x,y := util.Bezoute_Coefficients(*a,*b)
	fmt.Printf("x[%s],y[%s]\n",x.String(),y.String())

	//crtStr := readCert("testCert.pem")
	//crtStr2 := readCert("test_server.crt")
	//test := rsaAccumulator.New()
	//fmt.Println("A:"+test.GetA().String())
	//fmt.Println("A0:"+test.GetA0().String())
	//test.AddMember(crtStr)
	//crt1_proof := test.ProveMembership(crtStr)
	//fmt.Println("Crt1Proof:"+crt1_proof.String())
	//fmt.Println("A:"+test.GetA().String())
	//test.AddMember(crtStr2)
	//fmt.Println("A:"+test.GetA().String())
	//var test = new(big.Int)
	//test.Mod(a,b)

	//贝祖
	x1,y1 := util.Bezoute_coefficients(165,7)
	fmt.Printf("x[%d],y[%d]\n",x1,y1)


	//start := time.Now()
	//fmt.Println(time.Since(start))
	fmt.Println(time.Since(start))
}

func readCert(fileName string)string{
	//以文本形式读取文件取
	file, err := os.Open(fileName)
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
	fmt.Println(crtStr)
	fmt.Println("文件读取结束...")
	return crtStr
}

func test(){
	//var list []*big.Int
	//a := big.NewInt(3)
	//list = append(list,a)
	//fmt.Println(list)
	test := rsaAccumulator.New()
	fmt.Println("A:"+test.GetA().String())
	fmt.Println("A0:"+test.GetA0().String())
	fmt.Println("N:"+test.GetN().String())
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
	fmt.Println(witnesses)
	fmt.Println(witnesses)
	fmt.Println("37的素数hash值：")
	fmt.Println(util.HashToPrime("37"))
	sorted_keys := make([]string, 0)
	for k, _ := range dict {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var index = 0
	for k,v := range sorted_keys{
		if v=="37"{
			index=k
		}
	}
	fmt.Println(test.VerifyMembership("37",witnesses[index]))

}
