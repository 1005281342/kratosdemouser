package ability

import (
	"context"
	"github.com/1005281342/kratosdemouser/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type IAbility interface {
	// Check 参数校验
	Check(ctx context.Context) error
	// Process 数据处理
	Process(ctx context.Context) error
}

type Option struct {
	Uc  *biz.KratosDemoUserUsecase
	Log *log.Helper
}

// OptionFunc 可选参数处理
type OptionFunc func(option *Option)

// WithUsecase ...
func WithUsecase(uc *biz.KratosDemoUserUsecase) OptionFunc {
	return func(option *Option) {
		option.Uc = uc
	}
}

// WithLogHelper ...
func WithLogHelper(log *log.Helper) OptionFunc {
	return func(option *Option) {
		option.Log = log
	}
}
