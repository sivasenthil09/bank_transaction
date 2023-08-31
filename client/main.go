package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sivasenthil09/bank_transaction/banktrans"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect :%v", err)
	}
	defer conn.Close()

	//	client := pb.NewCustomerServiceClient(conn)
	client := pb.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &pb.Customer{
		CustomerID: 123,
		FirstName:  "siva",
		LastName:   "D",
		BankID:     123,
		Balance:    10000,
		CreatedAt:  "2023.08.30",
		UpdatedAt:  "2023.08.30",
		IsActive:   true,
		
	})
	response2, err := client.CreateCustomer(context.Background(), &pb.Customer{
		CustomerID: 124,
		FirstName:  "senthil",
		LastName:   "D",
		BankID:     1234,
		Balance:    10000,
		CreatedAt:  "2023.08.30",
		UpdatedAt:  "2023.08.30",
		IsActive:   true,
		
	})
    
	if err != nil {
		log.Fatalf("Failed to create customer :%v", err)
	}
	response3,err :=client.Transaction(context.Background(),&pb.Customer{
		CustomerID : 124,
	})

	fmt.Printf("Response from customer : %v\n", response.CustomerID)
	fmt.Printf("Response from customer : %v\n", response2.CustomerID)
	fmt.Printf("Response from customer : %v\n", response3.CustomerID)
}
