// 策略模式

package main

import "fmt"

type Order struct {
	price            float64
	discountStrategy func(*Order) float64
}

func NewOrder(price float64, strategy func(*Order) float64) *Order {
	return &Order{
		price:            price,
		discountStrategy: strategy,
	}
}

func (order *Order) String() string {
	return fmt.Sprintf("Order<price=%.2f, price after discount: %.2f>", order.price, order.price-order.discountStrategy(order))
}

func (order *Order) priceAfterDiscount() float64 {
	discount := order.discountStrategy(order)
	return order.price - discount
}

func tenPersonsDiscount(order *Order) float64 {
	return order.price * 0.10
}

func onSaleDiscount(order *Order) float64 {
	return order.price*0.25 + 20
}

func main() {
	order := NewOrder(100.0, tenPersonsDiscount)
	fmt.Println(order)

	order = NewOrder(1000.0, onSaleDiscount)
	fmt.Println(order)
}
