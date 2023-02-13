// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"context"
	"douyin-user/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id            int64  `json:"id"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	VideoCount    int64  `json:"video_count"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// QueryUser query user by name
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetUser get userInfo by id
func GetUser(ctx context.Context, userId int64) (*User, error) {
	res := new(User)
	// WithContext change current instance db's context to ctx
	if err := DB.WithContext(ctx).Where("id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetUsers multiple get list of user info
//func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
//	res := make([]*User, 0)
//	if len(userIDs) == 0 {
//		return res, nil
//	}
//
//	// WithContext change current instance db's context to ctx
//	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
//		return nil, err
//	}
//	return res, nil
//}
