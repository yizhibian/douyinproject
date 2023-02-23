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

namespace go douyinfavorite

struct BaseResp {
    1:i64 status_code //状态码，0-成功，其他值-失败
    2:string status_message //返回状态描述
    3:i64 service_time //服务时长
}

struct Favorite {
    1:i64 user_id //用户id
    2:i64 video_id //视频id
    3:i32 action_type
}

struct LikeRequest{
    1:i64 user_id //用户id
    2:i64 video_id //视频id
    3:i32 action_type//点赞或者取消
}

struct LikeResponse{
   1:BaseResp baseResp
}

struct getVideoIdsRequest{
    1:i64 user_id //用户id
}

struct getVideoIdsResponse{
   1:BaseResp baseResp
   2:list<i64> video_list
}




service UserServer {
    LikeResponse like(1:LikeRequest req)
    getVideoIdsResponse getVideoIds(1:getVideoIdsRequest req)
}
