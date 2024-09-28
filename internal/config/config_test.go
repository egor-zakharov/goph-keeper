package config

import (
	"reflect"
	"testing"
)

func TestConfig_ParseFlag(t *testing.T) {
	type fields struct {
		FlagLogLevel    string
		FlagDB          string
		FlagRunGRPCAddr string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				FlagLogLevel:    "1",
				FlagDB:          "1",
				FlagRunGRPCAddr: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				FlagLogLevel:    tt.fields.FlagLogLevel,
				FlagDB:          tt.fields.FlagDB,
				FlagRunGRPCAddr: tt.fields.FlagRunGRPCAddr,
			}
			c.ParseFlag()
		})
	}
}

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "success",
			want: &Config{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
