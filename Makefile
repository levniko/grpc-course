# Define the protoc command as a variable for easy modification and readability
PROTOC_CMD = protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --go-grpc_opt=require_unimplemented_servers=false

# Define a target for generating protobuf files
generate-proto:
	$(PROTOC_CMD) greet/greetpb/greet.proto

# You can also add other targets here as needed
