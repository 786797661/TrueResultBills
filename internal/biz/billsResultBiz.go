package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"

	"github.com/go-kratos/kratos/v2/log"
	v1 "ngBills/api/billsResult/v1"
)

// GreeterRepo is a Greater repo.
type NewBillsRepo interface {
	GetTrueResult(context.Context, *v1.GetTrueResultRequest) (*v1.GetTrueResultReply, error)
}

// GreeterUsecase is a Greeter usecase.
type NewBillsResultBizSer struct {
	repo NewBillsRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewBillsResultBiz(repo NewBillsRepo, logger log.Logger) *NewBillsResultBizSer {
	return &NewBillsResultBizSer{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *NewBillsResultBizSer) GetTrueResult(ctx context.Context, request *v1.GetTrueResultRequest) (*v1.GetTrueResultReply, error) {
	jsonbyte, err := json.Marshal(request)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	fmt.Println(string(jsonbyte))

	eg, ctx := errgroup.WithContext(ctx)

	reploy := v1.GetTrueResultReply{}

	eg.Go(func() error {
		egReploy, egerr := uc.repo.GetTrueResult(ctx, request)
		reploy = *egReploy

		return egerr
	})
	err = eg.Wait()

	return &reploy, err
}
