package main

import (
	"awesomeProject/src/utils/util"
	"fmt"
	"math/big"
)

func main()  {
	//rsa生成密钥测试
	//privateKey, err := rsa.GenerateKey(rand.Reader, 12)
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(privateKey)

	//贝祖系数计算测试
	a := big.NewInt(1105)
	b := big.NewInt(7)
	x,y := util.Bezoute_Coefficients(*a,*b)
	fmt.Printf("x[%s],y[%s]\n",x.String(),y.String())

	//var test = new(big.Int)
	//test.Mod(a,b)
	//fmt.Println(test)

	//贝祖
	//x,y := util.Bezoute_coefficients(1105,7)
	//fmt.Printf("x[%d],y[%d]\n",x,y)

	//累加器添加成员测试
	//test := rsaAccumulator.New()
	//fmt.Println(test)
	//fmt.Println(test.GetP())
	//fmt.Println(test.GetQ())
	//fmt.Println(test.GetN())
	//fmt.Println("A0",test.GetA0())
	//fmt.Println("A:",test.GetA())
	//test.AddMember(big.NewInt(5))
	//test.AddMember(big.NewInt(7))
	//witness:=test.ProveMembership(big.NewInt(5))
	//fmt.Println("witnss",witness)
	//fmt.Println("A:",test.GetA())
	//test.VerifyMembership(big.NewInt(5),witness)


}
