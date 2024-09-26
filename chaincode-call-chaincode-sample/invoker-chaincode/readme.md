- Ex: ./network.sh deployCC -ccn <chaincode_name> -ccp <path_to>/invoker-chaincode/ -ccl go:
```shell
./network.sh deployCC -ccn invoker -ccp /home/huy/Documents/chilley/fabric-samples/chaincode-call-chaincode-sample/invoker-chaincode/ -ccl go
```

- To invoke this chaincode, in folder test-network run:
```shell
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=${PWD}/../config
export TARGET_TLS_OPTIONS=(-o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt")
```
```shell
peer chaincode invoke "${TARGET_TLS_OPTIONS[@]}" -C mychannel -n invoker -c '{"Args":["ReadAssetFromA", "basic", "asset1"]}'
```