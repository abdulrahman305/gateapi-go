/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Order cancellation result
type CancelOrderResult struct {
	// Order currency pair
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Order ID
	Id string `json:"id,omitempty"`
	// Custom order information
	Text string `json:"text,omitempty"`
	// Whether cancellation succeeded
	Succeeded bool `json:"succeeded,omitempty"`
	// Error label when failed to cancel the order; emtpy if succeeded
	Label string `json:"label,omitempty"`
	// Error message when failed to cancel the order; empty if succeeded
	Message string `json:"message,omitempty"`
	// Empty by default. If cancelled order is cross margin order, this field is set to `cross_margin`
	Account string `json:"account,omitempty"`
}
