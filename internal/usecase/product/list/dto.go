package list

type InputCustomerListDTO struct{}

type OutputProductListDTO struct {
	Products []OutputProductDTO `json:"customers"`
}

type OutputProductDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
