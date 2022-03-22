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
	buildFieldNames          = builder.RawFieldNames(&Build{})
	buildRows                = strings.Join(buildFieldNames, ",")
	buildRowsExpectAutoSet   = strings.Join(stringx.Remove(buildFieldNames, "`create_time`", "`update_time`"), ",")
	buildRowsWithPlaceHolder = strings.Join(stringx.Remove(buildFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBuildIdPrefix = "cache:build:id:"
)

type (
	BuildModel interface {
		Insert(data *Build) (sql.Result, error)
		FindOne(id string) (*Build, error)
		Update(data *Build) error
		Delete(id string) error
	}

	defaultBuildModel struct {
		sqlc.CachedConn
		table string
	}

	Build struct {
		Id             string         `db:"id"`
		CreateTime     time.Time      `db:"create_time"`
		UpdateTime     time.Time      `db:"update_time"`
		Name           string         `db:"name"`
		ProjectId      string         `db:"project_id"`
		BuildingDetail string         `db:"building_detail"`
		Desc           sql.NullString `db:"desc"`
	}
)

func NewBuildModel(conn sqlx.SqlConn, c cache.CacheConf) BuildModel {
	return &defaultBuildModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`build`",
	}
}

func (m *defaultBuildModel) Insert(data *Build) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, buildRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.Name, data.ProjectId, data.BuildingDetail, data.Desc)

	return ret, err
}

func (m *defaultBuildModel) FindOne(id string) (*Build, error) {
	buildIdKey := fmt.Sprintf("%s%v", cacheBuildIdPrefix, id)
	var resp Build
	err := m.QueryRow(&resp, buildIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", buildRows, m.table)
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

func (m *defaultBuildModel) Update(data *Build) error {
	buildIdKey := fmt.Sprintf("%s%v", cacheBuildIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, buildRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.ProjectId, data.BuildingDetail, data.Desc, data.Id)
	}, buildIdKey)
	return err
}

func (m *defaultBuildModel) Delete(id string) error {

	buildIdKey := fmt.Sprintf("%s%v", cacheBuildIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, buildIdKey)
	return err
}

func (m *defaultBuildModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBuildIdPrefix, primary)
}

func (m *defaultBuildModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", buildRows, m.table)
	return conn.QueryRow(v, query, primary)
}
