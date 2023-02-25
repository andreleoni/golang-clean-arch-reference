package list

type InputProductListDTO struct{}

type OutputProductListDTO struct {
	Products []OutputProductDTO `json:"products"`
}

type OutputProductDTO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Price  uint64 `json:"price"`
}
