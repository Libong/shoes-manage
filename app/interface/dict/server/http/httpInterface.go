package http

import (
	"libong/common/context"
	"libong/common/server/http"
	"shoe-manager/app/interface/dict/api"
)

func addDict(ctx *http.Context) error {
	var (
		req = &api.AddDictReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.AddDict(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func updateDict(ctx *http.Context) error {
	var (
		req = &api.UpdateDictReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.UpdateDict(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func deleteDict(ctx *http.Context) error {
	var (
		req = &api.DeleteDictReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.DeleteDict(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func searchDictionariesPage(ctx *http.Context) error {
	var (
		req  = &api.SearchDictionariesPageReq{}
		resp *api.SearchDictionariesPageResp
		err  error
	)
	err = ctx.MarshalGet(req)
	if err != nil {
		return err
	}
	if resp, err = svc.SearchDictionariesPage(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}

func dictById(ctx *http.Context) error {
	var (
		req  = &api.DictByIdReq{}
		resp *api.DictByIdResp
		err  error
	)
	err = ctx.MarshalGet(req)
	if err != nil {
		return err
	}
	if resp, err = svc.DictById(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}
