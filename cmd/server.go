package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/eifzed/ares/pb"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func GetListener(port string) (net.Listener, error) {
	var listener net.Listener

	EINHORN_FDS := os.Getenv("EINHORN_FDS")
	if EINHORN_FDS != "" {
		fds, err := strconv.Atoi(EINHORN_FDS)
		if err != nil {
			return nil, err
		}
		log.Println("using socket master, listening on", EINHORN_FDS)
		f := os.NewFile(uintptr(fds), "listener")
		listener, err = net.FileListener(f)
		if err != nil {
			log.Fatalln("error create listener", err)
			return nil, err
		}
		return listener, nil
	}
	return net.Listen("tcp4", port)
}

func ListenAndServeGRPC(port string, mod *modules) error {
	listener, err := GetListener(port)
	if err != nil {
		return err
	}

	srv := grpc.NewServer(grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(grpcauth.UnaryServerInterceptor(mod.GRPCMiddleware.AuthHandler))))
	// srv := grpc.NewServer()
	reflection.Register(srv)
	pb.RegisterOrderServiceServer(srv, mod.GRPCHandler.OrderHandler)
	pb.RegisterUserServiceServer(srv, mod.GRPCHandler.UserHandler)
	go func() {
		fmt.Println("listening with pid " + fmt.Sprint(os.Getpid()))
		srv.Serve(listener)
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("got signal", <-sigs)
	fmt.Println("shutting down..")
	srv.Stop()
	return nil
}
