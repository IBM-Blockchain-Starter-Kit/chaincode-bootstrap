[![Build Status - Master](https://travis-ci.org/rolivieri/chaincode-bootstrap.svg?branch=master)](https://travis-ci.org/rolivieri/chaincode-bootstrap/builds)

# Scaffolding repository for use in Blockchain Starter Kit

## Chaincode Development Instructions
* Create a new component directory under the [chaincode](/chaincode) folder for each component.
* Populate this component directory with files to be deployed.
* The [ping](/chaincode/ping) directory is provided as a chaincode component example. 

## Chaincode Deployment Instructions
Populate the [deployment configuration](deploy_config.json) JSON file with information about network targets, inculding organizations, channels, and peers. Please note that included in this repostiory, there is a deployment configuration example file with the default network architecture to install and instantiate the ping chaincode component on a Hyperledger Fabric network.

## Environment
We have successfully tested and deployed this scaffolding chaincode component on [Hyperledger Fabric v1.1.0](https://hyperledger-fabric.readthedocs.io/en/release-1.1/releases.html), which requires [Go v1.9](https://golang.org/dl/).