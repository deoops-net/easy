package easy

import "testing"

func TestRootDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test root dir",
			want: ".",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RootDir(); got != tt.want {
				t.Errorf("RootDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
