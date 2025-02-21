package data

import (
	"crypto/tls"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"kratos-minio/internal/conf"
	"net/http"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewMinioClient)

// Data .

type Data struct {
	minio  *minio.Client
	logger *log.Helper
}

// NewData .
func NewData(
	minio *minio.Client,
	logger log.Logger,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		minio: minio,
	}, cleanup, nil
}

func NewMinioClient(c *conf.Data) *minio.Client {
	// 初始化 Minio 客户端

	// 跳过证书验证, 如果证书正常, 删除该代码
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 跳过证书验证
		},
	}
	minioClient, err := minio.New(c.Minio.Endpoint, &minio.Options{
		Creds:     credentials.NewStaticV4(c.Minio.AccessKey, c.Minio.SecretKey, c.Minio.Token),
		Secure:    c.Minio.Secure,
		Transport: transport,
	})
	if err != nil {
		panic("new minio client fail: ")
	}

	return minioClient
}
