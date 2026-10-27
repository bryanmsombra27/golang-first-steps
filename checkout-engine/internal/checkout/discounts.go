package checkout

// variadic
type Coupon struct {
Code string
Kind string
Value int
}

type DiscountFn func(Order) Money

type CompositeDiscount struct {
	Name string
	Fns  []DiscountFn
}

func ApplyCouponCodes(order *Order, codes...string){
	if order.Meta == nil {
		order.Meta= map[string]string{}
	}

	order.Meta["coupons"] = joinCoupons(codes)
}
func joinCoupons(coupons []string) string{
if len(coupons) == 0{
	return ""
}
// PASAMOS DE UN ["CUPON1","CUPON 2 ", "CUPON 3"] A  "CUPON1,CUPON 2 ,CUPON 3"
out:= coupons[0]

for i := 1; i < len(coupons); i++ {
	out+= "," + coupons[i]
	
}

return out

}

func FlatDiscount(amount Money)DiscountFn {
	return func(order Order) Money{ 
		return amount
	}
}

func ThresholdPercentDiscount(min Money,percent int) DiscountFn {

	return func(order Order) Money{
		sub:= CalcSubTotal(order)
		if sub < min {
			return 0
		}
		return sub * Money(percent)

	}

}

func MakeSkuDiscount(sku string, amount Money) DiscountFn{

	return func(order Order) Money{
		_, ok := order.FindItem(sku)

		if !ok {
			return 0 
		}
		return amount

	}

}

func ApplyDiscountRecursive(order Order,fns []DiscountFn) Money{

	if len(fns) == 0 {
		return 0
	}

	return fns[0](order) + ApplyDiscountRecursive(order,fns[1:])

}

func (composite CompositeDiscount) Apply(order Order) Money{
	return ApplyDiscountRecursive(order,composite.Fns)
}