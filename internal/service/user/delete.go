package user

import (
	"context"
	pb "github.com/1005281342/kratosdemouser/api/kratosdemouser"
	"github.com/1005281342/kratosdemouser/internal/errcode"
	"github.com/1005281342/kratosdemouser/internal/service/ability"
)

type ADelete struct {
	AUser
	opt *ability.Option
	req *pb.DeleteKratosDemoUserRequest
	rsp *pb.DeleteKratosDemoUserReply
}

func ConstructorADelete(req *pb.DeleteKratosDemoUserRequest, rsp *pb.DeleteKratosDemoUserReply, opts ...ability.OptionFunc) ADelete {
	var a = ADelete{req: req, rsp: rsp, opt: &ability.Option{}}
	for _, opt := range opts {
		opt(a.opt)
	}
	return a
}

func (a *ADelete) Check(ctx context.Context) error {
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

func (a *ADelete) Process(ctx context.Context) error {
	var (
		logger = a.opt.Log.WithContext(ctx)
		err    error
	)
	if err = a.opt.Uc.Delete(ctx, a.req.GetId()); err != nil {
		logger.Errorf("Delete err: %s", err.Error())
		return err
	}
	logger.Debugf("Delete %+v", a.req)
	return nil
}
