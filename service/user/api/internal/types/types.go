// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReply struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type RegisterUserReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterUserReply struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UpdateUserReq struct {
	Id       string `json:"id"`
	Name     string `json:"name,optional"`
	Password string `json:"password,optional"`
}

type UpdateUserReply struct {
	Result bool `json:"result"`
}

type DeleteUserReq struct {
	Id string `json:"id"`
}

type DeleteUserReply struct {
	Result bool `json:"result"`
}

type UserInfoReq struct {
	Email string `json:"email"`
}

type UserInfoReply struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"Name"`
}
