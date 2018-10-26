package bean

import ()

// FormParams FormParams
type FormParams struct {
	Params      string `form:"params" json:"params"`             // params
	Key         string `form:"key" json:"key"`                   // key
	Value       string `form:"value" json:"value"`               // value
	TTL         string `form:"ttl" json:"ttl"`                   // ttl
	Address     string `form:"address" json:"address"`           //
	From        string `form:"from" json:"from"`                 //
	To          string `form:"to" json:"to"`                     //
	Amount      string `form:"amount" json:"amount"`             //
	Decimals    string `form:"decimals" json:"decimals"`         //
	Pwd         string `form:"pwd" json:"pwd"`                   //
	Name        string `form:"name" json:"name"`                 //
	Symbol      string `form:"symbol" json:"symbol"`             //
	Total       string `form:"total" json:"total"`               //
	Desc        string `form:"desc" json:"desc"`                 //
	IsWait      string `form:"iswait" json:"iswait"`             //
	Number      string `form:"number" json:"number"`             //
	CallAddress string `form:"call_address" json:"call_address"` //
	Mnemonic    string `form:"mnemonic" json:"mnemonic"`         //
	Path        string `form:"path" json:"path"`                 //
}
