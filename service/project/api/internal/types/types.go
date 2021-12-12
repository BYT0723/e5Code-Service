// Code generated by goctl. DO NOT EDIT.
package types

type Project struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Url     string `json:"url"`
	OwnerId string `json:"owner_id"`
}

type GetProjectReq struct {
	ID string `json:"id"`
}

type GetProjectReply struct {
	Result Project `json:"result"`
}

type AddProjectReq struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Url  string `json:"url"`
}

type AddProjectReply struct {
	Result Project `json:"result"`
}

type UpdateProjectReq struct {
	Payload Project `json:"payload"`
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

type Deploy struct {
	ID              string          `json:"id"`
	ProjectID       string          `json:"projectID"`
	Name            string          `json:"name"`
	SSHConfig       SSHConfig       `json:"sshConfig"`
	ContainerConfig ContainerConfig `json:"containerConfig"`
}

type SSHConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	SSHType  string `json:"sshType"`
	Password string `json:"password"`
	SSHKey   string `json:"sshKey"`
}

type ContainerConfig struct {
	Name         string   `json:"name"`
	NetworkType  string   `json:"networkType"`
	IP           string   `json:"ip"`
	Ports        []string `json:"ports"`
	Environments []string `json:"environments"`
}

type GetDeployReq struct {
	ID string `json:"id"`
}

type GetDeployRsp struct {
	Result Deploy `json:"result"`
}

type AddDeployReq struct {
	ProjectID       string          `json:"projectID"`
	Name            string          `json:"name"`
	SSHConfig       SSHConfig       `json:"sshConfig"`
	ContainerConfig ContainerConfig `json:"containerConfig"`
}

type AddDeployRsp struct {
	Result Deploy `json:"result"`
}

type UpdateDeployReq struct {
	Payload Deploy `json:"payload"`
}

type UpdateDeployRsp struct {
	Result bool `json:"result"`
}

type DeleteDeployReq struct {
	ID string `json:"id"`
}

type DeleteDeployRsp struct {
	Result bool `json:"result"`
}

type DeployReq struct {
	ID string `json:"id"`
}

type DeployRsp struct {
	Result bool `json:"result"`
}
