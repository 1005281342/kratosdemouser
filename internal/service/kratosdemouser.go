package service

import (
	"context"
	"github.com/1005281342/kratosdemouser/internal/service/ability"
	"github.com/1005281342/kratosdemouser/internal/service/user"

	pb "github.com/1005281342/kratosdemouser/api/kratosdemouser"
)

func (s *KratosDemoUserService) CreateKratosDemoUser(ctx context.Context, req *pb.CreateKratosDemoUserRequest) (*pb.CreateKratosDemoUserReply, error) {
	var (
		rsp = &pb.CreateKratosDemoUserReply{}
		a   = user.ConstructorACreate(req, rsp, ability.WithUsecase(s.uc), ability.WithLogHelper(s.log))
		err error
	)
	if err = a.Check(ctx); err != nil {
		return nil, err
	}
	if err = a.Process(ctx); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (s *KratosDemoUserService) UpdateKratosDemoUser(ctx context.Context, req *pb.UpdateKratosDemoUserRequest) (*pb.UpdateKratosDemoUserReply, error) {
	var (
		rsp = &pb.UpdateKratosDemoUserReply{}
		a   = user.ConstructorAUpdate(req, rsp, ability.WithUsecase(s.uc), ability.WithLogHelper(s.log))
		err error
	)
	if err = a.Check(ctx); err != nil {
		return nil, err
	}
	if err = a.Process(ctx); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (s *KratosDemoUserService) DeleteKratosDemoUser(ctx context.Context, req *pb.DeleteKratosDemoUserRequest) (*pb.DeleteKratosDemoUserReply, error) {
	var (
		rsp = &pb.DeleteKratosDemoUserReply{}
		a   = user.ConstructorADelete(req, rsp, ability.WithUsecase(s.uc), ability.WithLogHelper(s.log))
		err error
	)
	if err = a.Check(ctx); err != nil {
		return nil, err
	}
	if err = a.Process(ctx); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (s *KratosDemoUserService) GetKratosDemoUser(ctx context.Context, req *pb.GetKratosDemoUserRequest) (*pb.GetKratosDemoUserReply, error) {
	var (
		rsp = &pb.GetKratosDemoUserReply{}
		a   = user.ConstructorAGet(req, rsp, ability.WithUsecase(s.uc), ability.WithLogHelper(s.log))
		err error
	)
	if err = a.Check(ctx); err != nil {
		return nil, err
	}
	if err = a.Process(ctx); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (s *KratosDemoUserService) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingRsp, error) {
	return &pb.PingRsp{Msg: req.GetMsg()}, nil
}
