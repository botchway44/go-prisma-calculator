gen :
	@protoc --proto_path=proto \
  --go_out=. \
  --go_opt=module=go-prisma-calculator \
  --go-grpc_out=. \
  --go-grpc_opt=module=go-prisma-calculator \
  proto/*.proto