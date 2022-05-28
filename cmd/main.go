package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/eifzed/ares/internal/config"
	"github.com/eifzed/ares/internal/handler"
	orderHandler "github.com/eifzed/ares/internal/handler/grpc/order"
	userHandler "github.com/eifzed/ares/internal/handler/grpc/user"
	"github.com/eifzed/ares/internal/handler/middleware/auth"
	orderDB "github.com/eifzed/ares/internal/repo/order"
	userDB "github.com/eifzed/ares/internal/repo/user"
	orderUC "github.com/eifzed/ares/internal/usecase/order"
	userUC "github.com/eifzed/ares/internal/usecase/user"
	"github.com/eifzed/ares/lib/database/mongodb/transaction"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	pemClientCA, err := ioutil.ReadFile("./files/etc/ares-secret/client-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("./files/etc/ares-secret/server-cert.pem", "./files/etc/ares-secret/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	secrets := config.GetSecretes()
	if secrets == nil {
		log.Fatal("failed to get secretes")
		return
	}
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg.Secrets = secrets
	client, err := getDBConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}
	aresDB := client.Database("ares-db")

	orderDB := orderDB.GetNewOrderDB(&orderDB.OrderDBOption{
		DB: aresDB,
	})

	userDB := userDB.GetNewUserDB(&userDB.Options{
		DB: aresDB,
	})

	tx := transaction.GetNewMongoDBTransaction(&transaction.Options{
		Client: client,
		DBName: "ares-db",
	})

	orderUC := orderUC.GetNewOrderUC(&orderUC.Options{
		OrderDB: orderDB,
		UserDB:  userDB,
		Config:  cfg,
		TX:      tx,
	})
	userUC := userUC.GetNewUserUC(&userUC.Options{
		UserDB: userDB,
		Config: cfg,
		TX:     tx,
	})

	orderHandler := orderHandler.GetNewOrderHandler(&orderHandler.Option{
		OrderUC: orderUC,
		Config:  cfg,
	})

	userHandler := userHandler.GetNewUserHandler(&userHandler.Option{
		UserUC: userUC,
	})

	authMiddleware := auth.NewAuthModule(&auth.Options{
		JWTCertificate: cfg.Secrets.Data.JWTCertificate,
		RouteRoles:     cfg.RouteRoles,
		Config:         cfg,
		UserDB:         userDB,
	})

	modules := getNewModules(&modules{
		GRPCHandler: handler.GRPCHandler{
			OrderHandler: orderHandler,
			UserHandler:  userHandler,
		},
		GRPCMiddleware: authMiddleware, // TODO: add logging middleware
	})

	ListenAndServeGRPC(cfg.Server.HTTP.Address, modules)
}
