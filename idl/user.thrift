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

namespace go douyinuser

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
//    5:bool is_follow //是否关注，true-已关注，false-未关注
}

struct UserRequest{
    1:string username // 注册用户名，最长32个字符
    2:string password // 密码，最长32个字符
}

struct UserResponse{
    1:i64 user_id
    2:string token
}

//创建用户
struct CreateUserRequest{
    1:UserRequest userReq
}

struct CreateUserResponse{
    1:UserResponse userResp
    2:BaseResp baseResp
}

//
struct CheckUserRequest{
    1:UserRequest userReq
}

struct CheckUserResponse{
    1:UserResponse userResp
    2:BaseResp baseResp
}

struct GetUserInfoRequest{
    1:i64 user_id// 用户id
}

struct GetUserInfoResponse{
    1:User user// 用户信息
    2:BaseResp baseResp
}



service UserServer {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    CheckUserResponse CheckUser(1:CheckUserRequest req)
    GetUserInfoResponse GetUserInfo(1:GetUserInfoRequest req)
}
