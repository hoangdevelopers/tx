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

type IData struct {
	Data   string `json:"data" form:"data" query:"data"`
	TxHash string `json:"txHash" form:"txHash" query:"txHash"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.POST("/send", sendTo)
	e.POST("/verify", verify)

	// Start server
	e.Logger.Fatal(e.Start(":1223"))
}

// Handler
func sendTo(c echo.Context) (err error) {
	data := new(IData)
	if err = c.Bind(data); err != nil {
		return
	}
	txBuilder := transactions.TransactionBuilder(NETWORK, SENDER_PRIVATE_KEY)
	txHash, _ := txBuilder.SendTo(TO_ADDRESS, data.Data)
	return c.String(http.StatusOK, "tx: "+txHash)
}
func verify(c echo.Context) (err error) {
	data := new(IData)
	if err = c.Bind(data); err != nil {
		return
	}
	txBuilder := transactions.TransactionBuilder(NETWORK, SENDER_PRIVATE_KEY)
	tx := txBuilder.GetTx(data.TxHash)
	dataResult := string(tx.Data())
	return c.String(http.StatusOK, strconv.FormatBool(dataResult == data.Data))
}
