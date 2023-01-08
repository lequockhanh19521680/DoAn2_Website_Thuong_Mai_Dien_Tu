package entities

type ProductSpecification struct {
	ID              uint   `gorm:"primarykey"`
	SpecificationID *uint  `gorm:"column:specification_id"`
	ProductID       uint   `gorm:"column:product_id"`
	Properties      string `gorm:"column:properties"`
}

func (ProductSpecification) TableName() string {
	return "ProductSpecification"
}
