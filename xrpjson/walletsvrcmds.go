package xrpjson

// GetBalanceCmd defines the getheight JSON-RPC command.
type GetAccountInfoCmd struct {
	Account     string `json:"account"`
	Strict      bool   `json:"strict"`
	LedgerIndex string `json:"ledger_index"`
	Queue       bool   `json:"queue"`
}

// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewGetAccountInfoCmd(account, ledger_index string) *GetAccountInfoCmd {
	return &GetAccountInfoCmd{
		Account:     account,
		Strict:      true,
		LedgerIndex: ledger_index,
		Queue:       false,
	}
}

// GetAccountTxCmd defines the gettransaction JSON-RPC command.
type GetAccountTxCmd struct {
	Account        string `json:"account"`
	Binary         bool   `json:"binary"`
	Forward        bool   `json:"forward"`
	LedgerIndexMax int64  `json:"ledger_index_max"`
	LedgerIndexMin int64  `json:"ledger_index_min"`
}

// NewGetAccountTxCmd returns a new instance which can be used to issue a
// tx JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewGetAccountTxCmd(account string, ledger_index int64) *GetAccountTxCmd {
	return &GetAccountTxCmd{
		Account:        account,
		Binary:         false,
		Forward:        false,
		LedgerIndexMax: ledger_index,
		LedgerIndexMin: ledger_index,
	}
}

type TxJSON struct {
	Account         string `json:"Account"`
	Amount          string `json:"Amount"`
	Destination     string `json:"Destination"`
	TransactionType string `json:"TransactionType"`
	DestinationTag  int64  `json:"DestinationTag"`
}

// SendToAddressCmd defines the sendtoaddress JSON-RPC command.
type SubmitCmd struct {
	Offline    bool   `json:"offline"`
	Secret     string `json:"secret"`
	TxJSON     TxJSON `json:"tx_json"`
	FeeMultMax int    `json:"fee_mult_max"`
}

// NewSendToAddressCmd returns a new instance which can be used to issue a
// sendtoaddress JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewSubmitCmd(account, secret, amount, destination string, destination_tag int64) *SubmitCmd {
	txJson := TxJSON{
		Account:         account,
		Amount:          amount,
		Destination:     destination,
		TransactionType: "Payment",
		DestinationTag:  destination_tag,
	}
	return &SubmitCmd{
		Offline:    false,
		Secret:     secret,
		TxJSON:     txJson,
		FeeMultMax: 1000,
	}
}

type GetTxCmd struct {
	Transaction string `json:"transaction"`
	Binary      bool   `json:"binary"`
}

func NewGetTxCmd(txid string) *GetTxCmd {
	return &GetTxCmd{
		Transaction: txid,
		Binary:      false,
	}
}

func init() {
	// The commands in this file are only usable with a wallet server.
	flags := UFWalletOnly

	MustRegisterCmd("account_info", (*GetAccountInfoCmd)(nil), flags)
	MustRegisterCmd("account_tx", (*GetAccountTxCmd)(nil), flags)
	MustRegisterCmd("submit", (*SubmitCmd)(nil), flags)
	MustRegisterCmd("tx", (*GetTxCmd)(nil), flags)
}
