package user

import (
	pb "github.com/1005281342/kratosdemouser/api/kratosdemouser"
	"github.com/1005281342/kratosdemouser/internal/biz"
)

type AUser struct{}

func (a *AUser) TKratosDemoUser2UserInfo(t *biz.TKratosDemoUser) *pb.UserInfo {
	return &pb.UserInfo{Id: t.FID, Name: t.FName}
}

func (a *AUser) UserInfo2TKratosDemoUser(t *pb.UserInfo) *biz.TKratosDemoUser {
	return &biz.TKratosDemoUser{FID: t.GetId(), FName: t.GetName()}
}
