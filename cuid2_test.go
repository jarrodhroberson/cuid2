package GOCUID2

import (
	"testing"
)

//TODO: make these test determistic

func TestNew(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want CUID2
	}{
		{
			name: "new 24",
			args: args{l: 24},
			want: "jy95r6ow2f1tu7wipv9d6ian",
		},
		{
			name: "new 36",
			args: args{l: 36},
			want: "jy95r6ow2f1tu7wipv9d6ian",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.l); got != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
