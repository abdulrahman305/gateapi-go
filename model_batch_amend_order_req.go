/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Modify contract order parameters
type BatchAmendOrderReq struct {
	// Order id, order_id and text must contain at least one
	OrderId int64 `json:"order_id,omitempty"`
	// User-defined order text, at least one of order_id and text must be passed
	Text string `json:"text,omitempty"`
	// The new order size, including the executed order size. - If it is less than or equal to the executed quantity, the order will be cancelled. - The new order direction must be consistent with the original one. - The size of the closing order cannot be modified. - For orders that only reduce positions, if the size is increased, other orders that only reduce positions may be kicked out. - If the price is not modified, reducing the size will not affect the depth of the queue, and increasing the size will place it at the end of the current price.
	Size int64 `json:"size,omitempty"`
	// New order price.
	Price string `json:"price,omitempty"`
	// Custom info during amending order
	AmendText string `json:"amend_text,omitempty"`
}
