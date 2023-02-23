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

package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                = 0
	ServiceErrCode             = 1
	ParamErrCode               = 1
	UserAlreadyExistErrCode    = 1
	AuthorizationFailedErrCode = 1
	NilValueErrCode            = 1
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")

	UserIdErr             = NewErrNo(ParamErrCode, "Wrong Parameter[UserId] has been given")
	ActionTypeErr         = NewErrNo(ParamErrCode, "Wrong Parameter[ActionType] has been given")
	CommentTextErr        = NewErrNo(ParamErrCode, "Wrong Parameter[CommentText] has been given")
	VideoErr              = NewErrNo(ParamErrCode, "Wrong Parameter[Video] has been given")
	CommentTextTooLongErr = NewErrNo(ParamErrCode, "Parameter[CommentText] too long")
	CommentIdErr          = NewErrNo(ParamErrCode, "Wrong Parameter[CommentIdErr] has been given")

	NoFavoriteErr      = NewErrNo(ParamErrCode, "Did not favorite this video before")
	FavoriteErr        = NewErrNo(ParamErrCode, "You have favorited this video before")
	WrongActionTypeErr = NewErrNo(ParamErrCode, "the action_type was wrong")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
