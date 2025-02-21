package service

import (
	"context"

	v1 "kratos-minio/api/helloworld/v1"
	"kratos-minio/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedFileServiceServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) OssUploadUrl(ctx context.Context, req *v1.OssUploadUrlRequest) (*v1.OssUploadUrlResponse, error) {
	response, err := s.uc.OssUploadUrl(ctx, &biz.OssUploadUrlRequest{
		Method:      biz.UploadMethod(req.Method),
		ContentType: req.ContentType,
		BucketName:  req.BucketName,
		FilePath:    req.FilePath,
		FileName:    req.FileName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.OssUploadUrlResponse{
		UploadUrl:   response.UploadUrl,
		DownloadUrl: response.DownloadUrl,
		BucketName:  response.BucketName,
		ObjectName:  response.ObjectName,
		FormData:    response.FormData,
	}, nil

}
