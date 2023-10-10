package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
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
	// startGRPCClientService()

	REST_api()

}

const (
	address = "a43f8d28bbef84bf7afd781cef4df17b-ccce9392309cd12b.elb.us-east-1.amazonaws.com:50051"
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

type Account struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func REST_api() {
	for i := 1; i <= 10; i++ {
		account := Account{
			FirstName: fmt.Sprintf("First_name_%d", rand.Intn(10000)),
			LastName:  fmt.Sprintf("Last_name_%d", rand.Intn(10000)),
		}

		body, err := json.Marshal(account)
		if err != nil {
			log.Fatalf("Error marshaling account: %v", err)
		}

		create_address := "http://a43f8d28bbef84bf7afd781cef4df17b-ccce9392309cd12b.elb.us-east-1.amazonaws.com:3000/account"

		resp, err := http.Post(create_address, "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Fatalf("Error sending POST request: %v", err)
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		log.Println(string(respBody))

	}
}
