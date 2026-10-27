package checkout

type Money = int //centavos
type Item struct {
	SKU      string
	Name     string
	Price    Money
	Quantity int
}
type Order struct {
	ID       string
	Customer string
	Items    []Item
	Meta     map[string]string
}

type Totals struct {
	Subtotal Money
	Discount Money
	Tax      Money
	Shipping Money
	Total    Money
}