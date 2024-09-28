package files

import (
	"context"
	"fmt"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/storage/files"
	"github.com/egor-zakharov/goph-keeper/internal/utils"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
	"io"
	"os"
	"path/filepath"
)

type service struct {
	fileStorage files.Storage
}

func New(fileStorage files.Storage) Service {
	return &service{fileStorage: fileStorage}
}

func (s *service) Add(ctx context.Context, stream pb.GophKeeper_UploadFileServer) (*models.FileData, error) {
	userID := ctx.Value(auth.UserIDContextKey).(string)

	file := utils.NewFile()
	var fileSize uint32
	fileSize = 0
	var meta string
	defer file.Close()
	for {
		req, err := stream.Recv()
		if err != nil {
			return nil, err
		}
		meta = req.GetFile().GetMeta()
		fileData := models.FileData{
			Name: req.GetFile().GetFilename(),
			Meta: meta,
		}
		createdFile, err := s.fileStorage.Create(ctx, fileData, userID)
		if err != nil {
			return nil, err
		}

		if file.FilePath == "" {
			errSetFile := file.SetFile(fmt.Sprintf("%s_%s", createdFile.ID, createdFile.Name), "uploaded")
			if errSetFile != nil {
				return nil, err
			}
		}
		chunk := req.GetFile().GetData()
		fileSize += uint32(len(chunk))

		err = file.Write(chunk)
		if err != nil {
			return nil, err
		}

		return createdFile, stream.Send(&pb.UploadFileResponse{Id: createdFile.ID})
	}
}

func (s *service) Read(ctx context.Context, userID string) (*[]models.FileData, error) {
	return s.fileStorage.Read(ctx, userID)
}

func (s *service) Delete(ctx context.Context, id string, userID string) error {
	return s.fileStorage.Delete(ctx, id, userID)
}

func (s *service) Download(in *pb.DownloadFileRequest, stream pb.GophKeeper_DownloadFileServer) error {
	ctx := stream.Context()
	userID := ctx.Value(auth.UserIDContextKey).(string)
	fileID := in.GetId()
	files, err := s.fileStorage.Read(ctx, userID)
	if err != nil {
		return err
	}
	var fileName string
	for _, data := range *files {
		if data.ID == fileID {
			fileName = data.Name
			break
		}
	}
	if fileName == "" {
		return fmt.Errorf("%s", "not found")
	}
	path := filepath.Join("uploads", fmt.Sprintf("%s_%s", fileID, fileName))

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var totalBytesStreamed int64

	for totalBytesStreamed < fileSize {
		shard := make([]byte, 1024*1024)
		bytesRead, err := f.Read(shard)
		if err == io.EOF {

			break
		}

		if err != nil {
			return err
		}

		if err := stream.Send(&pb.DownloadFileResponse{
			Data: shard,
		}); err != nil {
			return err
		}
		totalBytesStreamed += int64(bytesRead)
	}
	return nil
}
