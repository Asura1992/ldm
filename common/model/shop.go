package model

type Shop struct {
	ShopId        uint64  `gorm:"primary_key" json:"shop_id"`
	ShopShopifyId *int64  `gorm:"default:0" json:"shop_shopify_id"`
	ShopName      *string `gorm:"default:''" json:"shop_name"`
}
