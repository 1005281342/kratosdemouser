package user

import (
	"context"
	pb "github.com/1005281342/kratosdemouser/api/kratosdemouser"
	"github.com/1005281342/kratosdemouser/internal/biz"
	"github.com/1005281342/kratosdemouser/internal/errcode"
	"github.com/1005281342/kratosdemouser/internal/service/ability"
)

type ACreate struct {
	AUser
	opt *ability.Option
	req *pb.CreateKratosDemoUserRequest
	rsp *pb.CreateKratosDemoUserReply
}

func ConstructorACreate(req *pb.CreateKratosDemoUserRequest, rsp *pb.CreateKratosDemoUserReply, opts ...ability.OptionFunc) ACreate {
	var a = ACreate{req: req, rsp: rsp, opt: &ability.Option{}}
	for _, opt := range opts {
		opt(a.opt)
	}
	return a
}

func (a *ACreate) Check(ctx context.Context) error {
	var logger = a.opt.Log.WithContext(ctx)
	if a.req == nil || a.req.GetName() == "" {
		logger.Errorf("%s", errcode.ErrCheckParam.Error())
		return errcode.ErrCheckParam
	}
	if a.opt == nil || a.opt.Uc == nil {
		logger.Errorf("%s", errcode.ErrCheckParam.Error())
		return errcode.ErrCheckParam
	}
	return nil
}

func (a *ACreate) Process(ctx context.Context) error {
	var (
		logger = a.opt.Log.WithContext(ctx)
		user   = &biz.TKratosDemoUser{
			FName: a.req.GetName(),
		}
		err error
	)
	if err = a.opt.Uc.Create(ctx, user); err != nil {
		logger.Errorf("Create err: %s", err.Error())
		return err
	}
	a.rsp.UserInfo = a.TKratosDemoUser2UserInfo(user)
	logger.Debugf("Create get id: %d, user", a.rsp.GetUserInfo().GetId())
	return nil
}
