package middleware

import (
	"context"
	"fmt"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func AuthInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	switch info.FullMethod {
	case gophkeeper.GophKeeper_SignUp_FullMethodName,
		gophkeeper.GophKeeper_SignIn_FullMethodName:
		logger.Log().Debug("No protected method", zap.String("method", info.FullMethod))

		return handler(ctx, req)
	}

	logger.Log().Debug("Protected method")
	logger.Log().Debug(info.FullMethod)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Log().Debug("couldn't extract metadata from req")

		return nil, fmt.Errorf("%w", status.Error(codes.Internal, "couldn't extract metadata from req"))
	}

	authHeaders, ok := md[auth.AuthHeader]
	if !ok || len(authHeaders) != 1 {
		logger.Log().Debug("authorization not exists")

		return nil, status.Errorf(codes.Unauthenticated, "authorization not exists")
	}

	token := strings.TrimPrefix(authHeaders[0], auth.Bearer)
	if token == "" {
		logger.Log().Debug("token empty or not valid")

		return nil, status.Errorf(codes.Unauthenticated, "token empty or not valid")
	}

	if isValid, err := auth.IsValidToken(token); err != nil || !isValid {
		logger.Log().Debug("token is not valid")

		return nil, status.Errorf(codes.Unauthenticated, "token empty or not valid")
	}

	userID := auth.GetUserID(token)
	if userID == "" {
		logger.Log().Debug("cannot get userID")

		return nil, status.Errorf(codes.Unauthenticated, "token empty or not valid")
	}

	sessionID := auth.GetSessionID(token)
	if sessionID == "" {
		logger.Log().Debug("cannot get sessionID")

		return nil, status.Errorf(codes.Unauthenticated, "token empty or not valid")
	}

	userCtx := context.WithValue(ctx, auth.UserIdContextKey, userID)
	sessionCtx := context.WithValue(userCtx, auth.SessionIdContextKey, sessionID)

	return handler(sessionCtx, req)
}

func StreamAuthInterceptor(srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	// Получаем контекст из потокового сервера
	ctx := ss.Context()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Internal, "couldn't extract metadata from req")
	}

	authHeaders, ok := md[auth.AuthHeader]
	if !ok || len(authHeaders) != 1 {
		logger.Log().Debug("authorization not exists")
		return status.Error(codes.Unauthenticated, "authorization not exists")
	}

	token := strings.TrimPrefix(authHeaders[0], auth.Bearer)
	if token == "" {
		logger.Log().Debug("token empty or not valid")

		return status.Error(codes.Unauthenticated, "token empty or not valid")
	}

	if isValid, err := auth.IsValidToken(token); err != nil || !isValid {
		logger.Log().Debug("token is not valid")

		return status.Errorf(codes.Unauthenticated, "token empty or not valid")
	}

	userID := auth.GetUserID(token)
	if userID == "" {
		logger.Log().Debug("cannot get userID")

		return status.Errorf(codes.Unauthenticated, "token empty or not valid")
	}

	sessionID := auth.GetSessionID(token)
	if sessionID == "" {
		logger.Log().Debug("cannot get sessionID")

		return status.Errorf(codes.Unauthenticated, "token empty or not valid")
	}

	userCtx := context.WithValue(ctx, auth.UserIdContextKey, userID)
	sessionCtx := context.WithValue(userCtx, auth.SessionIdContextKey, sessionID)
	wrappedStream := grpc_middleware.WrapServerStream(ss)
	wrappedStream.WrappedContext = sessionCtx

	return handler(srv, wrappedStream)
}
