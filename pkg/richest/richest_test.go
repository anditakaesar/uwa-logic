package richest

import "testing"

func TestGetRichest(t *testing.T) {
	type args struct {
		bank [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 15",
			args: args{
				bank: [][]int{
					{1, 2, 3, 4, 5},
				},
			},
			want: 15,
		},
		{
			name: "should return 21",
			args: args{
				bank: [][]int{
					{1, 2, 3, 4, 5},
					{1, 2, 3, 4, 5, 6},
				},
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRichest(tt.args.bank); got != tt.want {
				t.Errorf("GetRichest() = %v, want %v", got, tt.want)
			}
		})
	}
}
