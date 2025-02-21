package biz

import (
	"context"
	v1 "kratos-minio/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}
type UploadMethod int32
type OssUploadUrlRequest struct {
	Method      UploadMethod
	ContentType *string
	BucketName  *string
	FilePath    *string
	FileName    *string
}

type OssUploadUrlResponse struct {
	UploadUrl   string
	DownloadUrl string
	BucketName  *string
	ObjectName  string
	FormData    map[string]string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	OssUploadUrl(ctx context.Context, req *OssUploadUrlRequest) (*OssUploadUrlResponse, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GreeterUsecase) OssUploadUrl(ctx context.Context, req *OssUploadUrlRequest) (*OssUploadUrlResponse, error) {
	uc.log.WithContext(ctx).Debugf("OssUploadUrl: %v", req)
	return uc.repo.OssUploadUrl(ctx, req)
}
