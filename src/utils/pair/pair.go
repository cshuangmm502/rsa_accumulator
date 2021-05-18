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

