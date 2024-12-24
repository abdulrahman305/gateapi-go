/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Order information that needs to be modified
type BatchAmendItem struct {
	// The order ID returned upon successful creation or the custom ID specified by the user during creation (i.e., the 'text' field).
	OrderId string `json:"order_id"`
	// Currency pair
	CurrencyPair string `json:"currency_pair"`
	// Default to spot, portfolio, and margin accounts if not specified. Use 'cross_margin' to query cross margin accounts. Only 'cross_margin' can be specified for portfolio margin accounts.
	Account string `json:"account,omitempty"`
	// trade amount, only one of amount and price can be specified
	Amount string `json:"amount,omitempty"`
	// trade price, only one of amount and price can be specified
	Price string `json:"price,omitempty"`
	// Custom info during amending order
	AmendText string `json:"amend_text,omitempty"`
	// Processing Mode: When placing an order, different fields are returned based on action_mode. This field is only valid during the request and is not included in the response result ACK: Asynchronous mode, only returns key order fields RESULT: No clearing information FULL: Full mode (default)
	ActionMode string `json:"action_mode,omitempty"`
}
