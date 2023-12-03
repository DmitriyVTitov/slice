package slice

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test1",
			args: args{
				s: []string{"a", "b", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test2",
			args: args{
				s: []string{"a", "b", "b", "c", "a", "a", "a", "b", "b", "b"},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}

			t.Log(len(tt.args.s), len(RemoveDuplicates(tt.args.s)))
		})
	}
}

func TestRemoveDuplicatesFunc(t *testing.T) {

	type T struct {
		s string
		i int
		b bool
	}

	f := func(a, b T) bool {
		return a.s == b.s && a.i == b.i
	}

	tc1 := []T{
		{s: "A", i: 1, b: true},
		{s: "B", i: 1, b: true},
		{s: "A", i: 2, b: true},
		{s: "A", i: 1, b: true},
		{s: "B", i: 1, b: true},
		{s: "B", i: 1, b: true},
		{s: "C", i: 3, b: true},
		{s: "C", i: 3, b: false},
	}

	want1 := []T{
		{s: "A", i: 1, b: true},
		{s: "B", i: 1, b: true},
		{s: "A", i: 2, b: true},
		{s: "C", i: 3, b: true},
	}

	tc2 := []T{
		{s: "A", i: 1, b: true},
		{s: "A", i: 1, b: false},
		{s: "A", i: 1, b: true},
		{s: "B", i: 2, b: true},
	}

	want2 := []T{
		{s: "A", i: 1, b: true},
		{s: "B", i: 2, b: true},
	}

	type args struct {
		s     []T
		equal func(a, b T) bool
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		{
			name: "Test #1",
			args: args{
				s:     tc1,
				equal: f,
			},
			want: want1,
		},
		{
			name: "Test #2",
			args: args{
				s:     tc2,
				equal: f,
			},
			want: want2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicatesFunc(tt.args.s, tt.args.equal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DedupFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
