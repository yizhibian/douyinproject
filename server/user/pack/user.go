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
	"douyin-user/idl/douyin_user/kitex_gen/douyinuser"
	"douyin-user/server/user/dal/db"
)

// UserInfo pack user detail info
func UserInfo(u *db.User) *douyinuser.User {
	if u == nil {
		return nil
	}
	return &douyinuser.User{
		Id:            u.Id,
		Name:          u.UserName,
		FollowCount:   u.FollowerCount,
		FollowerCount: u.FollowerCount}
}

// Users pack list of user info
//func Users(us []*db.User) []*douyinfavorite.User {
//	users := make([]*douyinfavorite.User, 0)
//	for _, u := range us {
//		if user2 := User(u); user2 != nil {
//			users = append(users, user2)
//		}
//	}
//	return users
//}
