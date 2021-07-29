package data

import (
	"context"
	"github.com/1005281342/kratosdemouser/internal/biz"
	"github.com/1005281342/kratosdemouser/internal/errcode"
	"github.com/go-kratos/kratos/v2/log"
)

type kratosDemoUserRepo struct {
	data *Data
	log  *log.Helper
}

// NewKratosDemoUserRepo .
func NewKratosDemoUserRepo(data *Data, logger log.Logger) biz.KratosDemoUserRepo {
	return &kratosDemoUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *kratosDemoUserRepo) CreateKratosDemoUser(ctx context.Context, g *biz.TKratosDemoUser) error {
	if g == nil || g.FName == "" {
		return errcode.ErrCheckParam
	}
	var res = r.data.db.Create(g)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		return errcode.ErrCheckRowsAffected
	}
	return nil
}

func (r *kratosDemoUserRepo) UpdateKratosDemoUser(ctx context.Context, g *biz.TKratosDemoUser) error {
	if g == nil || g.FID <= 0 {
		return errcode.ErrCheckParam
	}
	var res = r.data.db.Debug().Model(&biz.TKratosDemoUser{}).
		Where("f_id = ? and f_del_flag = 0", g.FID).Update("f_name", g.FName)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		return errcode.ErrCheckRowsAffected
	}
	return nil
}

func (r *kratosDemoUserRepo) DeleteKratosDemoUser(ctx context.Context, id int32) error {
	if id <= 0 {
		return errcode.ErrCheckParam
	}
	var res = r.data.db.Model(&biz.TKratosDemoUser{}).Where("f_id = ? and f_del_flag = 0", id).Update("f_del_flag", 1)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		return errcode.ErrCheckRowsAffected
	}
	return nil
}

func (r *kratosDemoUserRepo) GetKratosDemoUser(ctx context.Context, id int32) (biz.TKratosDemoUser, error) {
	var g biz.TKratosDemoUser
	if id <= 0 {
		return g, errcode.ErrCheckParam
	}
	var res = r.data.db.Model(&biz.TKratosDemoUser{}).Where("f_id = ? and f_del_flag = 0", id).First(&g)
	if res.Error != nil {
		return g, res.Error
	}
	return g, nil
}
