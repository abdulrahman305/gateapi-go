/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Options order detail
type OptionsOrder struct {
	// Options order ID
	Id int64 `json:"id,omitempty"`
	// User ID
	User int32 `json:"user,omitempty"`
	// Creation time of order
	CreateTime float64 `json:"create_time,omitempty"`
	// Order finished time. Not returned if order is open
	FinishTime float64 `json:"finish_time,omitempty"`
	// 结束方式，包括：  - filled: 完全成交 - cancelled: 用户撤销 - liquidated: 强制平仓撤销 - ioc: 未立即完全成交，因为tif设置为ioc - auto_deleveraged: 自动减仓撤销 - reduce_only: 增持仓位撤销，因为设置reduce_only或平仓 - position_closed: 因为仓位平掉了，所以挂单被撤掉 - reduce_out: 只减仓被排除的不容易成交的挂单 - mmp_cancelled: MMP撤销
	FinishAs string `json:"finish_as,omitempty"`
	// Order status  - `open`: waiting to be traded - `finished`: finished
	Status string `json:"status,omitempty"`
	// Contract name
	Contract string `json:"contract"`
	// Order size. Specify positive number to make a bid, and negative number to ask
	Size int64 `json:"size"`
	// Display size for iceberg order. 0 for non-iceberg. Note that you will have to pay the taker fee for the hidden size
	Iceberg int64 `json:"iceberg,omitempty"`
	// Order price. 0 for market order with `tif` set as `ioc` (USDT)
	Price string `json:"price,omitempty"`
	// Set as `true` to close the position, with `size` set to 0
	Close bool `json:"close,omitempty"`
	// Is the order to close position
	IsClose bool `json:"is_close,omitempty"`
	// Set as `true` to be reduce-only order
	ReduceOnly bool `json:"reduce_only,omitempty"`
	// Is the order reduce-only
	IsReduceOnly bool `json:"is_reduce_only,omitempty"`
	// Is the order for liquidation
	IsLiq bool `json:"is_liq,omitempty"`
	// 设置为 true 的时候，为MMP委托
	Mmp bool `json:"mmp,omitempty"`
	// 是否为MMP委托。对应请求中的`mmp`。
	IsMmp bool `json:"is_mmp,omitempty"`
	// Time in force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only - poc: PendingOrCancelled, makes a post-only order that always enjoys a maker fee
	Tif string `json:"tif,omitempty"`
	// Size left to be traded
	Left int64 `json:"left,omitempty"`
	// Fill price of the order
	FillPrice string `json:"fill_price,omitempty"`
	// User defined information. If not empty, must follow the rules below:  1. prefixed with `t-` 2. no longer than 28 bytes without `t-` prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.) Besides user defined information, reserved contents are listed below, denoting how the order is created:  - web: from web - api: from API - app: from mobile phones - auto_deleveraging: from ADL - liquidation: from liquidation - insurance: from insurance
	Text string `json:"text,omitempty"`
	// Taker fee
	Tkfr string `json:"tkfr,omitempty"`
	// Maker fee
	Mkfr string `json:"mkfr,omitempty"`
	// Reference user ID
	Refu int32 `json:"refu,omitempty"`
	// Referrer rebate
	Refr string `json:"refr,omitempty"`
}
