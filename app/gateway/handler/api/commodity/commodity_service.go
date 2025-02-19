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

// Code generated by hertz generator.

package api

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/west2-online/DomTok/app/gateway/pack"
	"github.com/west2-online/DomTok/app/gateway/rpc"
	"github.com/west2-online/DomTok/kitex_gen/commodity"
	"github.com/west2-online/DomTok/pkg/errno"
	"github.com/west2-online/DomTok/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	api "github.com/west2-online/DomTok/app/gateway/model/api/commodity"
)

// CreateCoupon .
// @router /api/commodity/coupon/create [POST]
func CreateCoupon(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateCouponReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CreateCouponResp)

	c.JSON(consts.StatusOK, resp)
}

// DeleteCoupon .
// @router /api/commodity/coupon/delete [DELETE]
func DeleteCoupon(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteCouponReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DeleteCouponResp)

	c.JSON(consts.StatusOK, resp)
}

// CreateUserCoupon .
// @router /api/commodity/coupon/receive [POST]
func CreateUserCoupon(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateUserCouponReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CreateUserCouponResp)

	c.JSON(consts.StatusOK, resp)
}

// ViewCoupon .
// @router /api/commodity/coupon/search [GET]
func ViewCoupon(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ViewCouponReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ViewCouponResp)

	c.JSON(consts.StatusOK, resp)
}

// ViewUserAllCoupon .
// @router /api/commodity/coupon/all [GET]
func ViewUserAllCoupon(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ViewUserAllCouponReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ViewUserAllCouponResp)

	c.JSON(consts.StatusOK, resp)
}

// UseUserCoupon .
// @router /api/commodity/coupon/use [POST]
func UseUserCoupon(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UseUserCouponReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UseUserCouponResp)

	c.JSON(consts.StatusOK, resp)
}

// CreateSpu .
// @router /api/commodity/spu/create [POST]
func CreateSpu(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateSpuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	resp := new(api.CreateSpuResp)

	file, err := c.FormFile("goodsHeadDrawing")
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	_, ok := utils.CheckImageFileType(file)
	if !ok {
		pack.RespError(c, errno.ParamVerifyError)
		return
	}

	datas, err := utils.FileToBytes(file)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	res, err := rpc.CreateSpuRPC(ctx, &commodity.CreateSpuReq{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		Price:       req.Price,
		ForSale:     req.ForSale,
		Shipping:    req.Shipping,
		BufferCount: int64(len(datas)),
	}, datas)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	resp.SpuID = res
	pack.RespData(c, resp)
}

// UpdateSpu .
// @router /api/commodity/spu/update [POST]
func UpdateSpu(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateSpuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}
	var l int64
	var datas [][]byte

	file, err := c.FormFile("goodsHeadDrawing")
	if err != nil {
		if !errors.Is(err, protocol.ErrMissingFile) {
			pack.RespError(c, errno.ParamVerifyError.WithError(err))
			return
		}
	} else {
		_, ok := utils.CheckImageFileType(file)
		if !ok {
			pack.RespError(c, errno.ParamVerifyError)
			return
		}
		datas, err = utils.FileToBytes(file)
		if err != nil {
			pack.RespError(c, err)
			return
		}
		l = int64(len(datas))
	}

	err = rpc.UpdateSpuRPC(ctx, &commodity.UpdateSpuReq{
		SpuID:       req.SpuID,
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		Price:       req.Price,
		ForSale:     req.ForSale,
		Shipping:    req.Shipping,
		BufferCount: &l,
	}, datas)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	pack.RespSuccess(c)
}

// ViewSpu .
// @router /api/commodity/spu/search [GET]
func ViewSpu(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ViewSpuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ViewSpuResp)

	c.JSON(consts.StatusOK, resp)
}

// DeleteSpu .
// @router /api/commodity/spu/delete [DELETE]
func DeleteSpu(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteSpuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithMessage(err.Error()))
		return
	}

	err = rpc.DeleteSpuRPC(ctx, &commodity.DeleteSpuReq{
		SpuID: req.SpuID,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}

	pack.RespSuccess(c)
}

// ViewSpuImage .
// @router /api/commodity/spu/image [GET]
func ViewSpuImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ViewSpuImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ViewSpuImageResp)

	c.JSON(consts.StatusOK, resp)
}

// CreateSku .
// @router /api/commodity/sku/create [POST]
func CreateSku(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateSkuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CreateSkuResp)

	c.JSON(consts.StatusOK, resp)
}

// UpdateSku .
// @router /api/commodity/sku/upadte [POST]
func UpdateSku(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateSkuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UpdateSkuResp)

	c.JSON(consts.StatusOK, resp)
}

// DeleteSku .
// @router /api/commodity/sku/delete [DELETE]
func DeleteSku(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteSkuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DeleteSkuResp)

	c.JSON(consts.StatusOK, resp)
}

// ViewSkuImage .
// @router /api/commodity/sku/image [GET]
func ViewSkuImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ViewSkuImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ViewSkuImageResp)

	c.JSON(consts.StatusOK, resp)
}

// ViewSku .
// @router /api/commodity/sku/search [GET]
func ViewSku(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ViewSkuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ViewSkuResp)

	c.JSON(consts.StatusOK, resp)
}

// UploadSkuAttr .
// @router /api/commodity/sku/attr [POST]
func UploadSkuAttr(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UploadSkuAttrReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UploadSkuAttrResp)

	c.JSON(consts.StatusOK, resp)
}

// ListSkuInfo .
// @router /api/commodity/sku/list [GET]
func ListSkuInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ListSkuInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ListSkuInfoResp)

	c.JSON(consts.StatusOK, resp)
}

// ViewHistory .
// @router /api/commodity/price/history [GET]
func ViewHistory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ViewHistoryPriceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ViewHistoryPriceResp)

	c.JSON(consts.StatusOK, resp)
}

// CreateCategory .
// @router /api/commodity/category/create [POST]
func CreateCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CreateCategoryResp)

	c.JSON(consts.StatusOK, resp)
}

// DeleteCategory .
// @router /api/commodity/category/delete [DELETE]
func DeleteCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DeleteCategoryResp)

	c.JSON(consts.StatusOK, resp)
}

// ViewCategory .
// @router /api/commodity/category/search [GET]
func ViewCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ViewCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ViewCategoryResp)

	c.JSON(consts.StatusOK, resp)
}

// UpdateCategory .
// @router /api/commodity/category/update [POST]
func UpdateCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UpdateCategoryResp)

	c.JSON(consts.StatusOK, resp)
}

// CreateSpuImage .
// @router /api/v1/commodity/spu/image/create [POST]
func CreateSpuImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateSpuImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}
	resp := new(api.CreateSpuImageResp)

	file, err := c.FormFile("spuImage")
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	_, ok := utils.CheckImageFileType(file)
	if !ok {
		pack.RespError(c, errno.ParamVerifyError)
		return
	}

	datas, err := utils.FileToBytes(file)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	id, err := rpc.CreateSpuImageRPC(ctx, &commodity.CreateSpuImageReq{
		BufferCount: int64(len(datas)),
		SpuID:       req.SpuID,
	}, datas)
	if err != nil {
		pack.RespError(c, err)
		return
	}
	resp.ImageID = id
	pack.RespData(c, resp)
}

// UpdateSpuImage .
// @router /api/v1/commodity/spu/image/update [POST]
func UpdateSpuImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateSpuImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	file, err := c.FormFile("spuImage")
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	_, ok := utils.CheckImageFileType(file)
	if !ok {
		pack.RespError(c, errno.ParamVerifyError)
		return
	}

	datas, err := utils.FileToBytes(file)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	err = rpc.UpdateSpuImageRPC(ctx, &commodity.UpdateSpuImageReq{
		BufferCount: int64(len(datas)),
		ImageID:     req.ImageID,
	}, datas)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	pack.RespSuccess(c)
}

// DeleteSpuImage .
// @router /api/v1/commodity/spu/image/delete [DELETE]
func DeleteSpuImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteSpuImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	err = rpc.DeleteSpuImageRPC(ctx, &commodity.DeleteSpuImageReq{
		req.SpuImageID,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}

	pack.RespSuccess(c)
}
