package main

import (
	"net/http"
	"strconv"

	"transactions"

	"gopkg.in/labstack/echo.v4"
)

const NETWORK = "https://rinkeby.infura.io/v3/083fe5705feb4e21aeae65d411396ed9"
const SENDER_PRIVATE_KEY = "0be685f4707d7d22bc971f78269d39fee695c6c764a8c2ae83158d62df660a31"
const TO_ADDRESS = "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d"
const DATA = "Banana,Orange,Apple,Mango"

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/send/:data", sendTo)
	e.GET("/verify/:txHash/:data", verify)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func sendTo(c echo.Context) error {
	data := c.Param("data")
	txBuilder := transactions.TransactionBuilder(NETWORK, SENDER_PRIVATE_KEY)
	txHash, _ := txBuilder.SendTo(TO_ADDRESS, data)
	return c.String(http.StatusOK, "tx: "+txHash)
}
func verify(c echo.Context) error {
	txHash := c.Param("txHash")
	dataExpected := c.Param("data")
	txBuilder := transactions.TransactionBuilder(NETWORK, SENDER_PRIVATE_KEY)
	tx := txBuilder.GetTx(txHash)
	dataResult := string(tx.Data())
	return c.String(http.StatusOK, strconv.FormatBool(dataResult == dataExpected))
}
