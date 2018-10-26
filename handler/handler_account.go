package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

	c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": string(keyjson)})
}
