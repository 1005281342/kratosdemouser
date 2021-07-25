package user

import (
	"context"
	pb "github.com/1005281342/kratosdemouser/api/kratosdemouser"
	"github.com/1005281342/kratosdemouser/internal/biz"
	"github.com/1005281342/kratosdemouser/internal/errcode"
	"github.com/1005281342/kratosdemouser/internal/service/ability"
)

type AGet struct {
	AUser
	opt *ability.Option
	req *pb.GetKratosDemoUserRequest
	rsp *pb.GetKratosDemoUserReply
}

func ConstructorAGet(req *pb.GetKratosDemoUserRequest, rsp *pb.GetKratosDemoUserReply, opts ...ability.OptionFunc) AGet {
	var a = AGet{req: req, rsp: rsp, opt: &ability.Option{}}
	for _, opt := range opts {
		opt(a.opt)
	}
	return a
}

func (a *AGet) Check(ctx context.Context) error {
	var logger = a.opt.Log.WithContext(ctx)
	if a.req == nil || a.req.GetId() <= 0 {
		logger.Errorf("%s", errcode.ErrCheckParam.Error())
		return errcode.ErrCheckParam
	}
	if a.opt == nil || a.opt.Uc == nil {
		logger.Errorf("%s", errcode.ErrCheckParam.Error())
		return errcode.ErrCheckParam
	}
	return nil
}

func (a *AGet) Process(ctx context.Context) error {
	var (
		logger = a.opt.Log.WithContext(ctx)
		err    error
		t      biz.TKratosDemoUser
	)
	if t, err = a.opt.Uc.Get(ctx, a.req.GetId()); err != nil {
		logger.Errorf("Get err: %s", err.Error())
		return err
	}
	logger.Debugf("Get %+v", a.req)
	a.rsp.UserInfo = a.TKratosDemoUser2UserInfo(&t)
	return nil
}
