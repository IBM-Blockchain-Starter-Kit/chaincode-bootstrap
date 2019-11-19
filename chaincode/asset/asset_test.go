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
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// TxID is just a dummy transactional ID for test cases
const TxID = "mockTxID"

// MockStubUUID - Dummy UUID value used accross all invocations to MockInit() method
const MockStubUUID = "d2490ad8-3901-11e8-b467-0ed5f89f718a"

func TestContractChaincode_Ping(t *testing.T) {
	scc := new(ContractChaincode)
	stub := shim.NewMockStub("ContractChaincode", scc)
	// Contract
	checkPing(t, stub)
}

func checkPing(t *testing.T, stub *shim.MockStub) {
	res := stub.MockInvoke(MockStubUUID, [][]byte{[]byte("ping")})
	checkResponse(t, res, "ping")

	result := string(res.Payload)
	if string(res.Payload) != "Ok" {
		fmt.Println("Contract return value not expected. Value expected is Ok but got ", result)
		t.FailNow()
	}
}

func TestContractChaincode_MyAssetExists(t *testing.T) {
	scc := new(ContractChaincode)
	stub := shim.NewMockStub("ContractChaincode", scc)

	// Store test data
	key := "key1"
	value := "value1"
	storeAsset(t, stub, key, value)

	// Test method
	checkMyAssetExists(t, stub, key)
	checkMyAssetDoesNotExist(t, stub, "invalidKey")
}

func checkMyAssetExists(t *testing.T, stub *shim.MockStub, key string) {
	mockInvokeArgs := [][]byte{[]byte("myAssetExists"), []byte(key)}
	res := stub.MockInvoke(MockStubUUID, mockInvokeArgs)
	checkResponse(t, res, "myAssetExists")

	result, _ := strconv.ParseBool(string(res.Payload))
	if !result {
		errMsg := fmt.Sprintf("Contract return unexpected value. Expected 'true' but got '%t'.", result)
		fmt.Println(errMsg)
		t.FailNow()
	}
}

func checkMyAssetDoesNotExist(t *testing.T, stub *shim.MockStub, key string) {
	mockInvokeArgs := [][]byte{[]byte("myAssetExists"), []byte(key)}
	res := stub.MockInvoke(MockStubUUID, mockInvokeArgs)
	checkResponse(t, res, "myAssetExists")

	result, _ := strconv.ParseBool(string(res.Payload))
	if result {
		errMsg := fmt.Sprintf("Contract return unexpected value. Expected 'false' but got '%t'.", result)
		fmt.Println(errMsg)
		t.FailNow()
	}
}

func TestContractChaincode_UpdateMyAsset(t *testing.T) {
	scc := new(ContractChaincode)
	stub := shim.NewMockStub("ContractChaincode", scc)

	key := "key1"
	value := "value1"
	newValue := "value2"

	storeAsset(t, stub, key, value)
	checkUpdateMyAsset(t, stub, key, newValue)
	var updatedAsset = readAsset(t, stub, key)

	if updatedAsset.Value != newValue {
		fmt.Println("Failed to update asset.")
		t.FailNow()
	}
}

func checkUpdateMyAsset(t *testing.T, stub *shim.MockStub, key string, newValue string) {
	mockInvokeArgs := [][]byte{[]byte("updateMyAsset"), []byte(key), []byte(newValue)}
	res := stub.MockInvoke(MockStubUUID, mockInvokeArgs)
	checkResponse(t, res, "updateMyAsset")

	asset := unmarshalAsset(t, res.Payload)
	if asset.Value != newValue {
		fmt.Println("Failed to update asset")
		t.FailNow()
	}
}

func TestContractChaincode_DeleteMyAsset(t *testing.T) {
	scc := new(ContractChaincode)
	stub := shim.NewMockStub("ContractChaincode", scc)
	key := "key1"
	value := "value1"
	storeAsset(t, stub, key, value)
	checkDeleteMyAsset(t, stub, key)
}

func checkDeleteMyAsset(t *testing.T, stub *shim.MockStub, key string) {
	mockInvokeArgs := [][]byte{[]byte("deleteMyAsset"), []byte(key)}
	res := stub.MockInvoke(MockStubUUID, mockInvokeArgs)
	checkResponse(t, res, "deleteMyAsset")

	retKey := string(res.Payload)
	if key != retKey {
		errMsg := fmt.Sprintf("Unexpected key returned: '%s'.", retKey)
		fmt.Println(errMsg)
		t.FailNow()
	}
}

func TestContractChaincode_GetMyAsset(t *testing.T) {
	scc := new(ContractChaincode)
	stub := shim.NewMockStub("ContractChaincode", scc)
	key := "key1"
	value := "value1"
	storeAsset(t, stub, key, value)
	checkGetMyAsset(t, stub, key, value)
}

func TestContractChaincode_CreateMyAsset(t *testing.T) {
	scc := new(ContractChaincode)
	stub := shim.NewMockStub("ContractChaincode", scc)
	key := "key1"
	value := "value1"
	checkCreateMyAsset(t, stub, key, value)

	asset := readAsset(t, stub, key)
	if asset.Value != value {
		fmt.Println("Failed to create asset.")
		t.FailNow()
	}
}

func checkCreateMyAsset(t *testing.T, stub *shim.MockStub, key string, value string) {
	mockInvokeArgs := [][]byte{[]byte("createMyAsset"), []byte(key), []byte(value)}
	res := stub.MockInvoke(MockStubUUID, mockInvokeArgs)
	checkResponse(t, res, "createMyAsset")

	asset := unmarshalAsset(t, res.Payload)
	if asset.Value != value {
		errMsg := fmt.Sprintf("Unexpected asset value '%s'.", asset.Value)
		fmt.Println(errMsg)
		t.FailNow()
	}
}

func checkGetMyAsset(t *testing.T, stub *shim.MockStub, key string, value string) {
	mockInvokeArgs := [][]byte{[]byte("getMyAsset"), []byte(key)}
	res := stub.MockInvoke(MockStubUUID, mockInvokeArgs)
	checkResponse(t, res, "getMyAsset")

	asset := unmarshalAsset(t, res.Payload)
	if asset.Value != value {
		errMsg := fmt.Sprintf("Unexpected asset value '%s'.", asset.Value)
		fmt.Println(errMsg)
		t.FailNow()
	}
}

func storeAsset(t *testing.T, stub *shim.MockStub, key string, value string) MyAsset {
	// Store test asset
	asset := MyAsset{Value: value}
	assetAsBytes, err := json.Marshal(asset)
	if err != nil {
		fmt.Println("Failed to marshal asset. Error: " + err.Error())
		t.FailNow()
	}

	// Need a dummy transaction before we can call the stub.PutState() method
	stub.MockTransactionStart(TxID)
	stub.PutState(key, assetAsBytes)
	stub.MockTransactionEnd(TxID)
	return asset
}

func readAsset(t *testing.T, stub *shim.MockStub, key string) MyAsset {
	stub.MockTransactionStart(TxID)
	assetAsBytes, err := stub.GetState(key)
	stub.MockTransactionEnd(TxID)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to read asset with key '%s'. Error: %s", key, err.Error())
		fmt.Println(errMsg)
		t.FailNow()
	}

	if assetAsBytes == nil {
		errMsg := fmt.Sprintf("Could not find asset with key '%s'.", key)
		fmt.Println(errMsg)
		t.FailNow()
	}

	asset := unmarshalAsset(t, assetAsBytes)
	return asset
}

func checkResponse(t *testing.T, res pb.Response, method string) {
	if res.Status != shim.OK {
		errMsg := fmt.Sprintf("%s failed. %s", method, string(res.Message))
		fmt.Println(errMsg)
		t.FailNow()
	}

	if res.Payload == nil {
		errMsg := fmt.Sprintf("%s failed to return a value.", method)
		fmt.Println(errMsg)
		t.FailNow()
	}
}

func unmarshalAsset(t *testing.T, assetAsBytes []byte) MyAsset {
	var asset MyAsset
	err := json.Unmarshal(assetAsBytes, &asset)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to unmarshal asset: %s", err.Error())
		fmt.Println(errMsg)
		t.FailNow()
	}
	return asset
}
