package binance

import "fmt"

// Struct Users
type UserID struct {
	ID     int
	Secret string
}

// BuyOrders
type BuyOrders struct {
	Symbol string
	Price  float64
	Amount float64
}

type SellOrders struct {
	Symbol string
	Price  float64
	Amount float64
}

type Portfolio struct {
	uid     UserID
	Symbol  string
	Balance float64
}

type BinanceManagerService interface {
	BuyOrder(secret string, order BuyOrders) string
	SellOrder(secret string, order SellOrders) string
	GetPort(secret string) Portfolio
}

//Call API To Binance for Place Order
func BuyOrder(secret string, order BuyOrders) string {
	fmt.Println("Place Order Buy!")
	return "Place Order Buy"
}

//Call API To Binance for Place Order
func SellOrder(secret string, order SellOrders) string {
	fmt.Println("Place Order Sell!")
	return "Place Order Sell"
}

//Call API To Binance for GetPort
func GetPort(secret string) Portfolio {
	fmt.Println("Get Port")
	var u UserID
	p := Portfolio{uid: u, Symbol: "BTCTUSD", Balance: 1474930}
	return p
}
