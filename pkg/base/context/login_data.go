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

package context

import (
	"context"

	"github.com/bytedance/sonic"

	"github.com/west2-online/DomTok/kitex_gen/model"
	"github.com/west2-online/DomTok/pkg/errno"
	"github.com/west2-online/DomTok/pkg/logger"
)

// TODO 放 constants 里
const loginDataKey string = "loginData"

// TODO 两个函数都改掉, 不要传什么 model.LoginData 了, 要么就找一个通用的地方声明一个专门的结构体来放结果, 要么直接返回 uid 得了

// WithLoginData 将LoginData加入到context中，通过metainfo传递到RPC server
func WithLoginData(ctx context.Context, loginData *model.LoginData) context.Context {
	value, err := sonic.MarshalString(*loginData)
	if err != nil {
		logger.Infof("Failed to marshal LoginData: %v", err)
	}
	return newContext(ctx, loginDataKey, value)
}

// GetLoginData 从context中取出LoginData
func GetLoginData(ctx context.Context) (*model.LoginData, error) {
	user, ok := fromContext(ctx, loginDataKey)
	if !ok {
		return nil, errno.NewErrNo(errno.ParamMissingErrorCode, "Failed to get header in context")
	}
	value := new(model.LoginData)
	err := sonic.UnmarshalString(user, value)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "Failed to get header in context when unmarshalling loginData")
	}
	return value, nil
}
