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
)

type Favorite struct {
	//gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (u *Favorite) TableName() string {
	return constants.FavoriteTableName
}

// CreateFavorite add a like
func CreateFavorite(ctx context.Context, favorite *Favorite) error {

	//if err := DB.WithContext(ctx).Where("userId = ? and video_id = ?", favorite.userId, favorite.VideoId).Error; err != nil {
	//	return err
	//}
	return DB.WithContext(ctx).Create(favorite).Error
}

// DelFavorite delete a like
func DelFavorite(ctx context.Context, favorite *Favorite) error {

	//if err := DB.WithContext(ctx).Where("userId = ? and video_id = ?", favorite.userId, favorite.VideoId).Error; err != nil {
	//	return err
	//}
	return DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", favorite.UserId, favorite.VideoId).Delete(favorite).Error
}

func GetFavorite(ctx context.Context, favorite *Favorite) (*Favorite, error) {
	res := new(Favorite)
	// WithContext change current instance db's context to ctx
	if err := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", favorite.UserId, favorite.VideoId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetFavorite multiple get list of video
func MGetFavorite(ctx context.Context, userId int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)

	// WithContext change current instance db's context to ctx
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetUser get userInfo by id
//func GetUser(ctx context.Context, userId int64) (*Favorite, error) {
//	res := new(Favorite)
//	// WithContext change current instance db's context to ctx
//	if err := DB.WithContext(ctx).Where("id = ?", userId).Find(&res).Error; err != nil {
//		return nil, err
//	}
//	return res, nil
//}
