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

// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "github.com/west2-online/DomTok/gateway/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_commodity := _api.Group("/commodity", _commodityMw()...)
			{
				_category := _commodity.Group("/category", _categoryMw()...)
				_category.POST("/create", append(_createcategoryMw(), api.CreateCategory)...)
				_category.DELETE("/delete", append(_deletecategoryMw(), api.DeleteCategory)...)
				_category.GET("/search", append(_viewcategoryMw(), api.ViewCategory)...)
				_category.POST("/update", append(_updatecategoryMw(), api.UpdateCategory)...)
			}
			{
				_coupon := _commodity.Group("/coupon", _couponMw()...)
				_coupon.GET("/all", append(_viewuserallcouponMw(), api.ViewUserAllCoupon)...)
				_coupon.POST("/create", append(_createcouponMw(), api.CreateCoupon)...)
				_coupon.DELETE("/delete", append(_deletecouponMw(), api.DeleteCoupon)...)
				_coupon.POST("/receive", append(_createusercouponMw(), api.CreateUserCoupon)...)
				_coupon.GET("/search", append(_viewcouponMw(), api.ViewCoupon)...)
				_coupon.POST("/use", append(_useusercouponMw(), api.UseUserCoupon)...)
			}
			{
				_price := _commodity.Group("/price", _priceMw()...)
				_price.GET("/history", append(_viewhistoryMw(), api.ViewHistory)...)
			}
			{
				_sku := _commodity.Group("/sku", _skuMw()...)
				_sku.POST("/attr", append(_uploadskuattrMw(), api.UploadSkuAttr)...)
				_sku.POST("/create", append(_createskuMw(), api.CreateSku)...)
				_sku.DELETE("/delete", append(_deleteskuMw(), api.DeleteSku)...)
				_sku.GET("/image", append(_viewskuimageMw(), api.ViewSkuImage)...)
				_sku.GET("/list", append(_listskuinfoMw(), api.ListSkuInfo)...)
				_sku.GET("/search", append(_viewskuMw(), api.ViewSku)...)
				_sku.POST("/upadte", append(_updateskuMw(), api.UpdateSku)...)
			}
			{
				_spu := _commodity.Group("/spu", _spuMw()...)
				_spu.POST("/create", append(_createspuMw(), api.CreateSpu)...)
				_spu.DELETE("/delete", append(_deletespuMw(), api.DeleteSpu)...)
				_spu.GET("/image", append(_viewspuimageMw(), api.ViewSpuImage)...)
				_spu.GET("/search", append(_viewspuMw(), api.ViewSpu)...)
				_spu.POST("/update", append(_updatespuMw(), api.UpdateSpu)...)
			}
		}
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			{
				_user := _v1.Group("/user", _userMw()...)
				_user.GET("/register", append(_registerMw(), api.Register)...)
			}
		}
	}
}
