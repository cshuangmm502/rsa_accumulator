package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
)

const PRIME_CERTAINTY  = 5

var smallPrimes = []int64{
	3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53,57,59,61,
}


func GenerateLargePrime(bitLength int) *big.Int {
	prime,err := rand.Prime(rand.Reader,bitLength)
	if err != nil {
		fmt.Println("素数生成出错:",err)
	}
	return prime
}

func CheckPrime(p *big.Int)bool{
	return p.ProbablyPrime(PRIME_CERTAINTY)
}

//func HashToPrime(input *big.Int) (*big.Int,*big.Int){
//	count := 0
//	x := HashToLength(input)
//	//fmt.Println(x)
//	for{
//		if CheckPrime(x)==true {
//			break
//		}
//		x.Add(x,big.NewInt(1))
//		count++
//	}
//	return x,big.NewInt(int64(count));
//}

func HashToPrime(input string) (*big.Int,int){
	count := 0
	x := HashToLength(input)
	//fmt.Println(x)
	for{
		if CheckPrime(x)==true {
			break
		}
		x.Add(x,big.NewInt(1))
		count++
	}
	return x,count;
}

func HashToPrimeWithNonce(input string,nonce int)(*big.Int){
	val := big.NewInt(int64(nonce))
	return val.Add(val,HashToLength(input))
}
//采用sha256做hash,截取前bitLength(256)位
//func HashToLength(x *big.Int) *big.Int {
//	var randomHexString string
//	randomHexString = x.String()
//	hash := sha256.New()
//	hash.Write([]byte(randomHexString))
//	hashResult := hash.Sum(nil)
//	//32个字节表示
//	//fmt.Println(hashResult)
//	hashcode := hex.EncodeToString(hashResult)
//	//64个十六进制数表示
//	//fmt.Println(hashcode)
//	n := new(big.Int)
//	n,ok := n.SetString(hashcode,16)
//	if !ok {
//		fmt.Println("SetString: error")
//		return big.NewInt(0)
//	}
//	//十进制表示
//	//fmt.Println(n)
//	return n
//}

func HashToLength(x string) *big.Int {
	var randomHexString string
	randomHexString = x
	hash := sha256.New()
	hash.Write([]byte(randomHexString))
	hashResult := hash.Sum(nil)
	//32个字节表示
	//fmt.Println(hashResult)
	hashcode := hex.EncodeToString(hashResult)
	//64个十六进制数表示
	//fmt.Println(hashcode)
	n := new(big.Int)
	n,ok := n.SetString(hashcode,16)
	if !ok {
		fmt.Println("SetString: error")
		return big.NewInt(0)
	}
	//十进制表示
	//fmt.Println(n)
	return n
}

func GenerateRandomNumber(min big.Int,max big.Int) *big.Int{
	temp := big.NewInt(0)
	temp.Sub(&max,&min)
	temp.Add(temp,big.NewInt(1))
	result,_ := rand.Int(rand.Reader,temp)
	return result.Add(result,&min)
}


func exgcd(a int,b int,x *int,y *int)(int){
	if b>a{
		return exgcd(b,a,y,x)
	}
	if b==0{
		*x = 1
		*y = 0
		return a
	}
	var x1 = new(int)
	var d = exgcd(b,a%b,x1,x)
	*y = *x1 - a/b**x
	return d
}

func Bezoute_coefficients(a int,b int)(int,int){
	var x = new(int)
	var y = new(int)
	_ = exgcd(a,b,x,y)
	return *x,*y
}

//func Exgcd(a big.Int,b big.Int,x *big.Int,y *big.Int) big.Int {
//	if b.Cmp(&a)==1{
//		return Exgcd(b,a,y,x)
//	}
//	if b.Cmp(big.NewInt(0))==0{
//		x.Set(big.NewInt(1))
//		y.Set(big.NewInt(0))
//		return a
//	}
//	var x1 = new(big.Int)
//	var temp,temp1,temp2 big.Int
//	temp.Mod(&a,&b)
//	var d = Exgcd(b,temp,x1,x)
//	temp1.Mod(&a,&b)
//	temp2.Mul(&temp1,x)
//	y.Sub(x1,&temp2)
//	return d
//}

func Exgcd(a big.Int,b big.Int) (big.Int,big.Int,big.Int){
	var x0,x1,y0,y1 big.Int
	x0.Set(big.NewInt(1))
	x1.Set(big.NewInt(0))
	y0.Set(big.NewInt(0))
	y1.Set(big.NewInt(1))
	for a.Cmp(big.NewInt(0))==1{
		var q big.Int
		var temp big.Int
		q.Div(&b,&a)
		temp.Set(&a)
		a.Mod(&b,&a)
		b.Set(&temp)
		temp.Set(&x1)
		x1.Mul(&q,&x1)
		x1.Sub(&x0,&x1)
		x0.Set(&temp)
		temp.Set(&y1)
		y1.Mul(&q,&y1)
		y1.Sub(&y0,&y1)
		y0.Set(&temp)
	}
	return b,x0,y0
}

//func Bezoute_Coefficients(a big.Int,b big.Int)(big.Int,big.Int){
//	var x = new(big.Int)
//	var y = new(big.Int)
//	_ = Exgcd(a,b,x,y)
//	return *x,*y
//}

func Mul_inv(b big.Int,n big.Int)(big.Int){
	g,x,_ := Exgcd(b,n)
	if g.Cmp(big.NewInt(1))==0{
		return *g.Mod(&x,&n)
	}
	return big.Int{}
}

func Bezoute_Coefficients(a big.Int,b big.Int)(big.Int,big.Int){
	_,x0,y0 := Exgcd(a,b)
	return x0,y0
}

func calculate_product(list []*big.Int)*big.Int{
	base := big.NewInt(1)
	for _,i := range list{
		base.Mul(base,i)
	}
	return base
}

func Create_all_membership_witness(A0 *big.Int,set map[string]int,N *big.Int)[]*big.Int{
	var primes []*big.Int
	sorted_keys := make([]string, 0)
	for k, _ := range set {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	for _,k := range sorted_keys{
		prime := HashToPrimeWithNonce(k,set[k])
		primes=append(primes, prime)
		fmt.Println(k,prime)
	}
	//for k,v := range set{
	//	prime := HashToPrimeWithNonce(k,v)
	//	primes=append(primes, prime)
	//	fmt.Println(k,prime)
	//}

	fmt.Println(primes)
	return Root_factor(A0,primes,N)
}

func Root_factor(g *big.Int,primes []*big.Int,N *big.Int)[]*big.Int{
	n := len(primes)
	if n==1{
		var result = make([]*big.Int,1)
		result[0]=g
		//var result1 []*big.Int
		//result1 = append(result1,g)
		return result
	}

	n_tag := n/2

	primes_L := primes[n_tag:n]
	product_L := calculate_product(primes_L)
	g_L := big.NewInt(1)
	g_L.Exp(g,product_L,N)

	primes_R := primes[0:n_tag]
	product_R := calculate_product(primes_R)
	g_R := big.NewInt(1)
	g_R.Exp(g,product_R,N)

	L := Root_factor(g_L, primes_R,N)
	R := Root_factor(g_R, primes_L,N)

	//var result []*big.Int
	////result = append(result, L...)
	//result = append(L, R...)
	return append(L,R...)
}


//func testEct(a int,b int)(int,int,int){
//
//}
//
//func Bezoute(a int,b int)(int,int){
//	pam1,pam2,pam3 := testEct(a,b)
//	var x = new(int)
//	var y = new(int)
//	_ = exgcd(a,b,x,y)
//	return *x,*y
//}