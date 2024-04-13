// Package models -> is package for describe all models, interfaces, data types, schemas, etc.
package models

// -> API response types below -<

type WalletItemResponse struct {
	UserID      uint64
	Address     string
	CoinName    string
	CoinBalance string  // *big.Float => balance.Text('f', -1)
	FiatValue   float64 // a balance in retreived currency type
	// update? *
}

type SendTransactionRequestDto struct {
	SenderAddress    string `json:"senderAddress" xml:"senderAddress" form:"senderAddress"`
	CoinName         string `json:"coinName" xml:"coinName" form:"coinName"`
	RecipientAddress string `json:"recipientAddress" xml:"recipientAddress" form:"recipientAddress"`
	Amount           string `json:"amount" xml:"amount" form:"amount"`
}

// WalletStat -> for the mongoDB for fast access from client API
// without ctypto API call
type WalletStats struct {
	ID           string  `bson:"_id,omitempty"`
	UserEmail    string  `json:"userEmail" bson:"userEmail"`       // receive via request
	UserId       int64   `json:"userId" bson:"userId"`             // receive via request
	TotalBalance float64 `json:"totalBalance" bson:"totalBalance"` // default 0 -> total user wallet balance (float value 0.00) displayed in chosen fiat currency
	CurrencyType string  `json:"currencyType" bson:"currencyType"` // default USD -> ["AUD","USD", "EUR", "AED", "RUB"]
	UpdatedAt    int64   `json:"updatedAt" bson:"updatedAt"`       // update once a 300 ms OR if manipulating wallet by user <-
}

// WalletCoinItem- > describe wallet item types
type WalletCoinItem struct {
	UserId      int64   `json:"userId" bson:"userId"` // receive via request
	CoinName    string  `json:"coinName" bson:"coinName"`
	Address     string  `json:"address" bson:"address"`
	CoinBalance float64 `json:"coinBalance" bson:"coinBalance"` // as float (0.00)
	FiatBalance float32 `json:"fiatBalance" bson:"fiatBalance"` // in chosen currency type ["AUD","USD", "EUR", "AED", "RUB"]
	WalletId    string  `json:"walletId" bson:"walletId"`       // ref to the WalletStat _id ->
}

type ParserCache struct {
	Status    string // current status
	UpdatedAt int64  // last status update
	StartFrom int64  // date to parse from
}

type BtcWallet struct {
	ID              int    `json:"id" bson:"id, omitempty"` // auto increment value
	Address         string `json:"address" bson:"address"`
	PrivateKey      string `json:"privateKey" bson:"privateKey"`
	PublicKey       string `json:"publicKey" bson:"publicKey"`
	Wif             string `json:"wif" bson:"wif"`
	PubKeys         string `json:"pubKeys" bson:"pubKeys"`
	ScriptType      string `json:"scriptType" bson:"scriptType"`
	OriginalAddress string `json:"originalAddress" bson:"originalAddress"`
	OAPAddress      string `json:"OAPAddress" bson:"OAPAddress"`
	CreatedAt       int64  `json:"createdAt" bson:"createdAt"` // stamp in ms as value
	UpdatedAt       int64  `json:"updatedAt" bson:"updatedAt"` // stamp in ms as value
	UserID          int64  `json:"userId" bson:"userId"`       // receive via request (telegram chatId)
}

type TonWallet struct {
	ID         int    `json:"id" bson:"id, omitempty"`      // auto increment value
	Address    string `json:"address" bson:"address"`       // address string
	AddrType   int    `json:"addrType" bson:"addrType"`     //
	PrivateKey string `json:"privateKey" bson:"privateKey"` // private
	BitsLen    int    `json:"bitsLen" bson:"bitsLen"`

	// Can be used to operate multiple wallets with the same key and version.
	// use GetSubwallet if you need it.
	Subwallet uint32 `json:"subwallet" bson:"subwallet"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"` // stamp in ms as value
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"` // stamp in ms as value
	UserID    int64  `json:"userId" bson:"userId"`       // receive via request (telegram chatId)
}