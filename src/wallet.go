package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Wallet object
type Wallet struct {
}

// Doc writes string to the blockchain (as JSON object) for a specific key
type Doc struct {
	Amount float64 `json:text`
}

func (doc *Doc) FromJson(input []byte) *Doc {
	json.Unmarshal(input, doc)
	return doc
}

func (doc *Doc) ToJson() []byte {
	jsonDoc, _ := json.Marshal(doc)
	return jsonDoc
}


// Init will do nothing
func (w *Wallet) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func Success(rc int32, doc string, payload []byte) pb.Response {
	return pb.Response{
		Status:  rc,
		Message: doc,
		Payload: payload,
	}
}

// Invoke wallet method
func (w *Wallet) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	switch fn {
	case "deposit":
		return deposit(stub, args)
	case "withdrawal":
		return withdrawal(stub, args)
	case "transfer":
		return transfer(stub, args)
	case "get":
		return get(stub, args)
	case "getAllKeys":
		return getAllKeys(stub, args)
	case "history":
		return history(stub, args)
	default:
		return shim.Error("Unsupported operation")
	}
}
