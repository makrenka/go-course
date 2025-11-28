package lessons

type User struct {
	id      int
	name    string
	email   string
	phone   string
	address Address
	cart    []Cart
}

type Address struct {
	street     string
	city       string
	postalCode string
}

type Cart struct {
	product  Product
	quantity int
}

type Product struct {
	id          int
	name        string
	description string
	price       int
	category    string
	brand       string
	rating      float64
	reviews     string
}

func PrintInfo(user User) {

}
