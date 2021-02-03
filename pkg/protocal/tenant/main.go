package tenant

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/bm-krishna-source/tenant-service/pkg/api/tenant"
	"google.golang.org/grpc"
)

func BootStrapServer(ctx context.Context, tenantServer tenant.TenantServer, port string) error {
	/*
		Regiseter Service:
			Create new grpc server pass to tenant RegisterXXXServer
	*/

	server := grpc.NewServer()
	tenant.RegisterTenantServer(server, tenantServer)

	// do graceful shoutdown

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	go func() {
		for range channel {
			log.Println("Shout down grpc server.")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()
	// server serving on net.Listner
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Failed to listen tcp port")
		return err
	}
	log.Println("Tenant server running on port", port)
	return server.Serve(listen)
}
