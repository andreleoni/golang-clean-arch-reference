package valueobject

type Address struct {
	Street     string `pg:"address_street"`
	Number     string `pg:"address_number"`
	Complement string `pg:"address_complement"`
	Zipcode    string `pg:"address_zipcode"`
	City       string `pg:"address_city"`
	Province   string `pg:"address_province"`
}
