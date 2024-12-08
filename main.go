package main

import (
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lil5/tigerbeetle_api/grpc"
	"github.com/lil5/tigerbeetle_api/rest"

	tigerbeetle_go "github.com/tigerbeetle/tigerbeetle-go"
	"github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

func main() {
	godotenv.Load()


	tbClusterIdStr := os.Getenv("TB_CLUSTER_ID")
	if tbClusterIdStr == "" {
		tbClusterIdStr = "0"
	}
    tbClusterId, _ := strconv.ParseUint(tbClusterIdStr, 10, 64)

	if host := os.Getenv("HOST"); host == "" {
		os.Setenv("HOST", "0.0.0.0")
	}

	if os.Getenv("TB_ADDRESSES") == "" {
		os.Setenv("TB_ADDRESSES", "127.0.0.1:3000")
	}
	tbAddressesArr := os.Getenv("TB_ADDRESSES")

	tbAddresses := strings.Split(tbAddressesArr, ",")

	slog.Info("Connecting to tigerbeetle cluster", "addresses", strings.Join(tbAddresses, ", "))

	// Connect to tigerbeetle
	tb, err := tigerbeetle_go.NewClient(types.ToUint128(uint64(tbClusterId)), tbAddresses)
	if err != nil {
		slog.Error("unable to connect to tigerbeetle", "err", err)
		os.Exit(1)
	}
	defer tb.Close()

	// Create server
	if os.Getenv("USE_GRPC") == "true" {
		if port, _ := strconv.Atoi(os.Getenv("PORT")); port == 0 {
			os.Setenv("PORT", "50051")
		}
		grpc.NewServer(tb)
	} else {
		if port, _ := strconv.Atoi(os.Getenv("PORT")); port == 0 {
			os.Setenv("PORT", "8000")
		}
		rest.NewServer(tb)
	}
}
