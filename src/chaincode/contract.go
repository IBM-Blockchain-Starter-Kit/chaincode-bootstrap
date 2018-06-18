/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("contract health")

// ContractChaincode implementation
type ContractChaincode struct {
}

// Init nothing to initialize
func (t *ContractChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### ping Init ###########")
	//nothing to initialize just return
	return shim.Success(nil)

}

// Invoke Support for calling chaincode to ensure operation is up
func (t *ContractChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### contract Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "Health" {
		// ping chaincode
		return t.Health(stub, args)
	}

	logger.Errorf("Unknown action, check the first argument, must be 'ping'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be 'ping'. But got: %v", args[0]))
}

// Health returns Ok if successful
func (t *ContractChaincode) Health(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Infof("Chaincode pinged successfully")
	return shim.Success([]byte("Ok"))
}

func main() {
	err := shim.Start(new(ContractChaincode))
	if err != nil {
		logger.Errorf("Error starting Contract chaincode: %s", err)
	}
}
