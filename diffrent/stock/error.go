package stock

type ErrorStock string

var (
	SuchProductExist    ErrorStock = "such product exist"
	SuchProductNotExist ErrorStock = "such product not exist"
	NotEnoughProduct    ErrorStock = "not enough product"
)

func (e ErrorStock) Error() string {
	return string(e)
}
