package stepstozero

import "testing"

func TestCountStepsToZero(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 1 from 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "should return 2 from 3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "should return 6 from 14",
			args: args{
				n: 14,
			},
			want: 6,
		},
		{
			name: "should return 4 from 8",
			args: args{
				n: 8,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountStepsToZero(tt.args.n); got != tt.want {
				t.Errorf("CountStepsToZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
