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

// Code generated by Kitex v0.12.0. DO NOT EDIT.

package cartservice

import (
	"context"

	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"

	cart "github.com/west2-online/DomTok/kitex_gen/cart"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddGoodsIntoCart(ctx context.Context, req *cart.AddGoodsIntoCartRequest, callOptions ...callopt.Option) (r *cart.AddGoodsIntoCartResponse, err error)
	ShowCartGoodsList(ctx context.Context, req *cart.ShowCartGoodsListRequest, callOptions ...callopt.Option) (r *cart.ShowCartGoodsListResponse, err error)
	UpdateCartGoods(ctx context.Context, req *cart.UpdateCartGoodsRequest, callOptions ...callopt.Option) (r *cart.UpdateCartGoodsResponse, err error)
	DeleteCartGoods(ctx context.Context, req *cart.DeleteAllCartGoodsRequest, callOptions ...callopt.Option) (r *cart.DeleteAllCartGoodsResponse, err error)
	DeleteAllCartGoods(ctx context.Context, req *cart.DeleteAllCartGoodsRequest, callOptions ...callopt.Option) (r *cart.DeleteAllCartGoodsResponse, err error)
	PayCartGoods(ctx context.Context, req *cart.PayCartGoodsRequest, callOptions ...callopt.Option) (r *cart.PayCartGoodsResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kCartServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCartServiceClient struct {
	*kClient
}

func (p *kCartServiceClient) AddGoodsIntoCart(ctx context.Context, req *cart.AddGoodsIntoCartRequest, callOptions ...callopt.Option) (r *cart.AddGoodsIntoCartResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddGoodsIntoCart(ctx, req)
}

func (p *kCartServiceClient) ShowCartGoodsList(ctx context.Context, req *cart.ShowCartGoodsListRequest, callOptions ...callopt.Option) (r *cart.ShowCartGoodsListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ShowCartGoodsList(ctx, req)
}

func (p *kCartServiceClient) UpdateCartGoods(ctx context.Context, req *cart.UpdateCartGoodsRequest, callOptions ...callopt.Option) (r *cart.UpdateCartGoodsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCartGoods(ctx, req)
}

func (p *kCartServiceClient) DeleteCartGoods(ctx context.Context, req *cart.DeleteAllCartGoodsRequest, callOptions ...callopt.Option) (r *cart.DeleteAllCartGoodsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCartGoods(ctx, req)
}

func (p *kCartServiceClient) DeleteAllCartGoods(ctx context.Context, req *cart.DeleteAllCartGoodsRequest, callOptions ...callopt.Option) (r *cart.DeleteAllCartGoodsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteAllCartGoods(ctx, req)
}

func (p *kCartServiceClient) PayCartGoods(ctx context.Context, req *cart.PayCartGoodsRequest, callOptions ...callopt.Option) (r *cart.PayCartGoodsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PayCartGoods(ctx, req)
}
