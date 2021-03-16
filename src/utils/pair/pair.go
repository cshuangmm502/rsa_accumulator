package pair

import (
	util2 "awesomeProject/src/utils/util"
	"fmt"
	"math/big"
)

type Pair struct {
	first *big.Int
	second *big.Int
}

func (p *Pair) GetFirst() *big.Int{
	return p.first
}

func (p *Pair) GetSecond() *big.Int {
	return p.second
}

func NewPair(bitLength int) *Pair{
	var p1 *big.Int
	var p2 *big.Int

	p1 = util2.GenerateLargePrime(bitLength)
	p2 = util2.GenerateLargePrime(bitLength)
	for p:=p1;p.String()==p2.String();p2=util2.GenerateLargePrime(bitLength){
	}

	return &Pair{first:p1,second:p2}
}

func (p *Pair)Print()  {
	fmt.Println(p.first.String())
	fmt.Println(p.second.String())
}

//func HashToPrime(x big.Int,bitLength int) Pair {
//	return HashToPrime(x,bitLength,big.NewInt(0))
//}

//hash x to a prime and get offset, thus x+nonce will be a large prime with certainty PRIME_CERTAINTY
//func HashToPrime(x *big.Int) *Pair {
//	result := util2.HashToPrime(x)
//	return &Pair{result,}
//}

//hash x to a designated bitLength
func HashToLength(x big.Int,bitLength int) *big.Int{
	return big.NewInt(0)
}