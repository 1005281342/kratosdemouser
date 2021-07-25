package user

import (
	"context"
	pb "github.com/1005281342/kratosdemouser/api/kratosdemouser"
	"github.com/1005281342/kratosdemouser/internal/errcode"
	"github.com/1005281342/kratosdemouser/internal/service/ability"
)

type AUpdate struct {
	AUser
	opt *ability.Option
	req *pb.UpdateKratosDemoUserRequest
	rsp *pb.UpdateKratosDemoUserReply
}

func ConstructorAUpdate(req *pb.UpdateKratosDemoUserRequest, rsp *pb.UpdateKratosDemoUserReply, opts ...ability.OptionFunc) AUpdate {
	var a = AUpdate{req: req, rsp: rsp, opt: &ability.Option{}}
	for _, opt := range opts {
		opt(a.opt)
	}
	return a
}

func (a *AUpdate) Check(ctx context.Context) error {
	var logger = a.opt.Log.WithContext(ctx)
	if a.req == nil || a.req.GetUserInfo() == nil {
		logger.Errorf("%s", errcode.ErrCheckParam.Error())
		return errcode.ErrCheckParam
	}
	if a.opt == nil || a.opt.Uc == nil {
		logger.Errorf("%s", errcode.ErrCheckParam.Error())
		return errcode.ErrCheckParam
	}
	return nil
}

func (a *AUpdate) Process(ctx context.Context) error {
	var (
		logger = a.opt.Log.WithContext(ctx)
		user   = a.UserInfo2TKratosDemoUser(a.req.GetUserInfo())
		err    error
	)
	if err = a.opt.Uc.Update(ctx, user); err != nil {
		logger.Errorf("Update err: %s", err.Error())
		return err
	}
	logger.Debugf("Update %+v", a.req.GetUserInfo())
	return nil
}
