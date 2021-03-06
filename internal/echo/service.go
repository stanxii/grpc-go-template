package echo

import (
	"context"
	"math/rand"
	"time"

	"github.com/hatlonely/grpc-go-template/api/echoapi"
	"github.com/sirupsen/logrus"
)

// Service 实现 Add 服务
type Service struct{}

// Do 接口实现
func (s *Service) Do(ctx context.Context, request *echoapi.Request) (*echoapi.Response, error) {
	// 50% 概率 sleep，模拟超时场景
	if rand.Int()%2 == 0 {
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	response := &echoapi.Response{
		Message: request.Message,
	}
	logrus.WithField("request", request).WithField("response", response).Info()
	return response, nil
}
