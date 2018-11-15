package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"fmt"
)

type SampleChaincode struct {}

//
// https://stackoverflow.com/questions/18559830/function-for-copying-arrays-in-go-language
func (sc *SampleChaincode) invokeOtherChaincode(stub shim.ChaincodeStubInterface, args []string ) peer.Response{

	chaincodeName := args[0]
	channelName := args[1]
	var byteArray [][]byte
	for _, arg := range args[2:] {
		byteArray = append(byteArray, []byte(arg))
	}


	return stub.InvokeChaincode(chaincodeName , byteArray, channelName  )

}


func (sc *SampleChaincode) Init( stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)

}

func (sc *SampleChaincode) Invoke( stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "invokeOtherChaincodeSetAsset" {
		return sc.invokeOtherChaincode(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")

}


func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SampleChaincode))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}

}