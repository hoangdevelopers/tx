package main

import (
	"transactions"
)

const NETWORK = "https://rinkeby.infura.io/v3/083fe5705feb4e21aeae65d411396ed9"
const SENDER_PRIVATE_KEY = "0be685f4707d7d22bc971f78269d39fee695c6c764a8c2ae83158d62df660a31"
const TO_ADDRESS = "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d"
const DATA = "Banana,Orange,Apple,Mango"

func main() {
	txBuilder := transactions.TransactionBuilder(NETWORK, SENDER_PRIVATE_KEY)
	txBuilder.SendTo(TO_ADDRESS, DATA)
}
