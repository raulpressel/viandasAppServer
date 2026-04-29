package db

import "viandasApp/db"

func CountProductsByCategory(categoryID int) (int64, error) {
	db := db.GetDB()
	var count int64
	err := db.Table("products").
		Where("product_category_id = ? AND deleted_at IS NULL", categoryID).
		Count(&count).Error
	return count, err
}
