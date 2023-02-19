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

package video_handler

type FeedParam struct {
	LatestTime string `query:"latest_time" json:"latest_time"`
	Token      string `query:"token" json:"token"`
}

type GetListParam struct {
	Token  string `query:"token" json:"token"`
	UserId int64  `query:"user_id" json:"user_id"`
}

type PublishParam struct {
	//Data  *multipart.FileHeader `query:"data" json:"data"`
	Data  []byte `query:"data" json:"data"`
	Token string `query:"token" json:"token"`
	Title string `query:"title" json:"title"`
}
