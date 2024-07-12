package request

type CreateKittensRequest struct {
	Name    string  `validate:"required,min=1,max=200" json:"name"`
	Age     int     `validate:"required,min=1,max=200" json:"age"`
	Color   string  `validate:"required,min=1,max=200" json:"color"`
	Breed   string  `validate:"required,min=1,max=200" json:"breed"`
	Price   float64 `validate:"required,min=1,max=200" json:"price"`
	Count   uint    `json:"count"`
	InStock bool    `json:"in_stock"`
}
