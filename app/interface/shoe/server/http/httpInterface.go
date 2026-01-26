package http

import (
	"libong/common/context"
	"libong/common/server/http"
	"shoe-manager/app/interface/shoe/api"
)

func addShoe(ctx *http.Context) error {
	var (
		req = &api.AddShoeReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.AddShoe(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func updateShoe(ctx *http.Context) error {
	var (
		req = &api.UpdateShoeReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.UpdateShoe(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func deleteShoe(ctx *http.Context) error {
	var (
		req = &api.DeleteShoeReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.DeleteShoe(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func searchShoesPage(ctx *http.Context) error {
	var (
		req  = &api.SearchShoesPageReq{}
		resp *api.SearchShoesPageResp
		err  error
	)
	err = ctx.MarshalGet(req)
	if err != nil {
		return err
	}
	if resp, err = svc.SearchShoesPage(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}

func shoeById(ctx *http.Context) error {
	var (
		req  = &api.ShoeByIdReq{}
		resp *api.ShoeByIdResp
		err  error
	)
	err = ctx.MarshalGet(req)
	if err != nil {
		return err
	}
	if resp, err = svc.ShoeById(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}
func updateShoeHot(ctx *http.Context) error {
	var (
		req = &api.UpdateShoeHotReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.UpdateShoeHot(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}
func changeShoeFavour(ctx *http.Context) error {
	var (
		req = &api.ChangeShoeFavourReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.ChangeShoeFavour(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}
func searchSelectList(ctx *http.Context) error {
	var (
		req  = &api.SearchSelectListReq{}
		resp *api.SearchSelectListResp
		err  error
	)
	err = ctx.MarshalGet(req)
	if err != nil {
		return err
	}
	if resp, err = svc.SearchSelectList(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}
