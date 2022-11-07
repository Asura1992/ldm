package model

import "ldm/common/model"

func (h HelloModel) Find(shopId uint64, shopInfo *model.Shop) error {
	return h.db.Find(&shopInfo, "shop_id = ?", 22).Error
}
