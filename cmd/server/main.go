package main

import (
	"context"
	"database/sql"
	"github.com/egor-zakharov/goph-keeper/internal/config"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/createauthdata"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/createcard"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/createtextdata"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/deleteauthdata"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/deletecard"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/deletefile"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/deletetextdata"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/downloadfile"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/getauthdata"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/getcards"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/getfiles"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/gettextdata"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/signin"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/signup"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/subcribetochanges"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/updateauthdata"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/updatecard"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/updatetextdata"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/uploadfile"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/middleware"
	"github.com/egor-zakharov/goph-keeper/internal/migrator"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/server"
	authService "github.com/egor-zakharov/goph-keeper/internal/service/authdata"
	cardsService "github.com/egor-zakharov/goph-keeper/internal/service/cards"
	filesService "github.com/egor-zakharov/goph-keeper/internal/service/files"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
	textService "github.com/egor-zakharov/goph-keeper/internal/service/textdata"
	usersService "github.com/egor-zakharov/goph-keeper/internal/service/users"
	authStorage "github.com/egor-zakharov/goph-keeper/internal/storage/authdata"
	cardsStorage "github.com/egor-zakharov/goph-keeper/internal/storage/cards"
	filesStorage "github.com/egor-zakharov/goph-keeper/internal/storage/files"
	sessionStorage "github.com/egor-zakharov/goph-keeper/internal/storage/session"
	textStorage "github.com/egor-zakharov/goph-keeper/internal/storage/textdata"
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
	textStore := textStorage.New(db)
	session := sessionStorage.New()
	filesStore := filesStorage.New(db)

	//Service
	usersService := usersService.New(usersStore)
	cardsService := cardsService.New(cardsStore)
	authService := authService.New(authStore)
	textService := textService.New(textStore)
	notificationService := notification.New(session)
	filesService := filesService.New(filesStore)

	//Handlers
	signUpHandler := signup.New(usersService)
	signInHandler := signin.New(usersService)
	createCardHandler := createcard.New(cardsService, notificationService)
	getCardsHandler := getcards.New(cardsService)
	updateCardHandler := updatecard.New(cardsService, notificationService)
	deleteCardHandler := deletecard.New(cardsService, notificationService)
	createAuthDataHandler := createauthdata.New(authService, notificationService)
	getAuthDataHandler := getauthdata.New(authService)
	updateAuthDataHandler := updateauthdata.New(authService, notificationService)
	deleteAuthDataHandler := deleteauthdata.New(authService, notificationService)
	createTextDataHandler := createtextdata.New(textService, notificationService)
	getTextDataHandler := gettextdata.New(textService)
	updateTextDataHandler := updatetextdata.New(textService, notificationService)
	deleteTextDataHandler := deletetextdata.New(textService, notificationService)
	filesUploadHandler := uploadfile.New(filesService, notificationService)
	getFilesHandler := getfiles.New(filesService)
	deleteFileHandler := deletefile.New(filesService, notificationService)
	downloadFileHandler := downloadfile.New(filesService)

	subscribeToChangesHandler := subcribetochanges.New(notificationService)

	//Server
	keeperServer := server.New(
		signUpHandler,
		signInHandler,
		createCardHandler,
		authService,
		notificationService,
		getCardsHandler,
		updateCardHandler,
		deleteCardHandler,
		createAuthDataHandler,
		getAuthDataHandler,
		updateAuthDataHandler,
		deleteAuthDataHandler,
		subscribeToChangesHandler,
		createTextDataHandler,
		getTextDataHandler,
		updateTextDataHandler,
		deleteTextDataHandler,
		filesUploadHandler,
		getFilesHandler,
		deleteFileHandler,
		downloadFileHandler,
	)

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
