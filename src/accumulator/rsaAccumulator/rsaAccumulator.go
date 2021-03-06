package rsaAccumulator

import (
	"awesomeProject/src/utils/pair"
	"awesomeProject/src/utils/util"
	"fmt"
	"math/big"
)

const (
	RSA_KEY_SIZE = 32
	RSA_PRIME_SIZE = RSA_KEY_SIZE/2
	ACCUMULATED_PRIME_SIZE = 128
)

//type RsaInter interface {
//	accumulator.Accumulator
//	GetN() big.Int
//}

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
	fmt.Println("?????????????????????",accumulatorState)
	fmt.Println("???????????????hash",hashPrime)
	fmt.Println("??????????????????????????????",proof)
	fmt.Println("??????result",result)
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