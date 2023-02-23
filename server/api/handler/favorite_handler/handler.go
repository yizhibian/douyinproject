package favorite_handler

type LikeParam struct {
	VideoId    int64  `query:"video_id" json:"video_id"`
	Token      string `query:"token" json:"token"`
	ActionType int32  `query:"action_type" json:"action_type"`
}
