package list

type InputCustomerListDTO struct {
	ID string
}

type OutputCustomerListDTO struct {
	Customers []OutputCustomerDTO `json:"customers"`
}

type OutputCustomerDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
