package gapi

import (
	"context"
	"database/sql"

	db "github.com/LouisChen-TW/simple_bank/db/sqlc"
	"github.com/LouisChen-TW/simple_bank/pb"
	"github.com/LouisChen-TW/simple_bank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "cannot find the user")
		}
		return nil, status.Errorf(codes.Internal, "server error: %s", err)
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect password")
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to create access token")
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to create refresh token")
	}

	mtdt := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to create session")
	}

	rsp := &pb.LoginUserResponse{
		User:                 convertUser(user),
		SessionId:            session.ID.String(),
		AccessToken:          accessToken,
		AccessTokenExpireAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:         refreshToken,
		RefreshTokenExpireAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return rsp, nil
}
