package model

import (
	"time"
)

type Shop struct {
	ShopId                uint64     `gorm:"primary_key" json:"shop_id"`
	ShopShopifyId         *int64     `gorm:"default:0" json:"shop_shopify_id"`
	ShopName              *string    `gorm:"default:''" json:"shop_name"`
	ShopOwner             *string    `gorm:"default:''" json:"shop_owner"`
	ShopEmail             *string    `gorm:"default:''" json:"shop_email"`
	ShopPhone             *string    `gorm:"default:''" json:"shop_phone"`
	ShopCustomerEmail     *string    `gorm:"default:''" json:"shop_customer_email"`
	ShopAddress1          *string    `gorm:"default:''" json:"shop_address1"`
	ShopAddress2          *string    `gorm:"default:''" json:"shop_address2"`
	ShopCity              *string    `gorm:"default:''" json:"shop_city"`
	ShopCountry           *string    `gorm:"default:''" json:"shop_country"`
	ShopCountryCode       *string    `gorm:"default:''" json:"shop_country_code"`
	ShopCountryName       *string    `gorm:"default:''" json:"shop_country_name"`
	ShopProvince          *string    `gorm:"default:''" json:"shop_province"`
	ShopProvinceCode      *string    `gorm:"default:''" json:"shop_province_code"`
	ShopPlanName          *string    `gorm:"default:''" json:"shop_plan_name"`
	ShopPlanDisplayName   *string    `gorm:"default:''" json:"shop_plan_display_name"`
	ShopPasswordEnabled   *int32     `gorm:"default:0" json:"shop_password_enabled"`
	ShopDomain            *string    `gorm:"default:''" json:"shop_domain"`
	ShopMyshopifyDomain   *string    `gorm:"default:''" json:"shop_myshopify_domain"`
	ShopTimezone          *string    `gorm:"default:''" json:"shop_timezone"`
	ShopCurrency          *string    `gorm:"default:''" json:"shop_currency"`
	ShopAccessToken       *string    `gorm:"default:''" json:"shop_access_token"`
	ShopIsAuthorized      *int32     `gorm:"default:2" json:"shop_is_authorized"`
	ShopAuthorizationTime *time.Time `gorm:""  json:"shop_authorization_time"`
	ShopAvailable         *int32     `gorm:"default:1" json:"shop_available"`
	ShopLoginTimes        *uint32    `gorm:"default:0" json:"shop_login_times"`
	ShopLastLoginTime     *time.Time `gorm:""  json:"shop_last_login_time"`
	ShopPricingPlanId     *uint32    `gorm:"default:0" json:"shop_pricing_plan_id"`
	ShopLoginIp           *string    `gorm:"default:0" json:"shop_login_ip"`
	ShopCreatedAt         *time.Time `gorm:""  json:"shop_created_at"`
	ShopUpdatedAt         *time.Time `gorm:""  json:"shop_updated_at"`
	ShopCreateTime        *time.Time `gorm:""  json:"shop_create_time"`
	ShopUpdateTime        *time.Time `gorm:""  json:"shop_update_time"`
}
