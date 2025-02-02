package polka

import (
	"encoding/json"
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/vedhavyas/go-subkey"
	"github.com/vedhavyas/go-subkey/sr25519"
	"githup.com/youthonline/ndk/base"
)

type Account struct {
	keypair  *signature.KeyringPair
	keystore *keystore

	publicKey []byte
	address   string

	Network int
}

// func KeyPairFromSeed(seed string, network int) (*signature.KeyringPair, error) {
// 	keyringPair, err := signature.KeyringPairFromSecret(seed, uint8(network))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &keyringPair, nil
// }

// func KeyPairFromSeed(prikey string, network int) (*Account, error) {
// 	seed, err := types.HexDecodeString(prikey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	kyr, err := sr25519.Scheme{}.FromSeed(seed)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ss58Address, err := kyr.SS58Address(uint8(network))
// 	if err != nil {
// 		return nil, err
// 	}
// 	var pk = kyr.Public()
// 	 signature.KeyringPair{
// 		URI:       prikey,
// 		Address:   ss58Address,
// 		PublicKey: pk,
// 	}, nil
// }

func NewAccountWithMnemonic(mnemonic string, network int) (*Account, error) {
	if len(mnemonic) == 0 {
		return nil, ErrSeedOrPhrase
	}
	keyringPair, err := signature.KeyringPairFromSecret(mnemonic, uint8(network))
	if err != nil {
		return nil, err
	}

	return &Account{
		keypair:   &keyringPair,
		publicKey: keyringPair.PublicKey,
		address:   keyringPair.Address,
		Network:   network,
	}, nil
}

func AccountWithPrivateKey(prikey string, network int) (*Account, error) {
	seed, err := types.HexDecodeString(prikey)
	if err != nil {
		return nil, err
	}
	kyr, err := sr25519.Scheme{}.FromSeed(seed)
	if err != nil {
		return nil, err
	}

	ss58Address, err := kyr.SS58Address(uint8(network))
	if err != nil {
		return nil, err
	}
	var pk = kyr.Public()
	keypair := signature.KeyringPair{
		URI:       prikey,
		Address:   ss58Address,
		PublicKey: pk,
	}

	return &Account{
		keypair:   &keypair,
		publicKey: keypair.PublicKey,
		address:   keypair.Address,
		Network:   network,
	}, nil
}

func NewAccountWithKeystore(keystoreString, password string, network int) (*Account, error) {
	var keyStore keystore
	err := json.Unmarshal([]byte(keystoreString), &keyStore)
	if err != nil {
		return nil, err
	}
	if err = keyStore.CheckPassword(password); err != nil {
		return nil, err
	}

	publicKeyHex, err := DecodeAddressToPublicKey(keyStore.Address)
	if err != nil {
		return nil, err
	}
	publicKey, err := types.HexDecodeString(publicKeyHex)
	if err != nil {
		return nil, err
	}
	address, err := EncodePublicKeyToAddress(publicKeyHex, network)
	if err != nil {
		return nil, err
	}

	return &Account{
		keystore:  &keyStore,
		publicKey: publicKey,
		address:   address,
		Network:   network,
	}, nil
}

func (a *Account) DeriveAccountAt(network int) (*Account, error) {
	address, err := EncodePublicKeyToAddress(a.PublicKeyHex(), network)
	if err != nil {
		return nil, err
	}
	return &Account{
		keypair:   a.keypair,
		keystore:  a.keystore,
		publicKey: a.publicKey,
		address:   address,
		Network:   network,
	}, nil
}

// MARK - Implement the protocol wallet.Account

// @return privateKey data
func (a *Account) PrivateKey() ([]byte, error) {
	if a.keypair == nil {
		return nil, ErrNilKey
	}

	scheme := sr25519.Scheme{}
	kyr, err := subkey.DeriveKeyPair(scheme, a.keypair.URI)
	if err != nil {
		return nil, err
	}
	return kyr.Seed(), nil
}

// @return privateKey string that will start with 0x.
func (a *Account) PrivateKeyHex() (string, error) {
	data, err := a.PrivateKey()
	if err != nil {
		return "", err
	}
	return types.HexEncodeToString(data), nil
}

// @return publicKey data
func (a *Account) PublicKey() []byte {
	return a.publicKey
}

// @return publicKey string that will start with 0x.
func (a *Account) PublicKeyHex() string {
	return types.HexEncodeToString(a.publicKey)
}

// @return address string
func (a *Account) Address() string {
	return a.address
}

func (a *Account) Sign(message []byte, password string) (data []byte, err error) {
	defer func() {
		errPanic := recover()
		if errPanic != nil {
			err = ErrSign
			fmt.Println(errPanic)
			return
		}
	}()
	if a.keypair != nil {
		data, err := signature.Sign(message, a.keypair.URI)
		return data, err // Must be separate to ensure that err can catch panic
	} else if a.keystore != nil {
		data, err := a.keystore.Sign(message, password)
		return data, err
	}
	return nil, ErrNilWallet
}

func (a *Account) SignHex(messageHex string, password string) (*base.OptionalString, error) {
	bytes, err := types.HexDecodeString(messageHex)
	if err != nil {
		return nil, err
	}
	signed, err := a.Sign(bytes, password)
	if err != nil {
		return nil, err
	}
	signedString := types.HexEncodeToString(signed)
	return &base.OptionalString{Value: signedString}, nil
}

// 内置账号，主要用来给用户未签名的交易签一下名
// 然后给用户去链上查询手续费，保护用户资产安全
func mockAccount() *Account {
	mnemonic := "infant carbon above canyon corn collect finger drip area feature mule autumn"
	a, _ := NewAccountWithMnemonic(mnemonic, 44)
	return a
}

// MARK - Implement the protocol AddressUtil

// @param publicKey can start with 0x or not.
func (a *Account) EncodePublicKeyToAddress(publicKey string) (string, error) {
	return EncodePublicKeyToAddress(publicKey, a.Network)
}

// @return publicKey that will start with 0x.
func (a *Account) DecodeAddressToPublicKey(address string) (string, error) {
	return DecodeAddressToPublicKey(address)
}

func (a *Account) IsValidAddress(address string) bool {
	return IsValidAddress(address)
}

func AsPolkaAccount(account base.Account) *Account {
	return account.(*Account)
}
