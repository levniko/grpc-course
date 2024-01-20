package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-module/greet/greetpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//doUnaryCall(c)
	//doServerStreaming(c)
	doClientStreaming(c)
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC!")
	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Muhammet",
				LastName:  "Aslan",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Hilal",
				LastName:  "Aslan",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Rabia",
				LastName:  "Aslan",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Halil",
				LastName:  "Aslan",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ayşe",
				LastName:  "Aslan",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Remzi",
				LastName:  "Aslan",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Println("Sending request: ", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}
	fmt.Println("LongGreet response: ", resp)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC!")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Muhammet",
			LastName:  "Aslan",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doUnaryCall(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC!")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Muhammet",
			LastName:  "Aslan",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet RPC: %v", res.Result)
}
