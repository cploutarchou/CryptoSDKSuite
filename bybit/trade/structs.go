package trade

type PlaceOrderRequest struct {
	Category         string  `json:"category"`
	Symbol           string  `json:"symbol"`
	IsLeverage       *int    `json:"isLeverage,omitempty"`
	Side             string  `json:"side"`
	OrderType        string  `json:"orderType"`
	Qty              string  `json:"qty"`
	Price            *string `json:"price,omitempty"`
	TriggerPrice     *string `json:"triggerPrice,omitempty"`
	TriggerDirection *int    `json:"triggerDirection,omitempty"`
	TriggerBy        *string `json:"triggerBy,omitempty"`
	OrderFilter      *string `json:"orderFilter,omitempty"`
	OrderIv          *string `json:"orderIv,omitempty"`
	TimeInForce      string  `json:"timeInForce"`
	PositionIdx      *int    `json:"positionIdx,omitempty"`
	OrderLinkId      string  `json:"orderLinkId"`
	TakeProfit       *string `json:"takeProfit,omitempty"`
	StopLoss         *string `json:"stopLoss,omitempty"`
	TpTriggerBy      *string `json:"tpTriggerBy,omitempty"`
	SlTriggerBy      *string `json:"slTriggerBy,omitempty"`
	ReduceOnly       *bool   `json:"reduceOnly,omitempty"`
	CloseOnTrigger   *bool   `json:"closeOnTrigger,omitempty"`
	SmpType          *string `json:"smpType,omitempty"`
	Mmp              *bool   `json:"mmp,omitempty"`
	TpslMode         *string `json:"tpslMode,omitempty"`
	TpLimitPrice     *string `json:"tpLimitPrice,omitempty"`
	SlLimitPrice     *string `json:"slLimitPrice,omitempty"`
	TpOrderType      *string `json:"tpOrderType,omitempty"`
	SlOrderType      *string `json:"slOrderType,omitempty"`
}

type PlaceOrderResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderID     string `json:"orderId"`
		OrderLinkID string `json:"orderLinkId"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}
