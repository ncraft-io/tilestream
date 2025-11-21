package postgis

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

func NewConfig(options core.Options) (*Config, error) {
	cfg := &Config{}
	str, err := jsoniter.Marshal(options)
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(str, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewConfigFrom(options *core.Object) (*Config, error) {
	cfg := &Config{}
	str, err := jsoniter.Marshal(options)
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(str, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (x *Config) ToOptions() core.Options {
	if x != nil {
		options := make(core.Options)
		str, _ := jsoniter.Marshal(x)
		_ = jsoniter.Unmarshal(str, options)
		return options
	}
	return nil
}

func (x *Config) ToObject() *core.Object {
	if x != nil {
		obj := core.NewObject()
		str, _ := jsoniter.Marshal(x)
		_ = jsoniter.Unmarshal(str, obj)
		return obj
	}
	return nil
}

func (x *Config) SetFilter(filter string) *Config {
	if x != nil && x.Provider != nil && x.Provider.Sql != nil {
		x.Provider.Sql.Filter = filter
	}
	return x
}

func (x *Config) GetSql() string {
	if x != nil && x.Provider != nil && x.Provider.Sql != nil {
		prd := x.Provider
		sql := fmt.Sprintf("SELECT ST_AsBinary(%s) AS geometry, %s, %s as feat_id", prd.GeometryField, prd.IdField, prd.IdField)
		for _, field := range prd.Sql.Fields {
			if field.Name == prd.GeometryField || field.Name == prd.IdField {
				continue
			}
			if len(field.Alias) > 0 {
				sql += fmt.Sprintf(", %s as %s", field.Name, field.Alias)
			} else {
				sql += fmt.Sprintf(", %s", field.Name)
			}
		}
		if len(x.Provider.Sql.Filter) > 0 {
			sql += fmt.Sprintf(" FROM %s WHERE geometry && !BBOX! and ( %s )", x.Name, x.Provider.Sql.Filter)
		} else {
			sql += fmt.Sprintf(" FROM %s WHERE geometry && !BBOX!", x.Name)
		}
		return sql
	}

	return ""
}
