package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "shuai190060/rest_grpc_k8s_deploy/gRPC_client/account_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	rand.New(rand.NewSource(1000000))
}

func main() {

	// // client service to write to postgresql
	startGRPCClientService()

}

const (
	address = "a6e66a42ba87a48c196c1742042dd4d9-f96108f28e54c2db.elb.us-east-1.amazonaws.com:50051"
)

func startGRPCClientService() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// create new client with this connection
	c := pb.NewAccountManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//create 100 random account
	var random_accounts [][]string

	for i := 0; i < 10; i++ {
		subSlice := []string{fmt.Sprintf("FirstName_%d", rand.Intn(10000)), fmt.Sprintf("LastName_%d", rand.Intn(10000))}
		random_accounts = append(random_accounts, subSlice)
	}

	for _, Name := range random_accounts {
		_, err := c.CreateAccount(ctx, &pb.NewAccount{
			FirstName: Name[0],
			LastName:  Name[1],
		})
		if err != nil {
			log.Fatalf("could not create new account:%v", err)
		}
	}

	// first_name := "bob"
	// last_name := "jack"
	// r, err := c.CreateAccount(ctx, &pb.NewAccount{
	// 	FirstName: first_name,
	// 	LastName:  last_name,
	// })
	// if err != nil {
	// 	log.Fatalf("could not create new account:%v", err)
	// }
	// log.Printf(`Account details:
	// First_name: %s
	// Last_name: %s
	// Number: %d
	// `, r.GetFirstName(), r.GetLastName(), r.GetNumber())

	params := &pb.GetAccountParams{}
	res_acc_list, err := c.GetAccount(ctx, params)
	if err != nil {
		log.Fatalf("could not retrieve accounts: %v", err)
	}
	log.Print("\nuser list is:\n")
	fmt.Printf("r.GetAccount():%v", res_acc_list)

}