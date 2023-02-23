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

package pack

import (
	"context"
	"douyin-user/idl/douyin_user/kitex_gen/douyinuser"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/server/video/rpc"

	"douyin-user/server/video/dal/db"
	"log"
)

// UserInfo pack user detail info
func VideoInfo(v *db.Video) *douyinvideo.Video {
	log.Printf("VideoInfo:%#v\n", v)
	req := douyinuser.GetUserInfoRequest{
		UserId: v.Id,
	}
	var author douyinvideo.User
	userInfo, err := rpc.GetUserInfo(context.Background(), &req)
	if err != nil {

	}
	author.Name = userInfo.Name
	author.Id = userInfo.Id
	author.FollowCount = author.FollowCount
	author.FollowerCount = author.FollowerCount
	author.IsFollow = author.IsFollow
	return &douyinvideo.Video{
		Id:            v.Id,
		Author:        &author,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}

func VideoInfos(v []*db.Video) []*douyinvideo.Video {
	l := len(v)
	if l == 0 {
		return nil
	}
	vs := make([]*douyinvideo.Video, l, l)
	for i := range v {
		vs[i] = VideoInfo(v[i])
	}
	log.Printf("VideoInfos:%#v\n", vs)
	return vs
}
