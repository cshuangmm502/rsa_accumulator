package accumulator

import (
	"math/big"
)

//const (
//	RSA_KEY_SIZE = 3072
//	RSA_PRIME_SIZE = RSA_KEY_SIZE/2
//	ACCUMULATED_PRIME_SIZE = 128
//)

type Accumulator interface {
	AddMember(bigInteger *big.Int) *big.Int
	ProveMembership(bigInteger big.Int) *big.Int
	DeleteMember(bigInteger big.Int) *big.Int
	VerifyMembership(bigInteger big.Int,proof big.Int) bool
}


//type RSAAccumulator struct {
//	data 	map[string]big.Int
//	pair 	pair.Pair
//	p		big.Int
//	q 		big.Int
//	N		big.Int
//	random 	int
//	A0		big.Int
//	A		big.Int
//}
//
//func (rsaObj *RSAAccumulator)getN() big.Int {
//	return rsaObj.N
//}
//
//func (rsaObj *RSAAccumulator)getVal(bigInteger big.Int) big.Int {
//	return rsaObj.data[bigInteger.String()]
//}
//
//func (rsaObj *RSAAccumulator)addMember(bigInteger big.Int) *big.Int {
//
//	return big.NewInt(0)
//}
//
//func (rsaObj *RSAAccumulator)proveMembership(bigInteger big.Int) big.Int {
//	return *big.NewInt(0)
//}
//
//func (rsaObj *RSAAccumulator)iterateAndGetProduct(bigInteger big.Int) *big.Int{
//	return big.NewInt(0)
//}
//
//func (rsaObj *RSAAccumulator)deleteMember(bigInteger big.Int) *big.Int{
//	return big.NewInt(0)
//}
//
//func (rsaObj *RSAAccumulator)verifyMembership(bigInteger big.Int) bool{
//	return true
//}

//func (rsaObj *RSAAccumulator)doVerifyMembership()

