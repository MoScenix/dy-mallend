package product

import (
	"context"
	product "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResp, err error) {
	resp, err = defaultClient.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (resp *product.SearchProductsResp, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddProduct(ctx context.Context, req *product.AddProductReq, callOptions ...callopt.Option) (resp *product.AddProductResp, err error) {
	resp, err = defaultClient.AddProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProduct(ctx context.Context, req *product.DeleteProductReq, callOptions ...callopt.Option) (resp *product.DeleteProductResp, err error) {
	resp, err = defaultClient.DeleteProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProDucts(ctx context.Context, req *product.GetProDuctsReq, callOptions ...callopt.Option) (resp *product.GetProDuctsResp, err error) {
	resp, err = defaultClient.GetProDucts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProDucts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProductsById(ctx context.Context, req *product.GetProductsByIdReq, callOptions ...callopt.Option) (resp *product.GetProductsByIdResp, err error) {
	resp, err = defaultClient.GetProductsById(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProductsById call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
