package main

import (
	"bot/binance"
	"bot/cryptodini"
	"math/rand"

	"github.com/gin-gonic/gin"
)

type AssetManagerService interface {
	Deposit(uid binance.UserID, amount float64) binance.BuyOrders
	Withdraw(uid binance.UserID, amount float64) binance.SellOrders
	GetPort(uid binance.UserID) binance.Portfolio
}

func main() {
	r := gin.Default()
	user := binance.UserID{ID: 1, Secret: "FFJAAFJHJAHUFAVBVVMEHYEBVEJEJBEJEJBVVA"}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	trade := r.Group("/api/trade")
	trade.GET("/buy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": Deposit(user, randFloats(10000, 200000000)),
		})
	})

	trade.GET("/sell", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": Withdraw(user, randFloats(10000, 200000000)),
		})
	})

	trade.GET("/adjust", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": cryptodini.Adjust(user, binance.GetPort(user.Secret)),
		})
	})

	trade.GET("/portfolio", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": GetPort(user),
		})
	})

	r.Run(":9091") // listen and serve on 0.0.0.0:8080

}

//Deposit Function
func Deposit(uid binance.UserID, amount float64) binance.BuyOrders {
	b := binance.BuyOrders{Symbol: "ETHTUSD", Price: randFloats(1, 2), Amount: randFloats(10, 2000000)}
	binance.BuyOrder(uid.Secret, b)
	return b
}

//Deposit Function
func Withdraw(uid binance.UserID, amount float64) binance.SellOrders {
	s := binance.SellOrders{Symbol: "OMGTUSD", Price: randFloats(1, 2), Amount: randFloats(10, 2000000)}
	binance.SellOrder(uid.Secret, s)
	return s
}

func GetPort(uid binance.UserID) binance.Portfolio {
	p := binance.GetPort(uid.Secret)
	return p
}

func randFloats(min, max float64) float64 {
	res := min + rand.Float64()*(max-min)
	return res
}
