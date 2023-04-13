package valueobject

type Address struct {
	Street     string `gorm:"column:street"`
	Number     string `gorm:"column:number"`
	Complement string `gorm:"column:complement"`
	Zipcode    string `gorm:"column:zipcode"`
	City       string `gorm:"column:city"`
	Province   string `gorm:"column:province"`
}
