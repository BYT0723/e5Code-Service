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
	projectFieldNames          = builder.RawFieldNames(&Project{})
	projectRows                = strings.Join(projectFieldNames, ",")
	projectRowsExpectAutoSet   = strings.Join(stringx.Remove(projectFieldNames, "`create_time`", "`update_time`"), ",")
	projectRowsWithPlaceHolder = strings.Join(stringx.Remove(projectFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheProjectIdPrefix = "cache:project:id:"
)

type (
	ProjectModel interface {
		Insert(data Project) (sql.Result, error)
		FindOne(id string) (*Project, error)
		Update(data Project) error
		Delete(id string) error
	}

	defaultProjectModel struct {
		sqlc.CachedConn
		table string
	}

	Project struct {
		Id         string         `db:"id"`
		CreateTime time.Time      `db:"create_time"`
		UpdateTime time.Time      `db:"update_time"`
		Name       string         `db:"name"`
		Desc       sql.NullString `db:"desc"`
		Url        string         `db:"url"`
		OwnerId    string         `db:"owner_id"`
	}
)

func NewProjectModel(conn sqlx.SqlConn, c cache.CacheConf) ProjectModel {
	return &defaultProjectModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`project`",
	}
}

func (m *defaultProjectModel) Insert(data Project) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, projectRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.Name, data.Desc, data.Url, data.OwnerId)

	return ret, err
}

func (m *defaultProjectModel) FindOne(id string) (*Project, error) {
	projectIdKey := fmt.Sprintf("%s%v", cacheProjectIdPrefix, id)
	var resp Project
	err := m.QueryRow(&resp, projectIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", projectRows, m.table)
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

func (m *defaultProjectModel) Update(data Project) error {
	projectIdKey := fmt.Sprintf("%s%v", cacheProjectIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, projectRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Desc, data.Url, data.OwnerId, data.Id)
	}, projectIdKey)
	return err
}

func (m *defaultProjectModel) Delete(id string) error {

	projectIdKey := fmt.Sprintf("%s%v", cacheProjectIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, projectIdKey)
	return err
}

func (m *defaultProjectModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheProjectIdPrefix, primary)
}

func (m *defaultProjectModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", projectRows, m.table)
	return conn.QueryRow(v, query, primary)
}
