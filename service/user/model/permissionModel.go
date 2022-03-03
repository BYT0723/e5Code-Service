package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	permissionFieldNames          = builder.RawFieldNames(&Permission{})
	permissionRows                = strings.Join(permissionFieldNames, ",")
	permissionRowsExpectAutoSet   = strings.Join(stringx.Remove(permissionFieldNames, "`create_time`", "`update_time`"), ",")
	permissionRowsWithPlaceHolder = strings.Join(stringx.Remove(permissionFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cachePermissionIdPrefix              = "cache:permission:id:"
	cachePermissionUserIdProjectIdPrefix = "cache:permission:userId:projectId:"
)

type (
	PermissionModel interface {
		Insert(data *Permission) (sql.Result, error)
		FindOne(id string) (*Permission, error)
		FindOneByUserIdProjectId(userId string, projectId string) (*Permission, error)
		Update(data *Permission) error
		Delete(id string) error
	}

	defaultPermissionModel struct {
		sqlc.CachedConn
		table string
	}

	Permission struct {
		Id         string    `db:"id"`
		UserId     string    `db:"user_id"`
		ProjectId  string    `db:"project_id"`
		Permission int64     `db:"permission"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewPermissionModel(conn sqlx.SqlConn, c cache.CacheConf) PermissionModel {
	return &defaultPermissionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`permission`",
	}
}

func (m *defaultPermissionModel) Insert(data *Permission) (sql.Result, error) {
	permissionIdKey := fmt.Sprintf("%s%v", cachePermissionIdPrefix, data.Id)
	permissionUserIdProjectIdKey := fmt.Sprintf("%s%v:%v", cachePermissionUserIdProjectIdPrefix, data.UserId, data.ProjectId)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, permissionRowsExpectAutoSet)
		return conn.Exec(query, data.Id, data.UserId, data.ProjectId, data.Permission)
	}, permissionIdKey, permissionUserIdProjectIdKey)
	return ret, err
}

func (m *defaultPermissionModel) FindOne(id string) (*Permission, error) {
	permissionIdKey := fmt.Sprintf("%s%v", cachePermissionIdPrefix, id)
	var resp Permission
	err := m.QueryRow(&resp, permissionIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", permissionRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPermissionModel) FindOneByUserIdProjectId(userId string, projectId string) (*Permission, error) {
	permissionUserIdProjectIdKey := fmt.Sprintf("%s%v:%v", cachePermissionUserIdProjectIdPrefix, userId, projectId)
	var resp Permission
	err := m.QueryRowIndex(&resp, permissionUserIdProjectIdKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `project_id` = ? limit 1", permissionRows, m.table)
		if err := conn.QueryRow(&resp, query, userId, projectId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPermissionModel) Update(data *Permission) error {
	permissionIdKey := fmt.Sprintf("%s%v", cachePermissionIdPrefix, data.Id)
	permissionUserIdProjectIdKey := fmt.Sprintf("%s%v:%v", cachePermissionUserIdProjectIdPrefix, data.UserId, data.ProjectId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, permissionRowsWithPlaceHolder)
		return conn.Exec(query, data.UserId, data.ProjectId, data.Permission, data.Id)
	}, permissionUserIdProjectIdKey, permissionIdKey)
	return err
}

func (m *defaultPermissionModel) Delete(id string) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	permissionIdKey := fmt.Sprintf("%s%v", cachePermissionIdPrefix, id)
	permissionUserIdProjectIdKey := fmt.Sprintf("%s%v:%v", cachePermissionUserIdProjectIdPrefix, data.UserId, data.ProjectId)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, permissionUserIdProjectIdKey, permissionIdKey)
	return err
}

func (m *defaultPermissionModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePermissionIdPrefix, primary)
}

func (m *defaultPermissionModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", permissionRows, m.table)
	return conn.QueryRow(v, query, primary)
}
