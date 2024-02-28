package utility
import (
	"fmt"
	"testing"
	"reflect"
)

func TestConvertStringSliceToIntSlice(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []int
		wantErr bool
	}{
		{
			name:    "Valid Integers",
			input:   []string{"1", "2", "3"},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "Invalid Integers",
			input:   []string{"1", "abc", "3"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Empty List",
			input:   []string{},
			want:    []int{},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ConvertStringSliceToIntSlice(tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ConvertStringSliceToIntSlice() error = %v, wantErr %v", err, tc.wantErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ConvertStringSliceToIntSlice() = %v, want %v", got, tc.want)
			}
		})
	}
}