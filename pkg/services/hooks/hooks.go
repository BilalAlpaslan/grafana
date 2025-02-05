package hooks

import (
	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/models"
	contextmodel "github.com/grafana/grafana/pkg/services/contexthandler/model"
)

type IndexDataHook func(indexData *dtos.IndexViewData, req *contextmodel.ReqContext)

type LoginHook func(loginInfo *models.LoginInfo, req *contextmodel.ReqContext)

type HooksService struct {
	indexDataHooks []IndexDataHook
	loginHooks     []LoginHook
}

func ProvideService() *HooksService {
	return &HooksService{}
}

func (srv *HooksService) AddIndexDataHook(hook IndexDataHook) {
	srv.indexDataHooks = append(srv.indexDataHooks, hook)
}

func (srv *HooksService) RunIndexDataHooks(indexData *dtos.IndexViewData, req *contextmodel.ReqContext) {
	for _, hook := range srv.indexDataHooks {
		hook(indexData, req)
	}
}

func (srv *HooksService) AddLoginHook(hook LoginHook) {
	srv.loginHooks = append(srv.loginHooks, hook)
}

func (srv *HooksService) RunLoginHook(loginInfo *models.LoginInfo, req *contextmodel.ReqContext) {
	for _, hook := range srv.loginHooks {
		hook(loginInfo, req)
	}
}
