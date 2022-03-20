// Code generated by goctl. DO NOT EDIT.
package types

type Project struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Url     string `json:"url"`
	OwnerID string `json:"owner_id"`
}

type GetProjectReq struct {
	ID string `json:"id"`
}

type GetProjectReply struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Url     string `json:"url"`
	OwnerID string `json:"owner_id"`
}

type AddProjectReq struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Url  string `json:"url"`
}

type AddProjectReply struct {
	ID string `json:"id"`
}

type UpdateProjectReq struct {
	ID   string `json:"id"`
	Name string `json:"name,optional"`
	Desc string `json:"desc,optional"`
	Url  string `json:"url,optional"`
}

type UpdateProjectReply struct {
	Result bool `json:"result"`
}

type DeleteProjectReq struct {
	ID string `json:"id"`
}

type DeleteProjectReply struct {
	Result bool `json:"result"`
}

type AddUserReq struct {
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
}

type AddUserReply struct {
	Result bool `json:"result"`
}

type RemoveUserReq struct {
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
}

type RemoveUserReply struct {
	Result bool `json:"result"`
}

type ModifyPermissionReq struct {
	UserID       string `json:"user_id"`
	ProjectID    string `json:"project_id"`
	ModifiedType int64  `json:"modified_type"`
	Value        int64  `json:"value"`
}

type ModifyPermissionReply struct {
	Result bool `json:"result"`
}
