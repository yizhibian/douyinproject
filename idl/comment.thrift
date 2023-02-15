// idl/comment.thrift
namespace go comment

struct User {
     1: i64 id;
     2: string name;
     3: i64 follow_count;
     4: i64 follower_count;
     5: bool is_follow;
}

struct Comment {
     1: i64 id;
     2: User user;
     3: string content;
     4: string create_date;
}

struct CommentRequest {
    1: i64 UserId (api.body="user_id");
    2: i64 VideoId (api.body="video_id");
    3: i64 actionType (api.body="action_type");
    4: string commentText (api.body="comment_text");
    5: i64 commentId (api.body="comment_id");
}

struct CommentResponse {
    1: i64 status_code;
    2: string status_msg;
    3: Comment comment;
}

struct CommentListRequest {
   1: i64 UserId (api.body="user_id");
   2: i64 VideoId (api.body="video_id");
}

struct CommentListResponse {
   1: i64 status_code;
   2: string status_msg;
   3: list<Comment> comment_list;
}

service CommentService {
    CommentResponse action(1: CommentRequest req) (api.post="/comment/action/");
    CommentListResponse list(1: CommentListRequest req) (api.get="/comment/list/");
}
