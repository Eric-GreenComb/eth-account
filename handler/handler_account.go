package handler

import (
	"fmt"
	"net/http"
	"strings"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	// "github.com/ethereum/go-ethereum/rlp"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/eth-account/badger"
	"github.com/Eric-GreenComb/eth-account/bean"
	"github.com/Eric-GreenComb/eth-account/ethereum"
	"github.com/Eric-GreenComb/go-bip39"
)

// CreateAccount CreateAccount
func CreateAccount(c *gin.Context) {

	// _name := c.Params.ByName("name")
	_passphrase := c.Params.ByName("passphrase")

	_key, err := ethereum.Ks.NewKey()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	fmt.Println(_key.Address.String())
	keyjson, err := ethereum.Ks.GenKeystore(_key, _passphrase)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.String(http.StatusOK, string(keyjson))
}

// CreateBIP39 CreateBIP39
func CreateBIP39(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_seed := bip39.NewSeed(_formParams.Mnemonic, "")

	// fmt.Println(string(_seed))

	_wallet, err := ethereum.NewFromSeed(_seed)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_path := ethereum.MustParseDerivationPath(_formParams.Path)
	_account, err := _wallet.Derive(_path, false)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_pubKey, _ := _wallet.PublicKeyHex(_account)

	_privateKey, _ := _wallet.PrivateKeyHex(_account)

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "addr": _account.Address.Hex(), "pubkey": _pubKey, "privatekey": _privateKey})
}

// CreateBIP39Keysore CreateBIP39Keysore
func CreateBIP39Keysore(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_seed := bip39.NewSeed(_formParams.Mnemonic, "")

	_wallet, err := ethereum.NewFromSeed(_seed)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_path := ethereum.MustParseDerivationPath(_formParams.Path)
	_account, err := _wallet.Derive(_path, false)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_privateKey, _ := _wallet.PrivateKey(_account)

	_key := ethereum.Ks.GenKeyFromECDSA(_privateKey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_passphrase := _formParams.Pwd
	keyjson, err := ethereum.Ks.GenKeystore(_key, _passphrase)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	fmt.Println(strings.ToLower(_account.Address.Hex()))
	err = badger.NewWrite().Set(strings.ToLower(_account.Address.Hex()), keyjson)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": string(keyjson)})
	c.String(http.StatusOK, string(keyjson))
}

// GetAddressByPriv GetAddressByPriv
func GetAddressByPriv(c *gin.Context) {

	var _param bean.FormParams
	c.BindJSON(&_param)

	_priv, err := crypto.HexToECDSA(_param.Value)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_addr := crypto.PubkeyToAddress(_priv.PublicKey)
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "addr": _addr.Hex()})
}

// SignAndRecover SignAndRecover
func SignAndRecover(c *gin.Context) {

	var _param bean.FormParams
	c.BindJSON(&_param)

	_seed := bip39.NewSeed(_param.Mnemonic, "")

	_wallet, err := ethereum.NewFromSeed(_seed)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_path := ethereum.MustParseDerivationPath(_param.Path)
	_account, err := _wallet.Derive(_path, false)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_pubKey, _ := _wallet.PublicKeyHex(_account)

	_privateKey, _ := _wallet.PrivateKeyHex(_account)

	fmt.Println("PubKey", _pubKey)
	fmt.Println("PrivKey", _privateKey)

	// secp256k1.RecoverPubkey
	_msg := "kdsjfkajoiqhkfi189ihsfaklsjfk222"
	_bytePrivKey, err := _wallet.PrivateKeyBytes(_account)
	if err != nil {
		fmt.Println(err.Error())
	}
	_sig, err := secp256k1.Sign([]byte(_msg), _bytePrivKey)
	if err != nil {
		fmt.Println(err.Error())
	}
	pubkey2, err := secp256k1.RecoverPubkey([]byte(_msg), _sig)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("PubKey2", hexutil.Encode(pubkey2)[:68])

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "addr": _account.Address.Hex(), "pubkey": _pubKey, "privatekey": _privateKey})
}
