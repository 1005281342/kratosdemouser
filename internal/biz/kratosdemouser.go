package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// TKratosDemoUser user表
type TKratosDemoUser struct {
	FID   int32  `gorm:"column:f_id;primary_key"`
	FName string `gorm:"column:f_name"`
	//CreateAt time.Time `gorm:"column:f_created_at;" json:"create_at"` // 创建时间
	//UpdateAt time.Time `gorm:"column:f_updated_at;" json:"update_at"` // 更新时间
}

type KratosDemoUserRepo interface {
	CreateKratosDemoUser(context.Context, *TKratosDemoUser) error
	UpdateKratosDemoUser(context.Context, *TKratosDemoUser) error
	DeleteKratosDemoUser(ctx context.Context, id int32) error
	GetKratosDemoUser(ctx context.Context, id int32) (TKratosDemoUser, error)
}

type KratosDemoUserUsecase struct {
	repo KratosDemoUserRepo
	log  *log.Helper
}

func NewKratosDemoUserUsecase(repo KratosDemoUserRepo, logger log.Logger) *KratosDemoUserUsecase {
	return &KratosDemoUserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *KratosDemoUserUsecase) Create(ctx context.Context, g *TKratosDemoUser) error {
	return uc.repo.CreateKratosDemoUser(ctx, g)
}

func (uc *KratosDemoUserUsecase) Update(ctx context.Context, g *TKratosDemoUser) error {
	return uc.repo.UpdateKratosDemoUser(ctx, g)
}

func (uc *KratosDemoUserUsecase) Delete(ctx context.Context, id int32) error {
	return uc.repo.DeleteKratosDemoUser(ctx, id)
}

func (uc *KratosDemoUserUsecase) Get(ctx context.Context, id int32) (TKratosDemoUser, error) {
	return uc.repo.GetKratosDemoUser(ctx, id)
}
