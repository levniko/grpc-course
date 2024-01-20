#!/bin/bash

protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --go-grpc_opt=require_unimplemented_servers=false greet/greetpb/greet.proto
