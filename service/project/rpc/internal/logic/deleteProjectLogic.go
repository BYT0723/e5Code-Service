package logic

import (
	"context"
	"strings"

	"e5Code-Service/common/contextx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DeleteProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProjectLogic {
	return &DeleteProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProjectLogic) DeleteProject(in *project.DeleteProjectReq) (*project.DeleteProjectRsp, error) {
	// 从metadata获取UserID
	uID, err := contextx.GetValueFromMetadata(l.ctx, contextx.UserID)
	if err != nil {
		logx.Error("Fail to getUserID: ", err.Error())
		return nil, status.Error(codesx.ContextError, err.Error())
	}

	// 获取User
	us, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Id: uID})
	if err != nil {
		logx.Error("Fail to getUser on DeleteProject: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	// 判断project是否存在
	p := &model.Project{}
	if err := l.svcCtx.DB.Where("id = ?", in.Id).First(p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logx.Error("Fail to GetProject on DeleteProject: ", err.Error())
			return nil, status.Error(codesx.NotFound, "ProjectNotFound")
		}
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	// 销毁仓库
	if strings.Contains(p.Url, l.svcCtx.Config.GitRegistryUrl.Http) || strings.Contains(p.Url, l.svcCtx.Config.GitRegistryUrl.SSH) {
		if res, err := l.svcCtx.GitCli.DestoryRegistry(us.Name, p.Name); err != nil {
			logx.Error("Fail to DestoryRegistry on deleteProject: ", err.Error())
			return nil, status.Error(codesx.GitError, res)
		}
	}

	// 删除project
	if err := l.svcCtx.DB.Delete(&model.Project{ID: in.Id}).Error; err != nil {
		logx.Error("Fail to delete Project, err: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &project.DeleteProjectRsp{}, nil
}
