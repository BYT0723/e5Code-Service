package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	containerFieldNames          = builderx.RawFieldNames(&Container{})
	containerRows                = strings.Join(containerFieldNames, ",")
	containerRowsExpectAutoSet   = strings.Join(stringx.Remove(containerFieldNames, "`create_time`", "`update_time`"), ",")
	containerRowsWithPlaceHolder = strings.Join(stringx.Remove(containerFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheContainerIdPrefix = "cache:container:id:"
)

type (
	ContainerModel interface {
		Insert(data Container) (sql.Result, error)
		FindOne(id string) (*Container, error)
		Update(data Container) error
		Delete(id string) error
	}

	defaultContainerModel struct {
		sqlc.CachedConn
		table string
	}

	Container struct {
		Id       string    `db:"id"`
		CreateAt time.Time `db:"create_at"`
		UpdateAt time.Time `db:"update_at"`
		Name     string    `db:"name"`
		Url      string    `db:"url"`
		DeployId string    `db:"deploy_id"`
		Status   string    `db:"status"`
	}
)

func NewContainerModel(conn sqlx.SqlConn, c cache.CacheConf) ContainerModel {
	return &defaultContainerModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`container`",
	}
}

func (m *defaultContainerModel) Insert(data Container) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, containerRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.CreateAt, data.UpdateAt, data.Name, data.Url, data.DeployId, data.Status)

	return ret, err
}

func (m *defaultContainerModel) FindOne(id string) (*Container, error) {
	containerIdKey := fmt.Sprintf("%s%v", cacheContainerIdPrefix, id)
	var resp Container
	err := m.QueryRow(&resp, containerIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", containerRows, m.table)
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

func (m *defaultContainerModel) Update(data Container) error {
	containerIdKey := fmt.Sprintf("%s%v", cacheContainerIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, containerRowsWithPlaceHolder)
		return conn.Exec(query, data.CreateAt, data.UpdateAt, data.Name, data.Url, data.DeployId, data.Status, data.Id)
	}, containerIdKey)
	return err
}

func (m *defaultContainerModel) Delete(id string) error {

	containerIdKey := fmt.Sprintf("%s%v", cacheContainerIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, containerIdKey)
	return err
}

func (m *defaultContainerModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheContainerIdPrefix, primary)
}

func (m *defaultContainerModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", containerRows, m.table)
	return conn.QueryRow(v, query, primary)
}
