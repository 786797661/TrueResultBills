package service

import (
	"context"

	v1 "ngBills/api/billsResult/v1"
	"ngBills/internal/biz"
)

// GreeterService is a greeter service.
type NewBillsResultBizServiceSer struct {
	v1.UnimplementedBillsResultServer

	uc *biz.NewBillsResultBizSer
}

// NewGreeterService new a greeter service.
func NewBillsResultBizService(uc *biz.NewBillsResultBizSer) *NewBillsResultBizServiceSer {
	return &NewBillsResultBizServiceSer{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (uc *NewBillsResultBizServiceSer) GetTrueResult(ctx context.Context, request *v1.GetTrueResultRequest) (*v1.GetTrueResultReply, error) {
	return uc.uc.GetTrueResult(ctx, request)

}
