package value

import "testing"

func TestIsStaleNaN(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "error",
			args: args{v: 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStaleNaN(tt.args.v); got != tt.want {
				t.Errorf("IsStaleNaN() = %v, want %v", got, tt.want)
			}
		})
	}
}
