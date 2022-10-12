package main

import "testing"

func TestTDD(t *testing.T) {
	t.Parallel()

	t.Cleanup(func() {
		t.Log("all test is passed")
	})

	type args struct {
		i int
		j int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				i: 1,
				j: 1,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				i: 1,
				j: 1,
			},
			want:    3,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Add(tt.args.i, tt.args.j)
			if tt.want != got && !tt.wantErr {
				t.Errorf("%s is failed: want = %d, got = %d", tt.name, tt.args.i, tt.args.j)
				return
			}
		})
	}
}
