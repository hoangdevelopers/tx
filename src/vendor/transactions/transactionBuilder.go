package transactions

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

type SenToFunc func(string, string) (string, string)
type GetTxFunc func(string) *types.Transaction

type ITransactionBuilder struct {
	SendTo SenToFunc
	GetTx  GetTxFunc
}

func TransactionBuilder(netword string, senderPrivateKey string) ITransactionBuilder {
	client, err := ethclient.Dial(netword)

	if err != nil {
		log.Fatal(err)
	}

	return ITransactionBuilder{
		SendTo: func(toAddressHex string, data string) (string, string) {

			privateKey, err := crypto.HexToECDSA(senderPrivateKey)
			if err != nil {
				log.Fatal(err)
			}

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("error casting public key to ECDSA")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			value := big.NewInt(0000000000000000000) // in wei (1 eth)
			gasLimit := uint64(210000)               // in units
			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			toAddress := common.HexToAddress(toAddressHex)
			var dataByteArr []byte = []byte(data)
			tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, dataByteArr)

			chainID, err := client.NetworkID(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			ts := types.Transactions{signedTx}
			rawTxBytes := ts.GetRlp(0)
			rawTxHex := hex.EncodeToString(rawTxBytes)

			//
			rlp.DecodeBytes(rawTxBytes, &tx)

			err = client.SendTransaction(context.Background(), tx)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("tx sent: %s", tx.Hash().Hex())
			return tx.Hash().Hex(), rawTxHex
		},
		GetTx: func(txHashStr string) *types.Transaction {
			txHash := common.HexToHash(txHashStr)
			tx, _, err := client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
			fmt.Println(tx.Value().String())    // 10000000000000000
			fmt.Println(tx.Gas())               // 105000
			fmt.Println(tx.GasPrice().Uint64()) // 102000000000
			fmt.Println(tx.Nonce())             // 110644
			fmt.Println(tx.Data())              // []
			fmt.Println(tx.To().Hex())
			return tx

		},
	}
}
