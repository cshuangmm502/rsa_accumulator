package rsaAccumulator

import (
	"awesomeProject/src/utils/pair"
	"awesomeProject/src/utils/util"
	"fmt"
	"math/big"
)

const (
	RSA_KEY_SIZE = 12
	RSA_PRIME_SIZE = RSA_KEY_SIZE/2
	ACCUMULATED_PRIME_SIZE = 128
)

//type RsaInter interface {
//	accumulator.Accumulator
//	GetN() big.Int
//}

type non_mem_witness struct {
	A *big.Int
	B *big.Int
}

type RSAAccumulator struct {
	data 	map[string]int			//["key":hashPrime]
	pair 	*pair.Pair
	p		*big.Int
	q 		*big.Int
	n		*big.Int
	//random 	big.Int
	a0		*big.Int
	a		*big.Int
}

func (rsaObj *RSAAccumulator)GetP() *big.Int {
	return rsaObj.p
}

func (rsaObj *RSAAccumulator)GetQ() *big.Int {
	return rsaObj.q
}

func (rsaObj *RSAAccumulator)GetN() *big.Int {
	return rsaObj.n
}

func (rsaObj *RSAAccumulator)GetA() *big.Int {
	return rsaObj.a
}

func (rsaObj *RSAAccumulator)GetA0() *big.Int {
	return rsaObj.a0
}

func (rsaObj *RSAAccumulator)GetVal(key string) int {
	return rsaObj.data[key]
}

//func (rsaObj *RSAAccumulator)AddMember(key *big.Int) *big.Int {
//	_,ok := rsaObj.data[key.String()]
//	if ok{
//		return rsaObj.a
//	}
//	hashPrime,_ :=util.HashToPrime(key)
//	//fmt.Println(hashPrime)
//	rsaObj.a.Exp(rsaObj.a,hashPrime,rsaObj.n)
//	rsaObj.data[key.String()]=hashPrime
//	return rsaObj.a
//}

func (rsaObj *RSAAccumulator)AddMember(key string) *big.Int {
	_,ok := rsaObj.data[key]
	if ok{
		return rsaObj.a
	}
	hashPrime,count :=util.HashToPrime(key)
	//fmt.Println(hashPrime)
	rsaObj.a.Exp(rsaObj.a,hashPrime,rsaObj.n)
	rsaObj.data[key]=count
	return rsaObj.a
}

//func (rsaObj *RSAAccumulator)UpdateExistProof(proof string,key string) *big.Int{
//	hashPrime,_ :=util.HashToPrime(key)
//}

func (rsaObj *RSAAccumulator)ProveMembership(key string) *big.Int {
	_,ok := rsaObj.data[key]
	if !ok{
		return nil
	}
	witness := rsaObj.iterateAndGetProductWithoutX(key)
	return witness.Exp(rsaObj.a0,witness,rsaObj.n)
}

func (rsaObj *RSAAccumulator)DeleteMember(bigInteger big.Int) *big.Int{
	return big.NewInt(0)
}

func (rsaObj *RSAAccumulator)VerifyMembership(key string,proof *big.Int) bool{
	hashPrime,_ := util.HashToPrime(key)
	return	doVerifyMembership(rsaObj.a,hashPrime,proof,rsaObj.n)
}

func (rsaObj *RSAAccumulator)ProveNonMembership(A big.Int,set []string,x string,g big.Int) *non_mem_witness{
	primes := big.NewInt(1)
	for _,element := range set{
		prime,_ := util.HashToPrime(element)
		primes.Mul(primes,prime)
	}
	x_prime,_ := util.HashToPrime(x)
	b,a := util.Bezoute_Coefficients(*primes,*x_prime)
	fmt.Println(&b)
	fmt.Println(&a)
	result_b := big.NewInt(1)
	result_b.Exp(&g,&b,rsaObj.n)
	non_mem_witness := &non_mem_witness{
		A: &a,
		B: result_b,
	}
	return non_mem_witness
}

//func (rsaObj *RSAAccumulator)ProveNonmembership(A0 big.Int,set []string,x string,n big.Int) *non_mem_witness{
//	for _,val := range set{
//		if x==val{
//			return nil
//		}
//	}
//	primes := big.NewInt(1)
//	for _,element := range set{
//		prime,_ := util.HashToPrime(element)
//		primes.Mul(primes,prime)
//	}
//	x_prime,_ := util.HashToPrime(x)
//	a,b := util.Bezoute_Coefficients(*primes,*x_prime)
//	d := big.NewInt(1)
//	inverse_A0 := big.NewInt(1)
//	if a.Cmp(big.NewInt(0))<0{
//		a.Abs(&a)
//		inverse_A0 := util.Mul_inv(A0,n)
//		b.Exp(&inverse_A0,&a,&n)
//	}else{
//
//	}
//}

//func (rsaObj *RSAAccumulator)VerifyNonMembership(An big.Int,Am big.Int,x string,proof non_mem_witness){
//	x_prime,_ := util.HashToPrime(x)
//
//}

//func (rsaObj *RSAAccumulator)ProveNoMembership(key *big.Int) *big.Int{
//	v,ok := rsaObj.data[key.String()]
//	if ok{
//		return nil
//	}
//	witness := rsaObj.iterateAndGetProduct()
//
//	return big.NewInt(0)
//}

//func (rsaObj *RSAAccumulator)VerifyNoMembership(){
//
//}

func doVerifyMembership(accumulatorState *big.Int,hashPrime *big.Int,proof *big.Int,n *big.Int) bool{
	result := big.NewInt(1)
	result.Exp(proof,hashPrime,n)
	fmt.Println("当前累加器状态",accumulatorState)
	fmt.Println("当前关键字hash",hashPrime)
	fmt.Println("当前关键字存在性证明",proof)
	fmt.Println("当前result",result)
	if result.Cmp(accumulatorState)==0{
		return true
	}
	return false
}

func (rsaObj *RSAAccumulator)iterateAndGetProductWithoutX(key string) *big.Int{
	result := big.NewInt(1)
	for k,v := range rsaObj.data{
		if k!=key{
			prime := util.HashToPrimeWithNonce(k,v)
			result.Mul(result,prime)
		}
	}
	return result
}

func (rsaObj *RSAAccumulator)iterateAndGetProduct() *big.Int{
	result := big.NewInt(1)
	for k,v := range rsaObj.data{
		prime := util.HashToPrimeWithNonce(k,v)
		result.Mul(result,prime)
		result.Mul(result,prime)
	}
	return result
}

func (rsaObj *RSAAccumulator)getPair() *pair.Pair {
	return rsaObj.pair
}


func New() *RSAAccumulator {
	data := make(map[string]int)
	pair := pair.NewPair(RSA_PRIME_SIZE)
	var N = new(big.Int)
	N.Mul(pair.GetFirst(), pair.GetSecond())
	random := util.GenerateRandomNumber(*big.NewInt(0), *N)
	random2 := big.NewInt(0)
	random2.Set(random)
	return &RSAAccumulator{
		data: data,
		pair: pair,
		p:    pair.GetFirst(),
		q:    pair.GetSecond(),
		n:    N,
		a:    random,
		a0:   random2,
	}
}