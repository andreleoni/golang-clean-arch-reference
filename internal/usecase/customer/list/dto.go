package list

type InputCustomerListDTO struct{}

type OutputCustomerListDTO struct {
	Customers []OutputCustomerDTO `json:"customers"`
}

type OutputCustomerDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
