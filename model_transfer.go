/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Accounts available to transfer:  - `spot`: spot account - `margin`: margin account - `futures`: perpetual futures account - `delivery`: delivery futures account - `options`: options account
type Transfer struct {
	// Transfer currency. For futures account, `currency` can be set to `POINT` or settle currency
	Currency string `json:"currency"`
	// Account to transfer from
	From string `json:"from"`
	// Account to transfer to
	To string `json:"to"`
	// Transfer amount
	Amount string `json:"amount"`
	// Margin currency pair. Required if transfer from or to margin account
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Futures settle currency. Required if transferring from or to futures account
	Settle string `json:"settle,omitempty"`
}
