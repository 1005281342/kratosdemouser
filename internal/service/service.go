package service

import (
	pb "github.com/1005281342/kratosdemouser/api/kratosdemouser"
	"github.com/1005281342/kratosdemouser/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewKratosDemoUserService)

type KratosDemoUserService struct {
	pb.UnimplementedKratosDemoUserServer

	uc  *biz.KratosDemoUserUsecase
	log *log.Helper
}

func NewKratosDemoUserService(uc *biz.KratosDemoUserUsecase, logger log.Logger) *KratosDemoUserService {
	return &KratosDemoUserService{uc: uc, log: log.NewHelper(logger)}
}
