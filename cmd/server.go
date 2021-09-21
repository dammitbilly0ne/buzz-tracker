package cmd

import (
	"log"
	"net"

	"github.com/dammitbilly0ne/buzz-tracker/internal/handlers"
	"github.com/dammitbilly0ne/buzz-tracker/protos"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("starting server")
		port := "5000"
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		handler, err := handlers.NewAlpha(&handlers.AlphaConfig{})
		if err != nil {
			log.Fatal("err returned from handlers.NewAlpha()")
		}

		protos.RegisterBuzzTrackerAPIServer(s, handler)

		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}
