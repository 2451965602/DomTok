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

package usecase

import (
	"context"
	"fmt"
	"github.com/west2-online/DomTok/app/user/domain/model"
	"github.com/west2-online/DomTok/pkg/errno"
)

// Login 用户登录
func (uc *useCase) Login(ctx context.Context, user *model.User) (*model.User, error) {
	user, err := uc.db.GetUserInfo(ctx, user.UserName)
	if err != nil {
		return nil, fmt.Errorf("get user info failed: %w", err)
	}

	if err = uc.svc.CheckPassword(user.Password, user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *useCase) RegisterUser(ctx context.Context, u *model.User) (uid int64, err error) {
	// 这里进行了简单的密码和邮箱格式的校验, 如果后续还需要对别的参数进行校验可以再加
	if err = uc.svc.Verify(uc.svc.VerifyEmail(u.Email), uc.svc.VerifyPassword(u.Password)); err != nil {
		return
	}
	// 判断是否已经注册过
	// 注意: 这里使用 uc 调用了 DB, 但显然这个方法其他地方也可能会用的上, 所以可以考虑包装在 service 里面
	exist, err := uc.db.IsUserExist(ctx, u.UserName)
	if err != nil {
		// 这里返回了 fmt.Errorf 而不是 errno 的原因是 db.IsUserExist 返回的已经是 errno 了
		// 这里是用 %w 占位符做了一层 wrap, 其实这个 error 的底部(origin error) 还是 errno 类型的
		return 0, fmt.Errorf("check user exist failed: %w", err)
	}
	if exist {
		return 0, errno.NewErrNo(errno.ServiceUserExist, "user already exist")
	}

	if u.Password, err = uc.svc.EncryptPassword(u.Password); err != nil {
		return 0, err
	}

	// 这里没有直接调用 db.CreateUser 是因为 svc.CreateUser 包含了一点业务逻辑, 这些细节不需要被 useCase 知道
	if err = uc.svc.CreateUser(ctx, u); err != nil {
		return
	}

	return u.Uid, nil
}
