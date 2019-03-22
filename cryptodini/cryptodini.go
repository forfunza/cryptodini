package cryptodini

import "bot/binance"

type CryptodiniService interface {
	Adjust(uid binance.UserID, desiredPort binance.Portfolio) string
	GetPort(uid binance.UserID) *binance.Portfolio
}

func Adjust(uid binance.UserID, desiredPort binance.Portfolio) string {
	//Algorithm for Adjust and just call Buy Or Sell
	return "Port has been Adjust"
}

func GetPort(uid binance.UserID) binance.Portfolio {
	return binance.GetPort(uid.Secret)
}
