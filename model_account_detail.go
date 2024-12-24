/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Account detail
type AccountDetail struct {
	// IP whitelist
	IpWhitelist []string `json:"ip_whitelist,omitempty"`
	// CurrencyPair whitelisting
	CurrencyPairs []string `json:"currency_pairs,omitempty"`
	// User ID
	UserId int64 `json:"user_id,omitempty"`
	// User VIP level
	Tier int64            `json:"tier,omitempty"`
	Key  AccountDetailKey `json:"key,omitempty"`
	// User role: 0 - Ordinary user 1 - Order leader 2 - Follower 3 - Order leader and follower
	CopyTradingRole int32 `json:"copy_trading_role,omitempty"`
}
