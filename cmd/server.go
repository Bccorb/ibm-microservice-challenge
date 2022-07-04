/*
server.go is a server api for receiving and confirming Pokemon trading requests from a client
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/spf13/cobra"

	pb "github.com/bccorb/pkg/gts/globalTradeSystem"
	"google.golang.org/grpc"
)

const (
	port = 9000
)

// server is used to implement gts.GlobalTradeSystemServer.
type Server struct {
	pb.UnimplementedGlobalTradeSystemServer
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption

		grpcServer := grpc.NewServer(opts...)

		// Register services
		pb.RegisterGlobalTradeSystemServer(grpcServer, &Server{})

		log.Printf("GRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

// GetTradeList implements GlobalTradeSystem.GlobalTradeSystemServer
func (s *Server) GetTradeList(stream pb.GlobalTradeSystem_GetTradeListServer) error {
	avaliablePokemon := [5]pb.TradeResponse{
		{RequestedPokemon: "Pikachu", OfferedPokemon: "Snorlax"},
		{RequestedPokemon: "Kadabra", OfferedPokemon: "Kadabra"},
		{RequestedPokemon: "Mewtow", OfferedPokemon: "Mewtow"},
		{RequestedPokemon: "Zubat", OfferedPokemon: "Pidgy"},
		{RequestedPokemon: "Haunter", OfferedPokemon: "Haunter"},
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Searching for a trade... %s for a %s", in.RequestedPokemon, in.OfferedPokemon)
		for i, pokemon := range avaliablePokemon {
			if pokemon.OfferedPokemon == in.RequestedPokemon {
				log.Printf("Found a %s to trade for a %s", in.RequestedPokemon, in.OfferedPokemon)
				if err := stream.Send(&avaliablePokemon[i]); err != nil {
					return err
				} else {
					log.Printf("Couldn't find a %s to trade for a %s", in.RequestedPokemon, in.OfferedPokemon)
				}
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
