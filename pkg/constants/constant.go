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

package constants

const (
	UserTableName                       = "user"
	SecretKey                           = "secret key"
	IdentityKey                         = "id"
	ApiServiceName                      = "douyinapi"
	UserServiceName                     = "douyinuser"
	VideoServiceName                    = "douyinvideo"
	CommentServiceName                  = "douyincomment"
	VideoServiceMySQLDefaultDSN         = "root:no@tcp(43.136.221.7:3306)/video_db?charset=utf8mb4&parseTime=True&loc=Local"
	MySQLDefaultDSN                     = "root:no@tcp(43.136.221.7:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                         = "43.136.22.7:2379"
	CPURateLimit                float64 = 80.0
)
