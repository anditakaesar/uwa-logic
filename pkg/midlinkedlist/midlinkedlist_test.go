package midlinkedlist

import (
	"reflect"
	"testing"
)

func TestBuildAndSearchMiddleNode(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "should return 3",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: buildNodes([]int{3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildAndSearchMiddleNode(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildAndSearchMiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
