/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/chaincode"
	// "gecgithub01.walmart.com/s0g0com/Hyperledger-DisputeNegationSystem/chaincode-go/chaincode"
)

func main() {
	tranChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating transaction chaincode: %v", err)
	}

	if err := tranChaincode.Start(); err != nil {
		log.Panicf("Error starting transaction chaincode: %v", err)
	}
}
