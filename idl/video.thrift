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

namespace go douyinvideo

struct BaseResp {
    1:i64 status_code //状态码，0-成功，其他值-失败
    2:string status_message //返回状态描述
    3:i64 service_time //服务时长
}

struct User {
    1:i64 id //用户id
    2:string name //用户名称
    3:i64 follow_count //关注总数
    4:i64 follower_count //粉丝总数
    5:bool is_follow //是否关注，true-已关注，false-未关注
}


struct Video {
    1:i64 id //视频id
    2:User author //用户名称
    3:string play_url
    4:string cover_url //粉丝总数
    5:i64 favorite_count
    6:i64 comment_count
    7:bool is_favorite
    8:string title
}



//视频流接口
struct FeedRequest{
    1:string latest_time
    2:string token
}

struct FeedResponse{
    1:i64 status_code //状态码，0-成功，其他值-失败
    2:string status_msg //返回状态描述
    3:i64 next_time
    4:list<Video> video_list
}

//投稿接口
struct PublishRequest{
    1:i64 author_id
    2:string token
    3:string title
    4:string play_url
    5:string cover_url
}

struct PublishResponse{
    1:i64 status_code //状态码，0-成功，其他值-失败
    2:string status_msg //返回状态描述
}

//发布列表接口
struct GetListRequest{
    1:string token
    2:i64 user_id// 用户id
}

struct GetListResponse{
    1:i64 status_code //状态码，0-成功，其他值-失败
    2:string status_msg //返回状态描述
    3:list<Video> video_list
}
//更新点赞
struct AddLikeRequest{
    1:i64 video_id// 用户id
    2:string action_type
}

struct AddLikeRespose{
    1:i64 status_code //状态码，0-成功，其他值-失败
    2:string status_msg //返回状态描述
}

//根据id获取视频列表
struct GetListByVIDRequest{
    1:list<i64> video_id// 用户id
}
struct GetListByVIDRespose{
    1:i64 status_code //状态码，0-成功，其他值-失败
    2:string status_msg //返回状态描述
    3:list<Video> video_list
}

service VideoServer {
    PublishResponse Publish(1:PublishRequest req)
    GetListResponse GetList(1:GetListRequest req)
    FeedResponse Feed(1:FeedRequest req)
}
