- Lưu ý đặt tên chaincode: viết thường, ngăn từ bởi dấu "-"
- Trong file chaincode/smartcontract.go phải đặt một tên biến chaincodeName chứa tên chaincode
- Ex: ./network.sh deployCC -ccn <chaincode_name> -ccp <path_to>/basic-chaincode/ -ccl go:

    ```shell
    ./network.sh deployCC -ccn basic -ccp /home/huy/Documents/chilley/fabric-samples/chaincode-call-chaincode-sample/basic-chaincode/ -ccl go
    ```