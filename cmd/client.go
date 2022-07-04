/*
client.go is a client api for sending Pokemon trading requests to a gRPC server
*/
package cmd

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "github.com/bccorb/pkg/gts/globalTradeSystem"
	"github.com/spf13/cobra"
)

const (
	address = "localhost:9000"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Query the gRPC server",

	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		client := pb.NewGlobalTradeSystemClient(conn)

		stream, _ := client.GetTradeList(context.Background())
		waitc := make(chan struct{})
		go func() {
			for {
				in, err := stream.Recv()
				if err == io.EOF {
					// read done.
					close(waitc)
					return
				}
				if err != nil {
					log.Fatalf("Failed to receive a Trade Response : %v", err)
				}
				log.Printf("We found a %s to exchange for your %s!", in.GetOfferedPokemon(), in.GetRequestedPokemon())
			}
		}()

		if err := stream.Send(&pb.TradeRequest{OfferedPokemon: "Zubat", RequestedPokemon: "Snorlax"}); err != nil {
			log.Fatalf("Failed to request pokemon trade: %v", err)
		}

		stream.CloseSend()
		<-waitc
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
