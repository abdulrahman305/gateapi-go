/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type PositionClose struct {
	// Position close time
	Time float64 `json:"time,omitempty"`
	// Futures contract
	Contract string `json:"contract,omitempty"`
	// Position side, long or short
	Side string `json:"side,omitempty"`
	// PNL
	Pnl string `json:"pnl,omitempty"`
	// PNL - Position P/L
	PnlPnl string `json:"pnl_pnl,omitempty"`
	// PNL - Funding Fees
	PnlFund string `json:"pnl_fund,omitempty"`
	// PNL - Transaction Fees
	PnlFee string `json:"pnl_fee,omitempty"`
	// Text of close order
	Text string `json:"text,omitempty"`
	// Max Trade Size
	MaxSize string `json:"max_size,omitempty"`
	// Cumulative closed position volume
	AccumSize string `json:"accum_size,omitempty"`
	// First Open Time
	FirstOpenTime int64 `json:"first_open_time,omitempty"`
	// When 'side' is 'long,' it indicates the opening average price; when 'side' is 'short,' it indicates the closing average price.
	LongPrice string `json:"long_price,omitempty"`
	// When 'side' is 'long,' it indicates the opening average price; when 'side' is 'short,' it indicates the closing average price
	ShortPrice string `json:"short_price,omitempty"`
}
