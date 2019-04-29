package utils

import (
	"reflect"
	"testing"
)

func TestQueryParams(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{"basicTest", args{"/hello/world?a=b&c=d"}, map[string]string{"a": "b", "c": "d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParams(tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
