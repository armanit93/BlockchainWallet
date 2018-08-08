package main

import (
	"fmt"
	"net/http"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func get(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Not enough arguments")
	}

	key := args[0]

	value, err := stub.GetState(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("Get balance error: %s", err))
	}

	doc := &Doc{Amount: byteToFloat64(value)}


	return Success(http.StatusOK, "OK", doc.ToJson())
}
