/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Options contract detail
type OptionsContract struct {
	// Options contract name
	Name string `json:"name,omitempty"`
	// tag
	Tag string `json:"tag,omitempty"`
	// Creation time
	CreateTime float64 `json:"create_time,omitempty"`
	// Expiration time
	ExpirationTime float64 `json:"expiration_time,omitempty"`
	// `true` means call options, while `false` is put options
	IsCall bool `json:"is_call,omitempty"`
	// Multiplier used in converting from invoicing to settlement currency
	Multiplier string `json:"multiplier,omitempty"`
	// Underlying
	Underlying string `json:"underlying,omitempty"`
	// Underlying price (quote currency)
	UnderlyingPrice string `json:"underlying_price,omitempty"`
	// Last trading price
	LastPrice string `json:"last_price,omitempty"`
	// Current mark price (quote currency)
	MarkPrice string `json:"mark_price,omitempty"`
	// Current index price (quote currency)
	IndexPrice string `json:"index_price,omitempty"`
	// Maker fee rate, where negative means rebate
	MakerFeeRate string `json:"maker_fee_rate,omitempty"`
	// Taker fee rate
	TakerFeeRate string `json:"taker_fee_rate,omitempty"`
	// Minimum order price increment
	OrderPriceRound string `json:"order_price_round,omitempty"`
	// Minimum mark price increment
	MarkPriceRound string `json:"mark_price_round,omitempty"`
	// Minimum order size the contract allowed
	OrderSizeMin int64 `json:"order_size_min,omitempty"`
	// Maximum order size the contract allowed
	OrderSizeMax int64 `json:"order_size_max,omitempty"`
	// The positive and negative offset allowed between the order price and the current mark price, that is, the order price `order_price` must meet the following conditions:   order_price is within the range of mark_price +/- order_price_deviate * underlying_price  and does not distinguish between buy and sell orders
	OrderPriceDeviate string `json:"order_price_deviate,omitempty"`
	// Referral fee rate discount
	RefDiscountRate string `json:"ref_discount_rate,omitempty"`
	// Referrer commission rate
	RefRebateRate string `json:"ref_rebate_rate,omitempty"`
	// Current orderbook ID
	OrderbookId int64 `json:"orderbook_id,omitempty"`
	// Current trade ID
	TradeId int64 `json:"trade_id,omitempty"`
	// Historical accumulated trade size
	TradeSize int64 `json:"trade_size,omitempty"`
	// Current total long position size
	PositionSize int64 `json:"position_size,omitempty"`
	// Maximum number of open orders
	OrdersLimit int32 `json:"orders_limit,omitempty"`
}
