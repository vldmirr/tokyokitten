package request

type UpdateKittensRequest struct {
	Id      int    `validate:"required"`
	Name    string `validate:"required,max=200,min=1" json:"name"`
	// Age     int    `validate:"required,min=1,max=200" json:"age"`
	// Color   string `validate:"required,min=1,max=200" json:"color"`
	// Breed   string `validate:"required,min=1,max=200" json:"breed"`
	// Price   string `validate:"required,min=1,max=200" json:"price"`
	// InStock string `validate:"required,min=1,max=200" json:"instock"`
}
