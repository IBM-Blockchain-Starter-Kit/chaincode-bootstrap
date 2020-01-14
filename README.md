[![Build Status - Master](https://travis-ci.org/IBM-Blockchain-Starter-Kit/chaincode-bootstrap.svg?branch=master)](https://travis-ci.org/IBM-Blockchain-Starter-Kit/chaincode-bootstrap/builds)

# GO chaincode scaffolding for Blockchain Starter Kit

This project contains the scaffolding necessary to create a Hyperledger Fabric based smart contract for Golang.  The source code in the the project has been structured to integrate with the DevOps pipeline scripts created for our custom Blockchain Automation Toolchain based on the IBM Cloud Continuos Delivery Service. You can use this scaffolding as a starting point for your smart contract and connect it to the Blockchain Automation Toolchain to have it built, tested and deployed in a CI/CD fashion on the IBM Cloud.

## Chaincode Development Instructions

This project is structured to support the development of multiple smart contracts that can be deployed as separate chaincode units into a Hyperledger Fabric network.

A sample smart contract can be found under the [asset](/chaincode/asset) directory. This smart contract provides a starting point with the necessary transactions to create and update a simple asset.  You can modify this smart contract with your own business logic, attributes and transactions.

Additional smart contracts can be developed and added under the [chaincode](/chaincode) directory. To create a new smart contract, create a new directory under [chaincode](/chaincode) and place the source code in the new directory i.e. `chaincode/new-asset`.

Notice that the purpose of this structure is to allow for the definition of multiple independent smart contracts that can easily be deployed separately by the Blockchain Automation Toolchain hence each smart contract should be self contained and self deployable.

## GO vendoring

If your GO chanincode requires vendoring of GO packages, you should include a `.govendor_packages` file inside each chaincode component folder. This is a custom mechanism that we implemented in the scaffolding for the go based chaincode to make it easier for developers to specify dependencies that have to be pulled into the chaincode at deploy time.  

At the moment, you should only specify in the `.govendor_packages` file libraries that are part of the Fabric binaries. Specifying libraries that are not in the Fabric binaries will more than likely result in compilation errors during the build phase of the toolchain (we plan to address this limitation soon).

The `.govendor_packages` file should contain one line for every GO package that should be vendored in. The syntax for specifying a GO dependency in the `.govendor_packages` file simply follows the format required by the [govendor](https://github.com/kardianos/govendor) tool.

```
github.com/hyperledger/fabric/core/chaincode/lib/cid@v1.2.1
```

The handling of the `.govendor_packages` has been implemented in the [build-lib](IBM-Blockchain-Starter-Kit/build-lib) project of the IBM-Blockchain-Starter-Kit.  

## Chaincode Deployment Instructions

When integrated with the Blockchain Automation Toolchain, the metadata specified in the [deployment configuration](deploy_config.json) file is used to indicate the smart contracts that must be deployed, the Hyperledger Fabric organizations where the smart contracts should be applied to and the channels where each smart contract should be instantiated.  

The sample [deployment configuration](deploy_config.json) file provided in this project is as follows: 

```
{
  "Org1MSP": {
    "chaincode": [
      {
        "name": "mycontract",
        "path": "chaincode/asset",
        "channels": [
          "mychannel"
        ],
        "instantiate": true,
        "install": true
      }
    ]
  }
}
```

The simple structure of this file allows you to specify the following information:

* Organization in which the smart contracts will be deployed. The organization must be represented by the MSP ID used by the organization in the Hyperledger Fabric network.  In this sample we are installing smart contracts on peers for organization known as `Org1MSP` in the network.  Notice that the Blockchain Automation Toolchain supports deploying smart contracts on multiple organizations.  To install smart contracts in additional organizations, add a new element to the [deployment configuration](deploy_config.json) file e.g.
    ```
    {
        "Org1MSP": {
            ...
        },
        "Org2MSP": {
            ...
        }
    }
    ```

* For every smart contract to be deployed, an element describing the details of the smart contract should be added to the chaincode array.  Notice that each element added to the array will be deployed as a separate smart contract in the Hyperledger Fabric network.  For each smart contract created under the [chaincode](/chaincode) directory you should add at least one entry in the array.
* The elements that describe a smart contract are as follows:
  * `name`: This is the name that will be used to deploy the smart contract. 
  * `path`: Path of where the smart contract source can be found. This path is relative to the [deployment configuration](deploy_config.json) file.
  * `channels`: Comma separated string array containing the names of the channels where the smart contracts should be instantiated.
  * `instantiate`: Boolean indicating if the smart contract should be instantiated after installation.
  * `install`: Boolean indicating if the smart contract should be installed or not.

## Environment

We have successfully tested and deployed this scaffolding chaincode component on [Hyperledger Fabric v1.4.4](https://hyperledger-fabric.readthedocs.io/en/release-1.4/releases.html), which requires [Go v1.12](https://golang.org/dl/).