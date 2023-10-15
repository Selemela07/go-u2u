package internaltx

import (
	"github.com/unicornultrafoundation/go-u2u/common"
	"github.com/unicornultrafoundation/go-u2u/core/types"
)

func IsInternal(tx *types.Transaction) bool {
	v, r, _ := tx.RawSignatureValues()
	return v.Sign() == 0 && r.Sign() == 0
}

func InternalSender(tx *types.Transaction) common.Address {
	_, _, s := tx.RawSignatureValues()
	return common.BytesToAddress(s.Bytes())
}

func Sender(signer types.Signer, tx *types.Transaction) (common.Address, error) {
	if !IsInternal(tx) {
		return types.Sender(signer, tx)
	}
	return InternalSender(tx), nil
}
