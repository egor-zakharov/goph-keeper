package client

import (
	"context"
	"crypto/tls"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/utils"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"os"
	"path/filepath"
)

type client struct {
	client        pb.GophKeeperClient
	token         string
	sessionID     string
	serverAddress string
	conn          *grpc.ClientConn
	isAvailable   bool
}

func New(serverAddress string) Client {
	return &client{
		serverAddress: serverAddress,
	}
}

func (c *client) Connect() error {
	conn, err := grpc.NewClient(c.serverAddress,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))
	if err != nil {
		logger.Log().Sugar().Errorw("failed connect to server", err)
		return err
	}

	c.conn = conn
	c.client = pb.NewGophKeeperClient(conn)

	c.isAvailable = true

	return nil
}

func (c *client) SignUp(login string, password string) error {
	resp, err := c.client.SignUp(context.Background(), &pb.SignUpRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		logger.Log().Sugar().Errorw("Sign up client error", err)
		return err
	}
	c.token = resp.GetToken()

	return nil
}

func (c *client) SignIn(login string, password string) error {
	resp, err := c.client.SignIn(context.Background(), &pb.SignInRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		logger.Log().Sugar().Errorw("Sign in client error", err)
		return err
	}
	c.token = resp.GetToken()

	return nil
}

func (c *client) CreateAuthData(ctx context.Context, meta string, login string, password string) (*models.AuthData, error) {
	resp, err := c.client.CreateAuthData(c.setToken(ctx, c.token), &pb.CreateAuthDataRequest{
		Data: &pb.CreateAuthDataRequest_Data{
			Meta:     meta,
			Login:    login,
			Password: password,
		},
	})
	if err != nil {
		logger.Log().Sugar().Errorw("Create auth data client error", err)
		return nil, err
	}

	data := models.AuthData{
		ID:       resp.Id,
		Meta:     meta,
		Login:    login,
		Password: password,
	}

	return &data, nil
}

func (c *client) CreateCard(ctx context.Context, number string, expirationDate string, holderName string, CVV string) (*models.Card, error) {
	resp, err := c.client.CreateCard(c.setToken(ctx, c.token), &pb.CreateCardRequest{
		Card: &pb.CreateCardRequest_Card{
			Number:         number,
			ExpirationDate: expirationDate,
			HolderName:     holderName,
			Cvv:            CVV,
		},
	})
	if err != nil {
		logger.Log().Sugar().Errorw("Create card client error", err)
		return nil, err
	}

	var data = models.Card{
		ID:             resp.CardID,
		Number:         number,
		ExpirationDate: expirationDate,
		HolderName:     holderName,
		CVV:            CVV,
	}
	return &data, nil
}

func (c *client) CreateTextData(ctx context.Context, meta string, data string) (*models.TextData, error) {
	resp, err := c.client.CreateTextData(c.setToken(ctx, c.token), &pb.CreateTextDataRequest{
		Data: &pb.CreateTextDataRequest_Data{
			Meta: meta,
			Text: data,
		},
	})
	if err != nil {
		logger.Log().Sugar().Errorw("Create text data client error", err)
		return nil, err
	}

	var textData = models.TextData{
		ID:   resp.Id,
		Meta: meta,
		Text: data,
	}
	return &textData, nil
}

func (c *client) DeleteAuthData(ctx context.Context, id string) error {
	_, err := c.client.DeleteAuthData(c.setToken(ctx, c.token), &pb.DeleteAuthDataRequest{Id: id})
	if err != nil {
		logger.Log().Sugar().Errorw("Delete auth data client error", err)
		return err
	}
	return nil
}

func (c *client) DeleteCard(ctx context.Context, id string) error {
	_, err := c.client.DeleteCard(c.setToken(ctx, c.token), &pb.DeleteCardRequest{Id: id})
	if err != nil {
		logger.Log().Sugar().Errorw("Delete card client error", err)
		return err
	}
	return nil
}

func (c *client) DeleteFile(ctx context.Context, id string) error {
	_, err := c.client.DeleteFile(c.setToken(ctx, c.token), &pb.DeleteFileRequest{Id: id})
	if err != nil {
		logger.Log().Sugar().Errorw("Delete file client error", err)
		return err
	}
	return nil
}

func (c *client) DeleteTextData(ctx context.Context, id string) error {
	_, err := c.client.DeleteTextData(c.setToken(ctx, c.token), &pb.DeleteTextDataRequest{Id: id})
	if err != nil {
		logger.Log().Sugar().Errorw("Delete text data client error", err)
		return err
	}
	return nil
}

func (c *client) DownloadFile(ctx context.Context, id string, fileName string) error {
	stream, err := c.client.DownloadFile(c.setToken(ctx, c.token), &pb.DownloadFileRequest{Id: id})
	if err != nil {
		logger.Log().Sugar().Errorw("Download file client error", err)
		return err
	}

	localFile := utils.NewFile()
	err = localFile.SetFile(fileName, "downloads")
	if err != nil {
		logger.Log().Sugar().Errorw("Set file error", err)
		return err
	}
	defer localFile.Close()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Log().Sugar().Errorw("receive chunk error", err)
			return err
		}

		if err := localFile.Write(resp.GetData()); err != nil {
			logger.Log().Sugar().Errorw("write file error", err)
			return err
		}
	}

	return nil
}

func (c *client) GetAuthData(ctx context.Context) (*[]models.AuthData, error) {
	resp, err := c.client.GetAuthData(c.setToken(ctx, c.token), &pb.GetAuthDataRequest{})
	if err != nil {
		logger.Log().Sugar().Errorw("Get auth data client error", err)
		return nil, err
	}
	response := make([]models.AuthData, 0, len(resp.Data))
	for _, data := range resp.Data {
		r := models.AuthData{
			ID:       data.Id,
			Meta:     data.Meta,
			Login:    data.Login,
			Password: data.Password,
		}
		response = append(response, r)
	}
	return &response, nil
}

func (c *client) GetCards(ctx context.Context) (*[]models.Card, error) {
	resp, err := c.client.GetCards(c.setToken(ctx, c.token), &pb.GetCardsRequest{})
	if err != nil {
		logger.Log().Sugar().Errorw("Get cards client error", err)
		return nil, err
	}
	response := make([]models.Card, 0, len(resp.Cards))
	for _, data := range resp.Cards {
		r := models.Card{
			ID:             data.Id,
			Number:         data.Number,
			ExpirationDate: data.ExpirationDate,
			HolderName:     data.HolderName,
			CVV:            data.Cvv,
		}
		response = append(response, r)
	}
	return &response, nil
}

func (c *client) GetFiles(ctx context.Context) (*[]models.FileData, error) {
	resp, err := c.client.GetFiles(c.setToken(ctx, c.token), &pb.GetFilesRequest{})
	if err != nil {
		logger.Log().Sugar().Errorw("Get files client error", err)
		return nil, err
	}
	response := make([]models.FileData, 0, len(resp.Files))
	for _, data := range resp.Files {
		r := models.FileData{
			ID:   data.Id,
			Name: data.Name,
			Meta: data.Meta,
		}
		response = append(response, r)
	}
	return &response, nil
}

func (c *client) GetTextData(ctx context.Context) (*[]models.TextData, error) {
	resp, err := c.client.GetTextData(c.setToken(ctx, c.token), &pb.GetTextDataRequest{})
	if err != nil {
		logger.Log().Sugar().Errorw("Get text data client error", err)
		return nil, err
	}
	response := make([]models.TextData, 0, len(resp.Data))
	for _, data := range resp.Data {
		r := models.TextData{
			ID:   data.Id,
			Text: data.Text,
			Meta: data.Meta,
		}
		response = append(response, r)
	}
	return &response, nil
}

func (c *client) SubscribeToChanges(ctx context.Context) (grpc.ServerStreamingClient[pb.SubscribeToChangesResponse], error) {
	return c.client.SubscribeToChanges(c.setToken(ctx, c.token), &pb.SubscribeToChangesRequest{})
}

func (c *client) UpdateAuthData(ctx context.Context, id string, meta string, login string, password string) error {
	_, err := c.client.UpdateAuthData(c.setToken(ctx, c.token), &pb.UpdateAuthDataRequest{
		Data: &pb.UpdateAuthDataRequest_Data{
			Id:       id,
			Meta:     meta,
			Login:    login,
			Password: password,
		}})
	if err != nil {
		logger.Log().Sugar().Errorw("Update auth data client error", err)
		return err
	}
	return nil
}

func (c *client) UpdateCard(ctx context.Context, id string, number string, expirationDate string, holderName string, CVV string) error {
	_, err := c.client.UpdateCard(c.setToken(ctx, c.token), &pb.UpdateCardRequest{
		Card: &pb.UpdateCardRequest_Card{
			Id:             id,
			Number:         number,
			ExpirationDate: expirationDate,
			HolderName:     holderName,
			Cvv:            CVV,
		},
	})
	if err != nil {
		logger.Log().Sugar().Errorw("Update card client error", err)
		return err
	}
	return nil
}

func (c *client) UpdateTextData(ctx context.Context, id string, meta string, data string) error {
	_, err := c.client.UpdateTextData(c.setToken(ctx, c.token), &pb.UpdateTextDataRequest{
		Data: &pb.UpdateTextDataRequest_Data{
			Id:   id,
			Meta: meta,
			Text: data,
		},
	})
	if err != nil {
		logger.Log().Sugar().Errorw("Update text data client error", err)
		return err
	}
	return nil
}

func (c *client) UploadFile(ctx context.Context, filePath string, meta string) (*models.FileData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.Log().Sugar().Errorw("Upload file file open error", err)
		return nil, err
	}
	defer file.Close()

	cleanFileName := filepath.Base(filePath)

	stream, err := c.client.UploadFile(c.setToken(ctx, c.token))
	if err != nil {
		logger.Log().Sugar().Errorw("Upload file client error", err)
		return nil, err
	}

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Log().Sugar().Errorw("Upload read file error", err)
			return nil, err
		}

		err = stream.Send(&pb.UploadFileRequest{
			File: &pb.UploadFileRequest_File{
				Data:     buffer[:n],
				Filename: cleanFileName,
				Meta:     meta,
			},
		})

		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Log().Sugar().Errorw("Upload file send data error", err)
			return nil, err
		}
	}

	if err := stream.CloseSend(); err != nil {
		logger.Log().Sugar().Errorw("Upload file close error", err)
		return nil, err
	}

	res, err := stream.Recv()
	if err != nil {
		logger.Log().Sugar().Errorw("Update file stream res err", err)
		return nil, err
	}

	return &models.FileData{
		ID:   res.Id,
		Name: cleanFileName,
		Meta: meta,
	}, nil
}

func (c *client) setToken(ctx context.Context, jwt string) context.Context {
	md := metadata.New(map[string]string{
		"authorization": jwt,
		"session":       c.sessionID,
	})

	return metadata.NewOutgoingContext(ctx, md)
}
