//go:build !windows
// +build !windows

package keystore

import (
	"github.com/eliohn/gotron-sdk/pkg/address"
	"github.com/ethereum/go-ethereum/crypto"
)

func RecoverPubkey(hash []byte, signature []byte) (address.Address, error) {

	if signature[64] >= 27 {
		signature[64] -= 27
	}

	sigPublicKey, err := crypto.Ecrecover(hash, signature)
	if err != nil {
		return nil, err
	}
	pubKey, err := UnmarshalPublic(sigPublicKey)
	if err != nil {
		return nil, err
	}

	addr := address.PubkeyToAddress(*pubKey)
	return addr, nil
}
