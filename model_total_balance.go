/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// User's balance in all accounts
type TotalBalance struct {
	Total AccountBalance `json:"total,omitempty"`
	// Total balances in different accounts  - cross_margin: cross margin account - spot: spot account - finance: finance account - margin: margin account - quant: quant account - futures: futures account - delivery: delivery account - warrant: warrant account - cbbc: cbbc account
	Details map[string]AccountBalance `json:"details,omitempty"`
}
