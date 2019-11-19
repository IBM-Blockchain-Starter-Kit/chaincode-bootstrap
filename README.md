[![Build Status - Master](https://travis-ci.org/IBM-Blockchain-Starter-Kit/chaincode-bootstrap.svg?branch=master)](https://travis-ci.org/IBM-Blockchain-Starter-Kit/chaincode-bootstrap/builds)

# Chaincode (GO) scaffolding repository for Blockchain Starter Kit

## Chaincode Development Instructions
* Create a new component directory under the [chaincode](/chaincode) folder for each component.
* Populate this component directory with files to be deployed.
* The [asset](/chaincode/asset) directory is provided as a chaincode component example. You can keep or modify this provided chaincode component as you like.

## Chaincode Deployment Instructions
Populate the [deployment configuration](deploy_config.json) JSON file with information about network targets, including organizations, channels, and peers. Please note that included in this repository, there is a deployment configuration example file with the default network architecture to install and instantiate the asset chaincode component on a Hyperledger Fabric network.

## Environment
We have successfully tested and deployed this scaffolding chaincode component on [Hyperledger Fabric v1.4.4](https://hyperledger-fabric.readthedocs.io/en/release-1.4/releases.html), which requires [Go v1.12](https://golang.org/dl/).