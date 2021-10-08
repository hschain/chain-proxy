package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"chainproxy/internal/logic"
	"chainproxy/internal/svc"
)

func ExRatesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewExRatesLogic(r.Context(), ctx)
		resp, err := l.ExRates()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
