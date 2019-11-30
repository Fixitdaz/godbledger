package main

import (
	"context"
	"log"
	"math/big"
	"time"

	pb "github.com/darcys22/godbledger/proto"
	"google.golang.org/grpc"

	"github.com/urfave/cli"
)

var commandDeleteTransaction = cli.Command{
	Name:      "delete",
	Usage:     "deletes a single transaction",
	ArgsUsage: "[]",
	Description: `
`,
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {

		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := pb.NewTransactorClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := &pb.DeleteRequest{
			Identifier: transactionLines,
			Signature:  t.Signature,
		}
		r, err := client.DeleteTransaction(ctx, req)
		if err != nil {
			log.Fatalf("could not delete: %v", err)
		}
		log.Printf("Response: %s", r.GetMessage())
		return nil

		return nil
	},
}
