package checkout

func newOrder(id string, customer string) Order {
	order := Order{
		ID:       id,
		Customer: customer,
		Items:    []Item{},
		Meta:     map[string]string{},
	}

	return order

}

func (order *Order) AddItem(item Item) {
	order.Items = append(order.Items, item)
}
func (order *Order) RemoveItem(sku string) bool {
	for i := range order.Items {
		if order.Items[i].SKU == sku {
			// eslice a la derecha slice a la izquierda y se elimina el elemento buscado
			order.Items = append(order.Items[:i], order.Items[i+1:]...)
			return true
		}
	}
	return false
}

func CalcLineTotal(item Item) Money {
	total := item.Price * Money(item.Quantity)

	return total
}

func CalcSubTotal(order Order) Money {
	var sum Money

	for _, item := range order.Items {
		sum += CalcLineTotal(item)
	}

	return sum
}
func (order Order) CalcTotalQty() int {
	var quantity int = 0
	for _, item := range order.Items {
		quantity += item.Quantity
	}
	return quantity

}

// variadic parametros variables
func (order *Order) AddItems(items ...Item) {
	// spread operator  aqui los 3 puntos se colocan posterior a la variable para esparcir los valores
	order.Items = append(order.Items, items...)
}
func (order *Order) FindItem(sku string) (Item, bool) {
	for _, item := range order.Items {
		if item.SKU == sku {
			return item, true
		}
	}
	return Item{}, false
}

func GetMeta(order Order, key string) (string, bool) {
	if order.Meta == nil {
		return "", false
	}

	value, ok := order.Meta[key]

	return value, ok

}

func IndexOfItem(order Order, sku string) (int, bool) {

	for index, item := range order.Items {
		if item.SKU == sku {
			return index, true
		}
	}
	return -1, false
}

func ApplyDiscount(order Order, fns ...DiscountFn) Money {
	var discount Money

	for _, fn := range fns {
		discount += fn(order)
	}

	sub := CalcSubTotal(order)
	if discount > sub {
		return sub
	}

	return discount
}

func Compute(order Order, bundle Money, taxFn TaxFn, shipfn ShippingFn, discountFn ...DiscountFn) (totals Totals, errorTotal error) {

	defer Track("Compute")()

	if errorTotal = ValidateOrder(order); errorTotal != nil {
		return Totals{}, errorTotal
	}

	totals.Subtotal = CalcSubTotal(order)

	if bundle > 0 {
		totals.Discount = bundle
	} else {
		totals.Discount = ApplyDiscount(order, discountFn...)

	}

	totals.Tax = taxFn(order)
	totals.Shipping = shipfn(order)
	totals.Total = totals.Subtotal - totals.Discount + totals.Shipping + totals.Tax

	return totals, nil

}