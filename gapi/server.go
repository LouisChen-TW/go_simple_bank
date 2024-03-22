package gapi

import (
	"fmt"

	db "github.com/LouisChen-TW/simple_bank/db/sqlc"
	"github.com/LouisChen-TW/simple_bank/pb"
	"github.com/LouisChen-TW/simple_bank/token"
	"github.com/LouisChen-TW/simple_bank/util"
)

// Server serves gRPC requests for banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer create a new gRPC service and set up routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
