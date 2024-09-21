package main

import (
	"context"
	"database/sql"
	"github.com/egor-zakharov/goph-keeper/internal/config"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/signin"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/signup"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/middleware"
	"github.com/egor-zakharov/goph-keeper/internal/migrator"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/server"
	authService "github.com/egor-zakharov/goph-keeper/internal/service/auth"
	cardsService "github.com/egor-zakharov/goph-keeper/internal/service/cards"
	usersService "github.com/egor-zakharov/goph-keeper/internal/service/users"
	authStorage "github.com/egor-zakharov/goph-keeper/internal/storage/auth"
	cardsStorage "github.com/egor-zakharov/goph-keeper/internal/storage/cards"
	usersStorage "github.com/egor-zakharov/goph-keeper/internal/storage/users"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.NewConfig()
	conf.ParseFlag()
	err := logger.Initialize(conf.FlagLogLevel)
	if err != nil {
		panic(err)
	}
	logger.Log().Sugar().Infow("Logging info", "level", conf.FlagLogLevel)

	//Migrator
	db, err := sql.Open("pgx", conf.FlagDB)
	if err != nil {
		logger.Log().Sugar().Errorw("Open DB migrations crashed: ", zap.Error(err))
		panic(err)
	}
	logger.Log().Sugar().Debugw("Running DB migrations")
	newMigrator := migrator.New(db)
	err = newMigrator.Run()
	if err != nil {
		logger.Log().Sugar().Errorw("Migrations crashed with error: ", zap.Error(err))
		panic(err)
	}

	//Storage
	//DB
	db, err = sql.Open("pgx", conf.FlagDB)
	if err != nil {
		logger.Log().Sugar().Errorw("Open DB storage crashed: ", zap.Error(err))
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logger.Log().Sugar().Errorw("Cannot ping DB: ", zap.Error(err))
		panic(err)
	}
	//Storages
	usersStore := usersStorage.New(db)
	cardsStore := cardsStorage.New(db)
	authStore := authStorage.New(db)

	//Service
	usersService := usersService.New(usersStore)
	cardsService := cardsService.New(cardsStore)
	authService := authService.New(authStore)

	//Handlers
	signUpHandler := signup.New(usersService)
	signInHandler := signin.New(usersService)

	//Server
	keeperServer := server.New(cardsService, signUpHandler, signInHandler, authService)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()
	go func() {
		listen, err := net.Listen("tcp", conf.FlagRunGRPCAddr)
		if err != nil {
			panic(err)
		}
		s := grpc.NewServer(
			grpc.UnaryInterceptor(middleware.AuthInterceptor),
			grpc.MaxConcurrentStreams(20),
			grpc.StreamInterceptor(middleware.StreamAuthInterceptor))
		pb.RegisterGophKeeperServerServer(s, keeperServer)
		reflection.Register(s)

		err = s.Serve(listen)
		logger.Log().Sugar().Infow("Running grpc keeperServer", "address", conf.FlagRunGRPCAddr)
		if err != nil {
			panic(err)
		}
	}()
	<-ctx.Done()
}
