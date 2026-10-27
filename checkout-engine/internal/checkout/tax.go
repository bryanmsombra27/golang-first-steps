package checkout

type TaxFn func(Order) Money

func NoTax(Order) Money {
	return 0
}

func IVA16(order Order) Money {
	sub := CalcSubTotal(order)

	return sub * 16 / 100
}
func NewTaxByState(state string) TaxFn {

	switch state {
	case "CDMX":
		return func(order Order) Money {
			return CalcSubTotal(order) * 16 / 100
		}
	case "NL":
		return func(order Order) Money {
			return CalcSubTotal(order) * 15 / 100
		}
	case "QRO":
		return func(order Order) Money {
			return CalcSubTotal(order) * 20 / 100
		}
	case "GDL":
		return func(order Order) Money {
			return CalcSubTotal(order) * 14 / 100
		}
	default:
		return NoTax
	}

}

func NewShippingByZone(zone string) ShippingFn {
	switch zone {
	case "LOCAL":
		return FreeShipping
	case "NATIONAL":
		return func(Order) Money {
			return 499
		}
	default:
		return func(Order) Money { return 2499 }
	}

}