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
	"encoding/json"
	"fmt"
	"strconv"

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
	// nothing to initialize just return
	return shim.Success(nil)
}

// Invoke Support for calling chaincode to ensure operation is up
func (t *ContractChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### Contract Invoke ###########")

	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "ping":
		return t.Ping(stub, args)
	case "myAssetExists":
		return t.MyAssetExists(stub, args)
	case "createMyAsset":
		return t.CreateMyAsset(stub, args)
	case "getMyAsset":
		return t.GetMyAsset(stub, args)
	case "updateMyAsset":
		return t.UpdateMyAsset(stub, args)
	case "deleteMyAsset":
		return t.DeleteMyAsset(stub, args)
	}

	errorMsg := fmt.Sprintf("Unknown action/function, please check the first argument: %s", function)
	logger.Errorf(errorMsg)
	return shim.Error(errorMsg)
}

// Ping returns Ok if successful
func (t *ContractChaincode) Ping(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Info("########### Ping ###########")
	logger.Info("Ping to chaincode successful.")
	return shim.Success([]byte("Ok"))
}

// MyAssetExists returns true if asset is found
func (t *ContractChaincode) MyAssetExists(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Info("########### MyAssetExists ###########")
	if len(args) < 1 {
		errMsg := "Key parameter missing."
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	key := args[0]
	assetAsBytes, err := stub.GetState(key)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to read asset with key '%s'. Error: %s", key, err.Error())
		return shim.Error(errMsg)
	}

	if assetAsBytes == nil {
		infoMsg := fmt.Sprintf("Asset with key '%s' was not found.", key)
		logger.Info(infoMsg)
		return shim.Success([]byte("false"))
	}

	infoMsg := fmt.Sprintf("Found asset with specified key: %s", key)
	logger.Info(infoMsg)
	return shim.Success([]byte("true"))
}

// CreateMyAsset returns true if asset is found
func (t *ContractChaincode) CreateMyAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Info("########### CreateMyAsset ###########")
	if len(args) < 2 {
		errMsg := "Key and/or value parameters missing."
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	res := t.MyAssetExists(stub, args)
	if res.Status != shim.OK || res.Payload == nil {
		return shim.Error(res.GetMessage())
	}

	key := args[0]
	assetFound, err := strconv.ParseBool(string(res.Payload))

	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse boolean, error = %s", err.Error())
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	if assetFound {
		errMsg := fmt.Sprintf("Asset with key '%s' already exists.", key)
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	value := args[1]
	asset := MyAsset{Value: value}
	assetAsBytes, err := json.Marshal(asset)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to marshal asset with key '%s', error = %s", key, err.Error())
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	err = stub.PutState(key, assetAsBytes)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to store asset with key '%s', error = %s", key, err.Error())
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	infoMsg := fmt.Sprintf("Asset with key '%s' successfully created.", key)
	logger.Info(infoMsg)
	return shim.Success(assetAsBytes)
}

// GetMyAsset returns true if asset is found
func (t *ContractChaincode) GetMyAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Info("########### GetMyAsset ###########")
	if len(args) < 1 {
		errMsg := "Key parameter missing."
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	key := args[0]
	assetAsBytes, err := stub.GetState(key)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to read asset with key '%s'. Error: %s", key, err.Error())
		return shim.Error(errMsg)
	}

	if assetAsBytes == nil {
		errMsg := fmt.Sprintf("Could not find asset with key '%s'", key)
		logger.Info(errMsg)
		return shim.Error(errMsg)
	}

	logger.Info("Asset successfully read.")
	return shim.Success(assetAsBytes)
}

// DeleteMyAsset returns true if asset is found
func (t *ContractChaincode) DeleteMyAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Info("########### DeleteMyAsset ###########")
	if len(args) < 1 {
		errMsg := "Key parameter missing."
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	res := t.MyAssetExists(stub, args)
	if res.Status != shim.OK || res.Payload == nil {
		return shim.Error(res.GetMessage())
	}

	key := args[0]
	assetFound, err := strconv.ParseBool(string(res.Payload))

	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse boolean, error = %s", err.Error())
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	if !assetFound {
		errMsg := fmt.Sprintf("Asset with key '%s' does not exist.", key)
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	err = stub.DelState(key)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to delete asset with key '%s'. Error: %s", key, err.Error())
		return shim.Error(errMsg)
	}

	logger.Info("Asset successfully deleted.")
	return shim.Success([]byte(key))
}

// UpdateMyAsset returns true if asset is found
func (t *ContractChaincode) UpdateMyAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Info("########### UpdateMyAsset ###########")

	if len(args) < 2 {
		errMsg := "Key and/or value parameters missing."
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	res := t.MyAssetExists(stub, args)
	if res.Status != shim.OK || res.Payload == nil {
		return shim.Error(res.GetMessage())
	}

	key := args[0]
	assetFound, err := strconv.ParseBool(string(res.Payload))

	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse boolean, error = %s", err.Error())
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	if !assetFound {
		errMsg := fmt.Sprintf("Asset with key '%s' does not exist.", key)
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	value := args[1]
	asset := MyAsset{Value: value}
	assetAsBytes, err := json.Marshal(asset)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to marshal asset with key '%s', error = %s", key, err.Error())
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	err = stub.PutState(key, assetAsBytes)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to store asset with key '%s', error = %s", key, err.Error())
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	logger.Info("Asset successfully updated.")
	return shim.Success(assetAsBytes)
}

func main() {
	err := shim.Start(new(ContractChaincode))
	if err != nil {
		logger.Errorf("Error starting ContractChaincode: %s", err)
	}
}
