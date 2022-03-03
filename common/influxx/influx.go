package influxx

import (
	"fmt"
	"time"

	influx "github.com/influxdata/influxdb1-client/v2"
)

type InfluxClient struct {
	influx.Client
	Database string
}

type InfluxConnConfig struct {
	Host string
	User string
	Pass string
	DB   string
}

func NewInfluxClient(c InfluxConnConfig) (*InfluxClient, error) {
	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     c.Host,
		Username: c.User,
		Password: c.Pass,
	})
	if err != nil {
		return nil, err
	}
	return &InfluxClient{client, c.DB}, nil
}

func (c *InfluxClient) Query(cmd string) ([]influx.Result, error) {
	rsp, err := c.Client.Query(influx.Query{
		Command:  cmd,
		Database: c.Database,
	})
	if err != nil {
		return nil, err
	}
	return rsp.Results, nil
}

func (c *InfluxClient) Insert(measurement string, tags map[string]string, fields map[string]interface{}) error {
	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database: c.Database,
	})
	if err != nil {
		return err
	}
	pt, err := influx.NewPoint(measurement, tags, fields, time.Now())
	if err != nil {
		return err
	}
	fmt.Printf("pt: %v\n", pt)
	bp.AddPoint(pt)
	if err := c.Write(bp); err != nil {
		return err
	}
	return nil
}
