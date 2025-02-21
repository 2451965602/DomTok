/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rpc

import (
	"bytes"
	"context"
	"fmt"

	"github.com/west2-online/DomTok/app/commodity/controllers/rpc/pack"
	"github.com/west2-online/DomTok/app/commodity/domain/model"
	"github.com/west2-online/DomTok/app/commodity/usecase"
	"github.com/west2-online/DomTok/kitex_gen/commodity"
	"github.com/west2-online/DomTok/pkg/base"
	"github.com/west2-online/DomTok/pkg/logger"
)

type CommodityHandler struct {
	useCase usecase.CommodityUseCase
}

func (c CommodityHandler) CreateCoupon(ctx context.Context, req *commodity.CreateCouponReq) (r *commodity.CreateCouponResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) DeleteCoupon(ctx context.Context, req *commodity.DeleteCouponReq) (r *commodity.DeleteCouponResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) CreateUserCoupon(ctx context.Context, req *commodity.CreateCouponReq) (r *commodity.CreateUserCouponResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) ViewCoupon(ctx context.Context, req *commodity.ViewCouponReq) (r *commodity.ViewCouponResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) ViewUserAllCoupon(ctx context.Context, req *commodity.ViewCouponReq) (r *commodity.ViewUserAllCouponResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) UseUserCoupon(ctx context.Context, req *commodity.UseUserCouponReq) (r *commodity.UseUserCouponResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) CreateSpu(streamServer commodity.CommodityService_CreateSpuServer) (err error) {
	resp := new(commodity.CreateSpuResp)

	req, err := streamServer.Recv()
	if err != nil {
		logger.Errorf("rpc.CreateSpu: receive error: %v", err)
		resp.Base = base.BuildBaseResp(err)
		return streamServer.SendAndClose(resp)
	}

	for i := 0; i < int(req.BufferCount); i++ {
		fileData, err := streamServer.Recv()
		if err != nil {
			logger.Errorf("rpc.CreateSpu: receive error: %v", err)
			resp.Base = base.BuildBaseResp(err)
			return streamServer.SendAndClose(resp)
		}
		req.GoodsHeadDrawing = bytes.Join([][]byte{req.GoodsHeadDrawing, fileData.GoodsHeadDrawing}, []byte(""))
	}

	id, err := c.useCase.CreateSpu(streamServer.Context(), &model.Spu{
		Name:             req.Name,
		Description:      req.Description,
		CategoryId:       req.CategoryID,
		Price:            req.Price,
		ForSale:          int(req.ForSale),
		Shipping:         req.Shipping,
		GoodsHeadDrawing: req.GoodsHeadDrawing,
	})
	if err != nil {
		logger.Errorf("rpc.CreateSpu: create spu error: %v", err)
		resp.Base = base.BuildBaseResp(err)
		return streamServer.SendAndClose(resp)
	}

	resp.Base = base.BuildBaseResp(nil)
	resp.SpuID = id
	return streamServer.SendAndClose(resp)
}

func (c CommodityHandler) UpdateSpu(streamServer commodity.CommodityService_UpdateSpuServer) (err error) {
	resp := new(commodity.UpdateSpuResp)

	req, err := streamServer.Recv()
	if err != nil {
		logger.Errorf("rpc.UpdateSpu: receive error: %v", err)
		resp.Base = base.BuildBaseResp(err)
		return streamServer.SendAndClose(resp)
	}

	for i := 0; i < int(*req.BufferCount); i++ {
		fileData, err := streamServer.Recv()
		if err != nil {
			logger.Errorf("rpc.UpdateSpu: receive error: %v", err)
			resp.Base = base.BuildBaseResp(err)
			return streamServer.SendAndClose(resp)
		}
		req.GoodsHeadDrawing = bytes.Join([][]byte{req.GoodsHeadDrawing, fileData.GoodsHeadDrawing}, []byte(""))
	}

	err = c.useCase.UpdateSpu(streamServer.Context(), &model.Spu{
		SpuId:            req.SpuID,
		Name:             req.GetName(),
		Description:      req.GetDescription(),
		CategoryId:       req.GetCategoryID(),
		Price:            req.GetPrice(),
		ForSale:          int(req.GetForSale()),
		Shipping:         req.GetShipping(),
		GoodsHeadDrawing: req.GetGoodsHeadDrawing(),
	})
	if err != nil {
		logger.Errorf("rpc.UpdateSpu: update spu error: %v", err)
		resp.Base = base.BuildBaseResp(err)
		return streamServer.SendAndClose(resp)
	}

	resp.Base = base.BuildBaseResp(nil)
	return streamServer.SendAndClose(resp)
}

func (c CommodityHandler) ViewSpu(ctx context.Context, req *commodity.ViewSpuReq) (r *commodity.ViewSpuResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) DeleteSpu(ctx context.Context, req *commodity.DeleteSpuReq) (r *commodity.DeleteSpuResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) ViewSpuImage(ctx context.Context, req *commodity.ViewSpuImageReq) (r *commodity.ViewSpuImageResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) CreateSku(streamServer commodity.CommodityService_CreateSkuServer) (err error) {
	r := new(commodity.CreateSkuResp)

	req, err := streamServer.Recv()
	if err != nil {
		logger.Errorf("rpc.CreateSpu: receive error: %v", err)
		r.Base = base.BuildBaseResp(err)
		return streamServer.SendAndClose(r)
	}

	for i := 0; i < int(req.BufferCount); i++ {
		fileData, err := streamServer.Recv()
		if err != nil {
			logger.Errorf("rpc.CreateSku: receive error: %v", err)
			r.Base = base.BuildBaseResp(err)
			return streamServer.SendAndClose(r)
		}
		req.StyleHeadDrawing = bytes.Join([][]byte{req.StyleHeadDrawing, fileData.StyleHeadDrawing}, []byte(""))
	}

	id, err := c.useCase.CreateSku(streamServer.Context(), &model.Sku{
		Name:             req.Name,
		Stock:            req.Stock,
		Description:      req.Description,
		StyleHeadDrawing: req.StyleHeadDrawing,
		Price:            req.Price,
		ForSale:          int(req.ForSale),
		SpuID:            req.SpuID,
	})
	if err != nil {
		logger.Errorf("rpc.CreateSku: create sku error: %v", err)
		r.Base = base.BuildBaseResp(err)
		return streamServer.SendAndClose(r)
	}

	r.Base = base.BuildBaseResp(nil)
	r.SkuID = id
	return streamServer.SendAndClose(r)
}

func (c CommodityHandler) UpdateSku(streamServer commodity.CommodityService_UpdateSkuServer) (rr error) {
	r := new(commodity.UpdateSkuResp)

	req, err := streamServer.Recv()
	if err != nil {
		logger.Errorf("rpc.UpdateSku: receive error: %v", err)
		r.Base = base.BuildBaseResp(err)
		return streamServer.SendAndClose(r)
	}

	for i := 0; i < int(*req.BufferCount); i++ {
		fileData, err := streamServer.Recv()
		if err != nil {
			logger.Errorf("rpc.UpdateSpu: receive error: %v", err)
			r.Base = base.BuildBaseResp(err)
			return streamServer.SendAndClose(r)
		}
		req.StyleHeadDrawing = bytes.Join([][]byte{req.StyleHeadDrawing, fileData.StyleHeadDrawing}, []byte(""))
	}

	err = c.useCase.UpdateSku(streamServer.Context(), &model.Sku{
		SkuID:            req.SkuID,
		Stock:            req.GetStock(),
		Description:      req.GetDescription(),
		StyleHeadDrawing: req.GetStyleHeadDrawing(),
		Price:            req.GetPrice(),
		ForSale:          int(req.GetForSale()),
	})
	if err != nil {
		logger.Errorf("rpc.UpdateSku: update sku error: %v", err)
		r.Base = base.BuildBaseResp(err)
		return streamServer.SendAndClose(r)
	}

	r.Base = base.BuildBaseResp(nil)
	return streamServer.SendAndClose(r)
}

func (c CommodityHandler) DeleteSku(ctx context.Context, req *commodity.DeleteSkuReq) (r *commodity.DeleteSkuResp, err error) {
	r = new(commodity.DeleteSkuResp)

	sku := &model.Sku{
		SkuID: req.SkuID,
	}

	if err = c.useCase.DeleteSku(ctx, sku); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Base = base.BuildBaseResp(nil)
	return
}

func (c CommodityHandler) ViewSkuImage(ctx context.Context, req *commodity.ViewSkuImageReq) (r *commodity.ViewSkuImageResp, err error) {
	r = new(commodity.ViewSkuImageResp)

	sku := &model.Sku{
		SkuID: req.SkuID,
	}

	var Images []*model.SkuImage

	if Images, err = c.useCase.ViewSkuImage(ctx, sku, req.PageNum, req.PageSize); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Base = base.BuildBaseResp(nil)
	r.Images = pack.BuildImages(Images)
	return
}

func (c CommodityHandler) ViewSku(ctx context.Context, req *commodity.ViewSkuReq) (r *commodity.ViewSkuResp, err error) {
	r = new(commodity.ViewSkuResp)

	var (
		isSpuId bool
		sku     model.Sku
	)

	if req.SkuID != nil {
		sku.SkuID = *req.SkuID
		isSpuId = false
	}
	if req.SpuID != nil {
		sku.SpuID = *req.SpuID
		isSpuId = true
	}
	if req.SkuID == nil && req.SpuID == nil {
		err = fmt.Errorf("ViewSku failed: skuID and spuID are both nil")
		r.Base = base.BuildBaseResp(err)
		return
	}

	Skus, err := c.useCase.ViewSku(ctx, &sku, req.PageNum, req.PageSize, isSpuId)
	if err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Base = base.BuildBaseResp(nil)
	r.Skus = pack.BuildSkus(Skus)
	return
}

func (c CommodityHandler) UploadSkuAttr(ctx context.Context, req *commodity.UploadSkuAttrReq) (r *commodity.UploadSkuAttrResp, err error) {
	r = new(commodity.UploadSkuAttrResp)

	attr := &model.AttrValue{
		SaleAttr:  req.SaleAttr,
		SaleValue: req.SaleValue,
	}

	sku := &model.Sku{
		SkuID: *req.SkuID,
	}

	if err = c.useCase.UploadSkuAttr(ctx, attr, sku); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Base = base.BuildBaseResp(nil)
	return
}

func (c CommodityHandler) ListSkuInfo(ctx context.Context, req *commodity.ListSkuInfoReq) (r *commodity.ListSkuInfoResp, err error) {
	r = new(commodity.ListSkuInfoResp)

	SkuInfos, err := c.useCase.ListSkuInfo(ctx, req.SkuIDs, req.PageNum, req.PageSize)
	if err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Base = base.BuildBaseResp(nil)
	r.SkuInfos = pack.BuildSkuInfos(SkuInfos)
	return
}

func (c CommodityHandler) ViewHistory(ctx context.Context, req *commodity.ViewHistoryPriceReq) (r *commodity.ViewHistoryPriceResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) DescSkuLockStock(ctx context.Context, req *commodity.DescSkuLockStockReq) (r *commodity.DescSkuLockStockResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) IncrSkuLockStock(ctx context.Context, req *commodity.IncrSkuLockStockReq) (r *commodity.IncrSkuLockStockResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) DescSkuStock(ctx context.Context, req *commodity.DescSkuStockReq) (r *commodity.DescSkuStockResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) CreateCategory(ctx context.Context, req *commodity.CreateCategoryReq) (r *commodity.CreateCategoryResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) DeleteCategory(ctx context.Context, req *commodity.DeleteCategoryReq) (r *commodity.DeleteCategoryResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) ViewCategory(ctx context.Context, req *commodity.ViewCategoryReq) (r *commodity.ViewCategoryResp, err error) {
	// TODO implement me
	panic("implement me")
}

func (c CommodityHandler) UpdateCategory(ctx context.Context, req *commodity.UpdateCategoryReq) (r *commodity.UpdateCategoryResp, err error) {
	// TODO implement me
	panic("implement me")
}

func NewCommodityHandler(useCase usecase.CommodityUseCase) *CommodityHandler {
	return &CommodityHandler{useCase}
}
