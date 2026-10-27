package checkout

import "fmt"

func RunDemo() {
	PrintHeader("Hola Checkout Engine :)")
	PrintDivider()

	order := newOrder("ORDER-0001", "Bryan")
	order.AddItem(Item{SKU: "KB001", Name: "Teclado", Price: 3500, Quantity: 1})
	order.AddItem(Item{SKU: "MB024", Name: "Monitor", Price: 15000, Quantity: 2})
	order.AddItem(Item{SKU: "MB054", Name: "CPU", Price: 45000, Quantity: 3})

	// VALIDANDO ITEMS
	PrintKV("VALIDADOR", ValidateOrder(order))

	PrintKV("OrderID", order.ID)
	PrintKV("Customer", order.Customer)
	PrintKV("Items", len(order.Items))

	remove := order.RemoveItem("KB001")
	PrintKV("Removed KB001: ", remove)

	PrintDivider()

	sub := CalcSubTotal(order)
	qty := order.CalcTotalQty()

	PrintKV("Subtotal: ", sub)
	PrintKV("Cantidad: ", qty)
	PrintDivider()

	TryChangeCustomerByValue(order, "Nuevo nombre")
	PrintKV("Customer no cambia: ", order.Customer)

	ChangeCustomerByPointer(&order, "Bryan Ochoa")
	PrintKV("Customer cambia: ", order.Customer)

	setCity(order, "CDMX")
	PrintKV("Ciudad(map si cambia): ", order.Meta["city"])
	PrintDivider()

	setZone(order, "NATIONAL")
	PrintKV("Zone(map si cambia): ", order.Meta["zone"])
	PrintDivider()

	items := []Item{
		{
			SKU:      "ms003",
			Name:     "mouse",
			Price:    1200,
			Quantity: 1,
		},
		{
			SKU:      "hd005",
			Name:     "HDMI",
			Price:    300,
			Quantity: 2,
		},
	}

	order.AddItems(items...)
	PrintKV("Cantidad Total: ", order.CalcTotalQty())
	PrintKV("Items: ", order.Items)

	PrintDivider()

	findItem, extraValueFound := order.FindItem("ms003")
	PrintKV2("Item encontrado: ", findItem, extraValueFound)

	getMeta, extraGetMeta := GetMeta(order, "city")
	PrintKV2("Meta Dato: ", getMeta, extraGetMeta)
	index, wasFound := IndexOfItem(order, "ms003")
	PrintKV2("Index encontrado: ", index, wasFound)
	PrintDivider()

	couponValue, errorCupon := ParseCoupon("SAVE10")
	PrintKV2("Cupon: ", couponValue, errorCupon)

	PrintDivider()

	// computeValue, totalError := Compute(order,)
	// PrintKV2("Cupon: ", computeValue, totalError)
	PrintDivider()

	PrintKV("Descuento: ", FlatDiscount(200))
	th := ThresholdPercentDiscount(2000, 20)
	PrintKV("Descuento %: ", th(order))
	PrintDivider()

	cityDiscount := func(order Order) Money {
		city, _ := GetMeta(order, "city")

		if city == "Tijuana" {
			return 200
		}
		return 0
	}
	PrintKV("Descuento especial por ciudad", cityDiscount(order))
	PrintDivider()
	
	discount := MakeSkuDiscount("TechProduct", 500)
	discountHDMI := MakeSkuDiscount("HD005", 150)
	
	fmt.Println(discount(order))
	fmt.Println(discountHDMI(order))
	PrintDivider()

	state,_ := GetMeta(order,"city")
	zone,_ := GetMeta(order,"zone")

	taxfn := NewTaxByState(state)
	shipfn := NewShippingByZone(zone) 

	promo:= CompositeDiscount{
		Name: "Promicion Febrero",
		Fns: []DiscountFn{
			FlatDiscount(100),
			ThresholdPercentDiscount(2000,20),
			MakeSkuDiscount("HD005",200),
		},
	}


	bundle:=	promo.Apply(order)
	PrintKV("Descuento Recursivo",StringUSD(bundle))
	PrintDivider()

	computeValue, totalError := Compute(order, bundle,taxfn,shipfn,FlatDiscount(5000),ThresholdPercentDiscount(2000,10))
	PrintKV2("computar valores por nombre (TOTALES): ", computeValue, totalError)	
	
}
