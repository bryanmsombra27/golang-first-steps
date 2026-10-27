package checkout

import (
	"errors"
	"fmt"
	"strings"
)

// paso por valor :pasa una copia del valor original
func TryChangeCustomerByValue(order Order,name string) {
	order.Customer = name

}

// paso por referencia pasa el valor original 
func ChangeCustomerByPointer(order *Order, name string){
	order.Customer = name
}

func setCity(order Order,city string){
	if order.Meta == nil {
		order.Meta = map[string]string{} 
	}
	order.Meta["city"] = city
}
func setZone(order Order,zone string){
	if order.Meta == nil {
		order.Meta = map[string]string{} 
	}
	order.Meta["zone"] = zone
}
// valores que se pasan por referencia
// map,slice, func,pointer, chan


// manejo de errores
func ValidateOrder(order Order) error {
if order.ID == ""{
	return errors.New("El id de la orden es obligatorio")
}
if order.Customer == ""{
	return errors.New("El cliente  es obligatorio")
}
if len(order.Items) == 0{
	return errors.New("La orden debe tener al menos un elemento ")
}

for index, item := range order.Items {
 if	item.SKU == ""{
	return fmt.Errorf("Elemento [%d]: el SKU es obligatorio",index)
 }

 if item.Quantity <=0 {
return 	fmt.Errorf("Item [%s]: Su cantidad debe ser mayor a 0",item.SKU)
 }

 if item.Price < 0{
	return fmt.Errorf("Item [%s]: el precio debe ser mayor a 0",item.SKU)
 }

}


	return nil
}

func ParseCoupon(code string) (Coupon,error) {
	coupon := strings.TrimSpace(strings.ToUpper(code))

	if coupon == "" {
		return Coupon{},errors.New("Ingresa un cupón valido")
	}


	switch coupon {
	case "SAVE10":
		return Coupon{
			Code: coupon,
			Kind: "PERCENT",
			Value: 10,
		},nil
	case "LESS500":
		return Coupon{
			Code: coupon,
			Kind: "FLAT",
			Value: 500,
		},nil
	case "FREESHIP":
		return Coupon{
			Code: coupon,
			Kind: "FREESHIP",
			Value:  0,
		},nil
	 default:
		return Coupon{
		
		},fmt.Errorf("Cupon %q: es Invalido",code)
		
	}

		
		
	

}