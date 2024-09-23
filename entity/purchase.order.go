package entity

import "gopkg.in/guregu/null.v4"

// PurchaseOrderStatistic 采购单统计
type PurchaseOrderStatistic struct {
	StatusNumMap       map[string]int `json:"statusNumMap"`       // key:采购单状态 value:该状态下数量。 订单状态 0-待接单；1-已接单，待发货；2-已送货；3-已收货；5-已验收，全部退回；6-已验收；7-已入库；8-作废；9-已超时； 10-已取消
	Total              int            `json:"total"`              // 总数
	TodayCanDeliverNum int            `json:"todayCanDeliverNum"` // 今日可发货数量
	DelayNumMap        map[string]int `json:"delayNumMap"`        // key:逾期状态 101-发货即将逾期，102-发货已逾期，201-到货即将逾期，202-到货已逾期 value:备货单数
}

type PurchaseOrderSkuQuantityDetailList struct {
	CurrencyType                 string   `json:"currencyType"`
	ClassName                    string   `json:"className"`
	RealReceiveAuthenticQuantity int      `json:"realReceiveAuthenticQuantity"`
	FulfilmentProductSkuId       int      `json:"fulfilmentProductSkuId"`
	CustomizationType            int      `json:"customizationType"`
	ProductSkuId                 int      `json:"productSkuId"`
	DeliverQuantity              int      `json:"deliverQuantity"`
	ThumbUrlList                 []string `json:"thumbUrlList"`
	QcResult                     any      `json:"qcResult"`
	ExtCode                      string   `json:"extCode"`
	PurchaseQuantity             int      `json:"purchaseQuantity"`
}

// PurchaseOrder 采购单
type PurchaseOrder struct {
	OriginalPurchaseOrderSn string                               `json:"originalPurchaseOrderSn"` // 母订单号（原始采购母单号）
	SubPurchaseOrderSn      string                               `json:"subPurchaseOrderSn"`      // 采购子单号
	Source                  int                                  `json:"source"`
	ProductName             string                               `json:"productName"`
	FulfilmentFormStatus    any                                  `json:"fulfilmentFormStatus"`
	IsFirst                 bool                                 `json:"isFirst"`
	SkuQuantityDetailList   []PurchaseOrderSkuQuantityDetailList `json:"skuQuantityDetailList"`
	DeliverInfo             struct {
		ReceiveTime                      null.Int    `json:"receiveTime"`
		DeliverTime                      null.Int    `json:"deliverTime"`
		ReceiveWarehouseId               null.Int    `json:"receiveWarehouseId"`
		ReceiveWarehouseName             null.String `json:"receiveWarehouseName"`
		ExpectLatestDeliverTimeOrDefault int         `json:"expectLatestDeliverTimeOrDefault"`
		ExpectLatestArrivalTimeOrDefault int         `json:"expectLatestArrivalTimeOrDefault"`
		DeliveryOrderSn                  null.String `json:"deliveryOrderSn"`
	} `json:"deliverInfo"`
	ProductSkcId         int `json:"productSkcId"`
	ProductId            int `json:"productId"`
	HasQcBill            int `json:"hasQcBill"`
	SupplyStatus         int `json:"supplyStatus"`
	ApplyDeleteStatus    int `json:"applyDeleteStatus"`
	SkuQuantityTotalInfo struct {
		CurrencyType                 any `json:"currencyType"`
		ClassName                    any `json:"className"`
		RealReceiveAuthenticQuantity int `json:"realReceiveAuthenticQuantity"`
		CustomizationType            any `json:"customizationType"`
		ProductSkuId                 any `json:"productSkuId"`
		DeliverQuantity              int `json:"deliverQuantity"`
		ExtCode                      any `json:"extCode"`
		PurchaseQuantity             int `json:"purchaseQuantity"`
	} `json:"skuQuantityTotalInfo"`
	IsCanJoinDeliverPlatform bool     `json:"isCanJoinDeliverPlatform"`
	CategoryType             int      `json:"categoryType"`
	Status                   int      `json:"status"`
	FulfilmentFormId         null.Int `json:"fulfilmentFormId"` // 关联履约函 ID
	ProductSkcPicture        string   `json:"productSkcPicture"`
	LackOrSoldOutTagList     []struct {
		IsLack     bool   `json:"isLack"`
		SkuDisplay string `json:"skuDisplay"`
		SoldOut    any    `json:"soldOut"`
	} `json:"lackOrSoldOutTagList"`
	QcReject          int `json:"qcReject"`
	PurchaseStockType int `json:"purchaseStockType"`
	SkuLackItemList   []struct {
		SkuDisplay string `json:"skuDisplay"`
	} `json:"skuLackItemList"`
	SkuLackSnapshot         int      `json:"skuLackSnapshot"`
	DeliveryOrderSn         *string  `json:"deliveryOrderSn"`
	SettlementType          int      `json:"settlementType"`
	UrgencyType             int      `json:"urgencyType"`
	ProductSn               *string  `json:"productSn"`
	SkuQcRejectItemList     any      `json:"skuQcRejectItemList"`
	DefectiveTime           null.Int `json:"defectiveTime"`
	TodayCanDeliver         bool     `json:"todayCanDeliver"`
	PurchaseTime            int      `json:"purchaseTime"`
	ApplyChangeSupplyStatus int      `json:"applyChangeSupplyStatus"`
	Category                string   `json:"category"`
}
