/*
Copyright IBM Corp. 2017,2018 All Rights Reserved.

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

var logger = shim.NewLogger("ContractChaincodeLog")

// ContractChaincode implementation
type ContractChaincode struct {
}

// Init nothing to initialize
func (t *ContractChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### Contract Init ###########")
	//nothing to initialize just return
	return shim.Success(nil)
}

// Invoke Support for calling chaincode to ensure operation is up
func (t *ContractChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### Contract Invoke ###########")

	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "Health":
		// contract chaincode
		return t.Health(stub, args)
	}

	errorMsg := fmt.Sprintf("Unknown action, please check the first argument, expecting 'Health'. Instead, got: %s", function)
	logger.Errorf(errorMsg)
	return shim.Error(errorMsg)
}

// Health returns Ok if successful
func (t *ContractChaincode) Health(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Infof("Chaincode is healthy.")
	return shim.Success([]byte("Ok"))
}

func main() {
	err := shim.Start(new(ContractChaincode))
	if err != nil {
		logger.Errorf("Error starting ContractChaincode: %s", err)
	}
}
