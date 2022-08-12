package simple_factory

import (
	"reflect"
	"testing"
)

func TestRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: JsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: YamlRuleConfigParser{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}

}
