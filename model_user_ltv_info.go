/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// User's currency statistics data
type UserLtvInfo struct {
	// Collateral
	CollateralCurrency string `json:"collateral_currency,omitempty"`
	// Borrowed currency
	BorrowCurrency string `json:"borrow_currency,omitempty"`
	// The initial collateralization rate
	InitLtv string `json:"init_ltv,omitempty"`
	// Warning collateralization ratio
	AlertLtv string `json:"alert_ltv,omitempty"`
	// The liquidation collateralization rate
	LiquidateLtv string `json:"liquidate_ltv,omitempty"`
	// Minimum borrowable amount for the loan currency
	MinBorrowAmount string `json:"min_borrow_amount,omitempty"`
	// Remaining borrowable amount for the loan currency
	LeftBorrowableAmount string `json:"left_borrowable_amount,omitempty"`
}
