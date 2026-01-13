package signer

import "github.com/ivanzzeth/ethsig"

type EOATradingSigner interface {
	ethsig.TypedDataSigner
	ethsig.PersonalSigner
	TransactionSignerAndAddrGetter
}
