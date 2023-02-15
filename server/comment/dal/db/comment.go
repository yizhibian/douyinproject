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

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	ID      uint   `gorm:"primaryKey"`
	VideoId int64  `json:"video_id"`
	Content string `json:"content"`
	UserId  int64  `json:"user_id"`
	Status  int64  `json:"status"`
	//CreatedAt string `json:"created_at"`
	//UpdatedAt string `json:"updated_at"`
	//DeletedAt string `json:"deleted_at"`
}

func (u *Comment) CommentName() string {
	return "comment"
}

// MGetComments  multiple get list of Comment info
func MGetComments(ctx context.Context, VideoId []int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if len(VideoId) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("video_id in ?", VideoId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateComment(ctx context.Context, comment []*Comment) error {
	return DB.WithContext(ctx).Create(comment).Error
}

func QueryComment(ctx context.Context, CommentID int64) (*Comment, error) {
	var total int64
	var res *Comment
	conn := DB.WithContext(ctx).Model(&Comment{}).Where("id = ?", CommentID).Find(&res)

	if err := conn.Count(&total).Error; err != nil {
		return res, err
	}
	return res, nil
}
func QueryComments(ctx context.Context, VideoID int64) ([]*Comment, error) {
	var total int64
	var res []*Comment
	conn := DB.WithContext(ctx).Model(&Comment{}).Where("video_id = ?", VideoID).Find(&res)
	if err := conn.Count(&total).Error; err != nil {
		return res, err
	}
	return res, nil
}

// DeleteComment delete note info
func DeleteComment(ctx context.Context, CommentId, userID, VideoID int64) error {
	//params := map[string]interface{}{}
	//	params["status"] = 0
	//	return DB.WithContext(ctx).Model(&Comment{}).Where("id = ? and user_id = ? ", CommentId, userID).Updates(params).Error
	return DB.WithContext(ctx).Where("id = ? and user_id = ? and video_id=? ", CommentId, userID, VideoID).Delete(&Comment{}).Error
}

//// QueryUser query list of user info
//func QueryComment(ctx context.Context, userName string) ([]*User, error) {
//	res := make([]*User, 0)
//	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
//		return nil, err
//	}
//	return res, nil
//}
