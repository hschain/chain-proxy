package logic

import (
	"context"
	"log"

	"chainproxy/internal/svc"
	"chainproxy/internal/types"
	"chainproxy/utils"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExRatesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExRatesLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExRatesLogic {
	return ExRatesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExRatesLogic) ExRates() (*types.ExRatesResponse, error) {

	var resp types.ExRatesResponse

	if err := utils.DoHTTP("GET", nil, nil, "https://api.coingecko.com/api/v3/exchange_rates", &resp); err != nil {
		// l.Logger.Printf() // todo
		log.Printf("err:%v.\n", err)
		return &types.ExRatesResponse{}, nil
	}

	return &resp, nil

}
