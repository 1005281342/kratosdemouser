package userclient

import (
	"context"
	"github.com/1005281342/kratosdemouser/api/kratosdemouser"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"testing"
)

var (
	gClient kratosdemouser.KratosDemoUserClient
	ctx     = context.Background()
)

func init() {
	var conn, _ = grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("127.0.0.1:8888"),
	)
	gClient = kratosdemouser.NewKratosDemoUserClient(conn)
}

func TestCreate(t *testing.T) {
	var req = &kratosdemouser.CreateKratosDemoUserRequest{Name: "jeson"}
	var rsp, err = gClient.CreateKratosDemoUser(ctx, req)
	if err != nil {
		t.Fatalf("create err: %s", err.Error())
	}
	t.Logf("create req: %+v, rsp: %+v", req, rsp)
	// user_client_test.go:29: create req: name:"jeson", rsp: user_info:{name:"jeson"  id:1}
}

func TestUpdate(t *testing.T) {
	var req = &kratosdemouser.UpdateKratosDemoUserRequest{UserInfo: &kratosdemouser.UserInfo{Name: "x123", Id: 1}}
	var rsp, err = gClient.UpdateKratosDemoUser(ctx, req)
	if err != nil {
		t.Fatalf("update err: %s", err.Error())
	}
	t.Logf("update req: %+v, rsp: %+v", req, rsp)
	//user_client_test.go:39: update req: user_info:{name:"x123" id:1}, rsp:
}

func TestGet(t *testing.T) {
	var req = &kratosdemouser.GetKratosDemoUserRequest{Id: 2}
	var rsp, err = gClient.GetKratosDemoUser(ctx, req)
	if err != nil {
		t.Fatalf("get err: %s", err.Error())
	}
	t.Logf("get req: %+v, rsp: %+v", req, rsp)
	// user_client_test.go:49: get req: id:1, rsp: user_info:{name:"x123" id:1}
}

func TestDelete(t *testing.T) {
	var req = &kratosdemouser.DeleteKratosDemoUserRequest{Id: 1}
	var rsp, err = gClient.DeleteKratosDemoUser(ctx, req)
	if err != nil {
		t.Fatalf("delete err: %s", err.Error())
	}
	t.Logf("delete req: %+v, rsp: %+v", req, rsp)
	// user_client_test.go:59: delete req: id:1, rsp:
	//  user_client_test.go:47: get err: rpc error: code = Unknown desc = record not found
	//  user_client_test.go:49: get req: id:2, rsp: user_info:{name:"jx"  id:2}
}
