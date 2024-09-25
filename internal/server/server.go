package server

import (
	"context"
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
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	authService "github.com/egor-zakharov/goph-keeper/internal/service/authdata"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
	"sync"
)

type GophKeeperServer struct {
	pb.UnimplementedGophKeeperServer
	signUp *signup.Handler
	signIn *signin.Handler

	createCard  *createcard.Handler
	getCards    *getcards.Handler
	updateCards *updatecard.Handler
	deleteCard  *deletecard.Handler

	createAuthData *createauthdata.Handler
	getAuthData    *getauthdata.Handler
	updateAuthData *updateauthdata.Handler
	deleteAuthData *deleteauthdata.Handler

	createTextData *createtextdata.Handler
	getTextData    *gettextdata.Handler
	updateTextData *updatetextdata.Handler
	deleteTextData *deletetextdata.Handler

	subscribe *subcribetochanges.Handler

	uploadFile   *uploadfile.Handler
	getFiles     *getfiles.Handler
	deleteFile   *deletefile.Handler
	downloadFile *downloadfile.Handler

	authService  authService.Service
	notification notification.Service
	rwMutex      sync.RWMutex
}

func New(
	signUp *signup.Handler,
	signIn *signin.Handler,
	createCard *createcard.Handler,
	authService authService.Service,
	notification notification.Service,
	getCards *getcards.Handler,
	updateCards *updatecard.Handler,
	deleteCard *deletecard.Handler,
	createAuthData *createauthdata.Handler,
	getAuthData *getauthdata.Handler,
	updateAuthData *updateauthdata.Handler,
	deleteAuthData *deleteauthdata.Handler,
	subscribe *subcribetochanges.Handler,
	createTextData *createtextdata.Handler,
	getTextData *gettextdata.Handler,
	updateTextData *updatetextdata.Handler,
	deleteTextData *deletetextdata.Handler,
	uploadFile *uploadfile.Handler,
	getFiles *getfiles.Handler,
	deleteFile *deletefile.Handler,
	downloadFile *downloadfile.Handler,
) *GophKeeperServer {
	return &GophKeeperServer{
		signUp:         signUp,
		signIn:         signIn,
		createCard:     createCard,
		authService:    authService,
		notification:   notification,
		getCards:       getCards,
		updateCards:    updateCards,
		deleteCard:     deleteCard,
		createAuthData: createAuthData,
		getAuthData:    getAuthData,
		updateAuthData: updateAuthData,
		deleteAuthData: deleteAuthData,
		subscribe:      subscribe,
		createTextData: createTextData,
		getTextData:    getTextData,
		updateTextData: updateTextData,
		deleteTextData: deleteTextData,
		uploadFile:     uploadFile,
		getFiles:       getFiles,
		deleteFile:     deleteFile,
		downloadFile:   downloadFile,
	}
}

func (s *GophKeeperServer) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return s.signUp.Handle(ctx, in)
}

func (s *GophKeeperServer) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	return s.signIn.Handle(ctx, in)
}

func (s *GophKeeperServer) CreateCard(ctx context.Context, in *pb.CreateCardRequest) (*pb.CreateCardResponse, error) {
	return s.createCard.Handle(ctx, in)
}

func (s *GophKeeperServer) GetCards(ctx context.Context, in *pb.GetCardsRequest) (*pb.GetCardsResponse, error) {
	return s.getCards.Handle(ctx, in)
}

func (s *GophKeeperServer) UpdateCard(ctx context.Context, in *pb.UpdateCardRequest) (*pb.UpdateCardResponse, error) {
	return s.updateCards.Handle(ctx, in)
}

func (s *GophKeeperServer) DeleteCard(ctx context.Context, in *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
	return s.deleteCard.Handle(ctx, in)
}

func (s *GophKeeperServer) CreateAuthData(ctx context.Context, in *pb.CreateAuthDataRequest) (*pb.CreateAuthDataResponse, error) {
	return s.createAuthData.Handle(ctx, in)
}

func (s *GophKeeperServer) GetAuthData(ctx context.Context, in *pb.GetAuthDataRequest) (*pb.GetAuthDataResponse, error) {
	return s.getAuthData.Handle(ctx, in)
}

func (s *GophKeeperServer) UpdateAuthData(ctx context.Context, in *pb.UpdateAuthDataRequest) (*pb.UpdateAuthDataResponse, error) {
	return s.updateAuthData.Handle(ctx, in)
}

func (s *GophKeeperServer) DeleteAuthData(ctx context.Context, in *pb.DeleteAuthDataRequest) (*pb.DeleteAuthDataResponse, error) {
	return s.deleteAuthData.Handle(ctx, in)
}

func (s *GophKeeperServer) CreateConfTextData(ctx context.Context, in *pb.CreateConfTextDataRequest) (*pb.CreateConfTextDataResponse, error) {
	return s.createTextData.Handle(ctx, in)
}

func (s *GophKeeperServer) GetConfTextData(ctx context.Context, in *pb.GetConfTextDataRequest) (*pb.GetConfTextDataResponse, error) {
	return s.getTextData.Handle(ctx, in)
}

func (s *GophKeeperServer) UpdateConfTextData(ctx context.Context, in *pb.UpdateConfTextDataRequest) (*pb.UpdateConfTextDataResponse, error) {
	return s.updateTextData.Handle(ctx, in)
}

func (s *GophKeeperServer) DeleteConfTextData(ctx context.Context, in *pb.DeleteConfTextDataRequest) (*pb.DeleteConfTextDataResponse, error) {
	return s.deleteTextData.Handle(ctx, in)
}

func (s *GophKeeperServer) SubscribeToChanges(in *pb.SubscribeToChangesRequest, stream pb.GophKeeper_SubscribeToChangesServer) error {
	return s.subscribe.SubscribeToChanges(in, stream)
}

func (s *GophKeeperServer) UploadFile(stream pb.GophKeeper_UploadFileServer) error {
	return s.uploadFile.Handle(stream)
}

func (s *GophKeeperServer) GetFiles(ctx context.Context, in *pb.GetFilesRequest) (*pb.GetFilesResponse, error) {
	return s.getFiles.Handle(ctx, in)
}

func (s *GophKeeperServer) DeleteFile(ctx context.Context, in *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	return s.deleteFile.Handle(ctx, in)
}

func (s *GophKeeperServer) DownloadFile(in *pb.DownloadFileRequest, stream pb.GophKeeper_DownloadFileServer) error {
	return s.downloadFile.Handle(in, stream)
}
