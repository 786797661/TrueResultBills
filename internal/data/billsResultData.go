package data

import (
	"context"
	"encoding/json"
	v1 "ngBills/api/billsResult/v1"
	nnres "ngBills/internal/data/nn/nnmodell"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"ngBills/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.NewBillsRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (uc *greeterRepo) GetTrueResult(ctx context.Context, request *v1.GetTrueResultRequest) (*v1.GetTrueResultReply, error) {
	return uc.getTrueResultHttp(ctx, request)
}

/**
通过http调用阿里云接口
*/
func (uc *greeterRepo) getTrueResultHttp(ctx context.Context, request *v1.GetTrueResultRequest) (*v1.GetTrueResultReply, error) {
	res := httpGet(request)
	//Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
	//第二个参数必须是指针，否则无法接收解析的数据，如stu仍为空对象Student{}
	reploy := transFormData(res)
	return reploy, nil
}

func httpGet(request *v1.GetTrueResultRequest) nnres.NNResult {
	return HttpGet(request)
}

func transFormData(res nnres.NNResult) *v1.GetTrueResultReply {
	if res.Code != 1 {
		errorReply := v1.GetTrueResultReply{
			Data:       nil,
			Success:    "false",
			SubMessage: "第三方验真提示：" + res.Msg,
			SubCode:    "-1",
		}
		return &errorReply
	}
	resmodle := nnres.NNFpYzResult{}

	json.Unmarshal([]byte(res.Data), &resmodle)

	kprq, _ := time.Parse("2006-01-02", resmodle.Kprq)
	bz := ""
	if resmodle.Bz != "" {
		bz = strings.ReplaceAll(resmodle.Bz, "<br/>", "")
	}
	trueResult := v1.GetTrueResultReplyMainModel{
		Fphm:   resmodle.Fphm,
		Fpdm:   resmodle.Fpdm,
		Fpzl:   resmodle.Fpzl,
		Kprq:   kprq.Format("20060102"),
		Gfmc:   resmodle.Gfmc,
		Gfsh:   resmodle.Gfsh,
		Gfdzdh: resmodle.Gfdzdh,
		Gfyhzh: resmodle.Gfyhzh,
		Xfsh:   resmodle.Xfsh,
		Xfmc:   resmodle.Xfmc,
		Xfdzdh: resmodle.Xfdzdh,
		Je:     resmodle.Je,
		Se:     resmodle.Se,
		Hjje:   resmodle.Hjje,
		Jqm:    resmodle.Jqm,
		Jym:    resmodle.Jym,
		Bz:     bz,
	}
	switch resmodle.Zfbz {
	case "N":
		trueResult.Zfbz = "N"
	case "Y":
		trueResult.Zfbz = "Y"
	case "0":
		trueResult.Zfbz = "N"
	case "2":
		trueResult.Zfbz = "Y"
	case "3":
		trueResult.Zfbz = "Y"
	default:
		trueResult.Zfbz = "N"
	}
	switch resmodle.Hongc {
	case "N":
		trueResult.Hongc = "N"
	case "Y":
		trueResult.Hongc = "Y"
	default:
		trueResult.Zfbz = "N"
	}
	dts := make([]*v1.GetTrueResultReplyNodeDd, 0)
	for _, detail := range resmodle.Details {
		dt := v1.GetTrueResultReplyNodeDd{
			Fphm: resmodle.Fphm,
			Fpdm: resmodle.Fpdm,
			Fpzl: resmodle.Fpzl,
			Spmc: detail.Spmc,
			Ggxh: detail.Ggxh,
			Jldw: detail.Jldw,
			Spsl: detail.Spsl,
			Dj:   detail.Dj,
			Je:   detail.Je,
			Slv:  detail.Slv,
			Se:   detail.Se,
		}
		dts = append(dts, &dt)
	}
	trueResult.InvoiceList = dts

	trueResultReply := v1.GetTrueResultReply{
		Data:       &trueResult,
		Success:    "true",
		SubMessage: "验真成功",
		SubCode:    "00",
	}
	return &trueResultReply
}
