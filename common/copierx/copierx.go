package copierx

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ConvertList []copier.TypeConverter

func init() {
	ConvertList = []copier.TypeConverter{
		{
			// time 转 timestamp
			SrcType: time.Time{},
			DstType: &timestamp.Timestamp{},
			Fn: func(src interface{}) (interface{}, error) {
				temp, _ := src.(time.Time)
				return timestamppb.New(temp), nil
			},
		},
		{
			// timestamp 转 int
			SrcType: &timestamppb.Timestamp{},
			DstType: int64(0),
			Fn: func(src interface{}) (interface{}, error) {
				temp, _ := src.(*timestamppb.Timestamp)
				return temp.AsTime().UnixNano(), nil
			},
		},
	}
}

func Copy(to, from interface{}) error {
	return copier.CopyWithOption(to, from, copier.Option{
		Converters: ConvertList,
	})
}

func DeepCopy(to, from interface{}) error {
	return copier.CopyWithOption(to, from, copier.Option{
		DeepCopy:   true,
		Converters: ConvertList,
	})
}
