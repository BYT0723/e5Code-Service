package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	deployFieldNames          = builder.RawFieldNames(&Deploy{})
	deployRows                = strings.Join(deployFieldNames, ",")
	deployRowsExpectAutoSet   = strings.Join(stringx.Remove(deployFieldNames, "`create_time`", "`update_time`"), ",")
	deployRowsWithPlaceHolder = strings.Join(stringx.Remove(deployFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheDeployIdPrefix = "cache:deploy:id:"
)

type (
	DeployModel interface {
		Insert(data *Deploy) (sql.Result, error)
		FindOne(id string) (*Deploy, error)
		Update(data *Deploy) error
		Delete(id string) error
	}

	defaultDeployModel struct {
		sqlc.CachedConn
		table string
	}

	Deploy struct {
		Id              string         `db:"id"`
		CreateTime      time.Time      `db:"create_time"`
		UpdateTime      time.Time      `db:"update_time"`
		Name            string         `db:"name"`
		ProjectId       string         `db:"project_id"`
		SshConfig       sql.NullString `db:"sshConfig"`
		ContainerConfig sql.NullString `db:"containerConfig"`
	}
)

func NewDeployModel(conn sqlx.SqlConn, c cache.CacheConf) DeployModel {
	return &defaultDeployModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`deploy`",
	}
}

func (m *defaultDeployModel) Insert(data *Deploy) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, deployRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.Name, data.ProjectId, data.SshConfig, data.ContainerConfig)

	return ret, err
}

func (m *defaultDeployModel) FindOne(id string) (*Deploy, error) {
	deployIdKey := fmt.Sprintf("%s%v", cacheDeployIdPrefix, id)
	var resp Deploy
	err := m.QueryRow(&resp, deployIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", deployRows, m.table)
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

func (m *defaultDeployModel) Update(data *Deploy) error {
	deployIdKey := fmt.Sprintf("%s%v", cacheDeployIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, deployRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.ProjectId, data.SshConfig, data.ContainerConfig, data.Id)
	}, deployIdKey)
	return err
}

func (m *defaultDeployModel) Delete(id string) error {

	deployIdKey := fmt.Sprintf("%s%v", cacheDeployIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, deployIdKey)
	return err
}

func (m *defaultDeployModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDeployIdPrefix, primary)
}

func (m *defaultDeployModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", deployRows, m.table)
	return conn.QueryRow(v, query, primary)
}
