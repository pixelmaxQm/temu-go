package entity

import "gopkg.in/guregu/null.v4"

// JitVirtualInventory 虚拟库存 Jit
type JitVirtualInventory struct {
	ProductSkuId     int64    `json:"productSkuId"`     // 货品SKUId
	SkuStockQuantity null.Int `json:"skuStockQuantity"` // 货品SKU虚拟库存, 不允许查看时返回null
}