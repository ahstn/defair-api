package actions

import (
	"reflect"
	"testing"

	"github.com/ahstn/defair/domain"
)

func Test_networkFilter(t *testing.T) {
	tests := []struct {
		name string
		args domain.DataFilter
		want []string
	}{
		{
			name: "Supplied Filter of 'all'",
			args: domain.DataFilter{Networks: []string{"all"}},
			want: []string{"avalanche", "harmony", "aurora"},
		},
		{
			name: "Supplied Filter of single network",
			args: domain.DataFilter{Networks: []string{"avalanche"}},
			want: []string{"avalanche"},
		},
		{
			name: "Supplied Filter of multiple networks",
			args: domain.DataFilter{Networks: []string{"harmony", "aurora"}},
			want: []string{"harmony", "aurora"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkFilter(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("networkFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
