SERVERADDR="localhost:50051"
GOPATH=~/go
HTTPS=false

# check https://github.com/improbable-eng/grpc-web/tree/master/go/grpcwebproxy

if [ "$HTTPS" = true ] ; then
  #https & http
  $GOPATH/bin/grpcwebproxy \
    --server_tls_cert_file=./misc/localhost.crt \
    --server_tls_key_file=./misc/localhost.key \
    --backend_addr=$SERVERADDR \
    --use_websockets \
    --allow_all_origins
else
  #only http
  $GOPATH/bin/grpcwebproxy \
    --backend_addr=$SERVERADDR \
    --use_websockets \
    --run_tls_server=false \
    --allow_all_origins
fi
