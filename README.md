Fabric 1.1
Go 1.9.6

## 教材準備  

##### 下載專案
    git clone https://github.com/BingHongLi/invoke_chaincode_lbh.git

##### 下載子專案
    cd invoke_chaincode_lbh
    git submodule update --init --recursive

##### 複製合約進入環境內
    cp -r chaincode_basic_tutorial_lbh fabric-samples/chaincode/
    cp -r invoke_chaincode_lbh fabric-samples/chaincode/

#####  啟用環境
    cd fabric-samples/basic-network
    sh start.sh
    docker-compose up -d cli

## 合約準備

#####  安裝invoke合約
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode install -n invoke_chaincode_lbh -v 1.0 -p github.com/invoke_chaincode_lbh
##### 激活invoke合約
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n invoke_chaincode_lbh -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')"

##### 安裝asset合約
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode install -n chaincode_basic_tutorial_lbh  -v 1.0 -p github.com/chaincode_basic_tutorial_lbh
##### 激活 asset合約
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n chaincode_basic_tutorial_lbh -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')"

## 調用合約

##### 調用asset 合約, 生成模擬資產
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n chaincode_basic_tutorial_lbh -c '{"function":"simulateTA","Args":[""]}'


##### 調用invoke合約，查詢模擬資產
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n invoke_chaincode_lbh -c '{"function":"invokeOtherChaincodeSetAsset","Args":["chaincode_basic_tutorial_lbh","mychannel","getTA","1"]}'

##### 調用invoke合約，新增模擬資產
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n invoke_chaincode_lbh -c '{"function":"invokeOtherChaincodeSetAsset","Args":["chaincode_basic_tutorial_lbh","mychannel","putTA","4","20181111","lbh","500","singleValentine"]}'

##### 調用invoke合約，查詢模擬資產
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n invoke_chaincode_lbh -c '{"function":"invokeOtherChaincodeSetAsset","Args":["chaincode_basic_tutorial_lbh","mychannel","getTA","4"]}'

##### 調用asset 合約，查詢資產
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n chaincode_basic_tutorial_lbh -c '{"function":"getTA","Args":["4","20181111","lbh","500","singleValentine"]}'
