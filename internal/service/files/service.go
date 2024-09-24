package files

import (
	"context"
	"fmt"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/storage/files"
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

func (s *service) Add(ctx context.Context, stream pb.GophKeeperServer_UploadFileServer) (*models.FileData, error) {
	userID := ctx.Value(auth.UserIdContextKey).(string)

	file := models.NewFile()
	var fileSize uint32
	fileSize = 0
	var meta string
	defer func() {
		if err := file.Close(); err != nil {
		}
	}()
	for {
		req, err := stream.Recv()
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
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		chunk := req.GetFile().GetData()
		fileSize += uint32(len(chunk))

		if err := file.Write(chunk); err != nil {
			return nil, err
		}

		return createdFile, stream.SendAndClose(&pb.UploadFileResponse{Id: createdFile.ID})
	}

	return nil, nil
}

func (s *service) Read(ctx context.Context, userID string) (*[]models.FileData, error) {
	return s.fileStorage.Read(ctx, userID)
}

func (s *service) Delete(ctx context.Context, id string, userID string) error {
	return s.fileStorage.Delete(ctx, id, userID)
}

func (s *service) Download(in *pb.DownloadFileRequest, stream pb.GophKeeperServer_DownloadFileServer) error {
	ctx := stream.Context()
	userID := ctx.Value(auth.UserIdContextKey).(string)
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
		return fmt.Errorf("%s", "not fount")
	}
	//id + name
	path := filepath.Join("uploaded", fmt.Sprintf("%s_%s", fileID, fileName))

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
