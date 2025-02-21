package data

import (
	"context"
	"github.com/minio/minio-go/v7"
	"time"

	"kratos-minio/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	defaultExpiryTime = time.Second * 24 * 60 * 60 // 1 day
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

func (r *greeterRepo) OssUploadUrl(ctx context.Context, req *biz.OssUploadUrlRequest) (*biz.OssUploadUrlResponse, error) {
	expiry := defaultExpiryTime

	policy := minio.NewPostPolicy()
	_ = policy.SetBucket(*req.BucketName)
	_ = policy.SetKey(*req.FileName)
	_ = policy.SetExpires(time.Now().UTC().Add(expiry))
	presignedURL, formData, err := r.data.minio.PresignedPostPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}

	url, err := r.data.minio.PresignedPutObject(ctx, *req.BucketName, *req.FileName, expiry)
	if err != nil {
		return nil, err
	}

	return &biz.OssUploadUrlResponse{
		UploadUrl:   presignedURL.String(),
		DownloadUrl: url.String(),
		BucketName:  req.BucketName,
		ObjectName:  *req.FileName,
		FormData:    formData,
	}, nil
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
