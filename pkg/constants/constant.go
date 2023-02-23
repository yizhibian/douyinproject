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
	VideoServiceMySQLDefaultDSN         = "root:123456@tcp(43.136.22.7:3306)/video_db?charset=utf8mb4&parseTime=True&loc=Local"
	MySQLDefaultDSN                     = "gorm:gorm@tcp(47.113.179.3:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                         = "47.115.227.234:2379"
	CPURateLimit                float64 = 80.0
)
