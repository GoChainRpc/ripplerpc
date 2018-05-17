package xrpjson

type GetAccountInfoResult struct {
	AccountData struct {
		Account           string `json:"Account"`
		Balance           string `json:"Balance"`
		Flags             int    `json:"Flags"`
		LedgerEntryType   string `json:"LedgerEntryType"`
		OwnerCount        int    `json:"OwnerCount"`
		PreviousTxnID     string `json:"PreviousTxnID"`
		PreviousTxnLgrSeq int64  `json:"PreviousTxnLgrSeq"`
		Sequence          int64  `json:"Sequence"`
		Index             string `json:"index"`
	} `json:"account_data"`
	LedgerCurrentIndex int64 `json:"ledger_current_index"`
	QueueData struct {
		TxnCount int `json:"txn_count"`
	} `json:"queue_data"`
	Status    string `json:"status"`
	Validated bool   `json:"validated"`
}

type SubmitResult struct {
	EngineResult        string `json:"engine_result"`
	EngineResultCode    int    `json:"engine_result_code"`
	EngineResultMessage string `json:"engine_result_message"`
	Status              string `json:"status"`
	TxBlob              string `json:"tx_blob"`
	TxJSON struct {
		Account         string `json:"Account"`
		Amount          string `json:"Amount"`
		Destination     string `json:"Destination"`
		DestinationTag  int    `json:"DestinationTag"`
		Fee             string `json:"Fee"`
		Flags           int64  `json:"Flags"`
		Sequence        int    `json:"Sequence"`
		SigningPubKey   string `json:"SigningPubKey"`
		TransactionType string `json:"TransactionType"`
		TxnSignature    string `json:"TxnSignature"`
		Hash            string `json:"hash"`
	} `json:"tx_json"`
}

type GetTxResult struct {
	Account         string `json:"Account"`
	Amount          string `json:"Amount"`
	Destination     string `json:"Destination"`
	DestinationTag  int    `json:"DestinationTag"`
	Fee             string `json:"Fee"`
	Flags           int64  `json:"Flags"`
	Sequence        int    `json:"Sequence"`
	SigningPubKey   string `json:"SigningPubKey"`
	TransactionType string `json:"TransactionType"`
	TxnSignature    string `json:"TxnSignature"`
	Date            int    `json:"date"`
	Hash            string `json:"hash"`
	InLedger        int    `json:"inLedger"`
	LedgerIndex     int    `json:"ledger_index"`
	Meta struct {
		AffectedNodes []struct {
			ModifiedNode struct {
				FinalFields struct {
					Account    string `json:"Account"`
					Balance    string `json:"Balance"`
					Flags      int    `json:"Flags"`
					OwnerCount int    `json:"OwnerCount"`
					Sequence   int    `json:"Sequence"`
				} `json:"FinalFields"`
				LedgerEntryType string `json:"LedgerEntryType"`
				LedgerIndex     string `json:"LedgerIndex"`
				PreviousFields struct {
					Balance string `json:"Balance"`
				} `json:"PreviousFields"`
				PreviousTxnID     string `json:"PreviousTxnID"`
				PreviousTxnLgrSeq int    `json:"PreviousTxnLgrSeq"`
			} `json:"ModifiedNode"`
		} `json:"AffectedNodes"`
		TransactionIndex  int    `json:"TransactionIndex"`
		TransactionResult string `json:"TransactionResult"`
		DeliveredAmount   string `json:"delivered_amount"`
	} `json:"meta"`
	Status    string `json:"status"`
	Validated bool   `json:"validated"`
}

type GetAccountTxResult struct {
	Account        string `json:"account"`
	LedgerIndexMax int    `json:"ledger_index_max"`
	LedgerIndexMin int    `json:"ledger_index_min"`
	Status         string `json:"status"`
	Transactions []struct {
		Meta struct {
			AffectedNodes []struct {
				ModifiedNode struct {
					FinalFields struct {
						Account    string `json:"Account"`
						Balance    string `json:"Balance"`
						Flags      int    `json:"Flags"`
						OwnerCount int    `json:"OwnerCount"`
						Sequence   int    `json:"Sequence"`
					} `json:"FinalFields"`
					LedgerEntryType string `json:"LedgerEntryType"`
					LedgerIndex     string `json:"LedgerIndex"`
					PreviousFields struct {
						Balance string `json:"Balance"`
					} `json:"PreviousFields"`
					PreviousTxnID     string `json:"PreviousTxnID"`
					PreviousTxnLgrSeq int    `json:"PreviousTxnLgrSeq"`
				} `json:"ModifiedNode"`
			} `json:"AffectedNodes"`
			TransactionIndex  int    `json:"TransactionIndex"`
			TransactionResult string `json:"TransactionResult"`
			DeliveredAmount   string `json:"delivered_amount"`
		} `json:"meta"`
		Tx struct {
			Account         string `json:"Account"`
			Amount          string `json:"Amount"`
			Destination     string `json:"Destination"`
			DestinationTag  int64  `json:"DestinationTag"`
			Fee             string `json:"Fee"`
			Flags           int64  `json:"Flags"`
			Sequence        int    `json:"Sequence"`
			SigningPubKey   string `json:"SigningPubKey"`
			TransactionType string `json:"TransactionType"`
			TxnSignature    string `json:"TxnSignature"`
			Date            int64  `json:"date"`
			Hash            string `json:"hash"`
			InLedger        int    `json:"inLedger"`
			LedgerIndex     int64  `json:"ledger_index"`
		} `json:"tx"`
		Validated bool `json:"validated"`
	} `json:"transactions"`
}
