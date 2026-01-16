package http

import (
	"libong/common/context"
	"libong/common/server/http"
	"shoe-manager/app/interface/category/api"
)

func addCategory(ctx *http.Context) error {
	var (
		req = &api.AddCategoryReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.AddCategory(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func updateCategory(ctx *http.Context) error {
	var (
		req = &api.UpdateCategoryReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.UpdateCategory(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func deleteCategory(ctx *http.Context) error {
	var (
		req = &api.DeleteCategoryReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.DeleteCategory(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func searchCategoriesPage(ctx *http.Context) error {
	var (
		req  = &api.SearchCategoriesPageReq{}
		resp *api.SearchCategoriesPageResp
		err  error
	)
	err = ctx.MarshalGet(req)
	if err != nil {
		return err
	}
	if resp, err = svc.SearchCategoriesPage(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}

func categoryById(ctx *http.Context) error {
	var (
		req  = &api.CategoryByIdReq{}
		resp *api.CategoryByIdResp
		err  error
	)
	err = ctx.MarshalGet(req)
	if err != nil {
		return err
	}
	if resp, err = svc.CategoryById(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}
