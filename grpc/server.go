package grpc

import (
	"fmt"
	"log/slog"
	"net"
	"os"

    tigerbeetlepb "github.com/lil5/tigerbeetle_api/proto/tigerbeetle"  // Add alias here
    tigerbeetle_go "github.com/tigerbeetle/tigerbeetle-go"
    "google.golang.org/grpc"
    "google.golang.org/grpc/health"
    healthpb "google.golang.org/grpc/health/grpc_health_v1"
    "google.golang.org/grpc/reflection"
)

func NewServer(tb tigerbeetle_go.Client) {
	s := grpc.NewServer()


	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	if err != nil {
		slog.Error("Failed to listen", "error", err)
		os.Exit(1)
	}

    // Add health check service
    healthServer := health.NewServer()
    healthpb.RegisterHealthServer(s, healthServer)
    healthServer.SetServingStatus("tigerbeetle.TigerBeetle", healthpb.HealthCheckResponse_SERVING)

	tigerbeetlepb.RegisterTigerBeetleServer(s, &Server{TB: tb})
	reflection.Register(s)

	slog.Info("Server listening at", "address", lis.Addr())
	if err := s.Serve(lis); err != nil {
		slog.Error("Failed to serve:", "error", err)
		os.Exit(1)
	}

	slog.Info("Server exiting")
}