package domain

type Order struct {
	ID        uint `gorm:"primaryKey"`
	ProductID int64
	Quantity  int32
}
