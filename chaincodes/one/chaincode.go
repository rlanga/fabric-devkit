package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("chaincode-one")

// SimpleChaincode representing a class of chaincode
type SimpleChaincode struct{}

// Init to initiate the SimpleChaincode class
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("chaincode-one: Hello Init")
	return shim.Success([]byte("Initialisation completed"))
}

// Invoke a method specified in the SimpleChaincode class
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("chaincode-one: Hello Invoke")
	return shim.Success([]byte("Invoke"))
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Debugf("Error: %s", err)
	}
}
