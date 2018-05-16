package ripplerpc

import (
	"encoding/json"
	"github.com/GoChainRpc/ripplerpc/xrpjson"
)

// FutureGetAccountTxResult is a future promise to deliver the result
// of a GetTransactionAsync RPC invocation (or an applicable error).
type FutureGetAccountTxResult chan *response

// Receive waits for the response promised by the future and returns detailed
// information about a wallet transaction.
func (r FutureGetAccountTxResult) Receive() (*xrpjson.GetAccountTxResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}
	// Unmarshal result as a gettransaction result object
	var getTx xrpjson.GetAccountTxResult
	err = json.Unmarshal(res, &getTx)
	if err != nil {
		return nil, err
	}

	return &getTx, nil
}

// GetAccountTxAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetTransaction for the blocking version and more details.
func (c *Client) GetAccountTxAsync(account string, ledger_index int64) FutureGetAccountTxResult {
	cmd := xrpjson.NewGetAccountTxCmd(account, ledger_index)
	return c.sendCmd(cmd)
}

// GetTransfers returns detailed information about a wallet transaction.
//
// See GetTransfers to return the raw transaction instead.
func (c *Client) GetAccountTx(account string, ledger_index int64) (*xrpjson.GetAccountTxResult, error) {
	return c.GetAccountTxAsync(account, ledger_index).Receive()
}

type FutureSubmitResult chan *response

func (r FutureSubmitResult) Receive() (*xrpjson.SubmitResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var transferResult xrpjson.SubmitResult
	err = json.Unmarshal(res, &transferResult)
	if err != nil {
		return nil, err
	}

	return &transferResult, nil
}

func (c *Client) SubmitAsync(account, secret, amount, destination string, destination_tag int64) FutureSubmitResult {
	cmd := xrpjson.NewSubmitCmd(account, secret, amount, destination, destination_tag)
	return c.sendCmd(cmd)
}

func (c *Client) Submit(account, secret, amount, destination string, destination_tag int64) (*xrpjson.SubmitResult, error) {
	return c.SubmitAsync(account, secret, amount, destination, destination_tag).Receive()
}

// FutureGetBalanceResult is a future promise to deliver the result of a
// GetBalanceAsync or GetBalanceMinConfAsync RPC invocation (or an applicable
// error).
type FutureGetAccountInfoResult chan *response

// Receive waits for the response promised by the future and returns the
// available balance from the server for the specified account.
func (r FutureGetAccountInfoResult) Receive() (*xrpjson.GetAccountInfoResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string
	var accountInfoResult xrpjson.GetAccountInfoResult
	err = json.Unmarshal(res, &accountInfoResult)
	if err != nil {
		return nil, err
	}

	return &accountInfoResult, nil
}

// GetBalanceAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetBalance for the blocking version and more details.
func (c *Client) GetAccountInfoAsync(account, ledger_index string) FutureGetAccountInfoResult {
	cmd := xrpjson.NewGetAccountInfoCmd(account, ledger_index)
	return c.sendCmd(cmd)
}

// GetBalance returns the available balance from the server for the specified
// account using the default number of minimum confirmations.  The account may
// be "*" for all accounts.
//
// See GetBalanceMinConf to override the minimum number of confirmations.
func (c *Client) GetAccountInfo(account, ledger_index string) (*xrpjson.GetAccountInfoResult, error) {
	return c.GetAccountInfoAsync(account, ledger_index).Receive()
}

type FutureTxResult chan *response

func (r FutureTxResult) Receive() (*xrpjson.GetTxResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var transferResult xrpjson.GetTxResult
	err = json.Unmarshal(res, &transferResult)
	if err != nil {
		return nil, err
	}

	return &transferResult, nil
}

func (c *Client) GetTxAsync(txid string) FutureTxResult {
	cmd := xrpjson.NewGetTxCmd(txid)
	return c.sendCmd(cmd)
}

func (c *Client) GetTx(txid string) (*xrpjson.GetTxResult, error) {
	return c.GetTxAsync(txid).Receive()
}
