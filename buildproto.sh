GOEnv="github.com/vinaycharlie07/ai"
PROTOPATH="aitech/proto"



protoc --go_out=. \
    --go_opt=module=$GOEnv \
    --go-grpc_out=. --go-grpc_opt=module=$GOEnv \
    $PROTOPATH/*.proto