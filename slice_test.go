package slice

import (
	"reflect"
	"testing"
)

type testStruct struct {
	s string
	i int
	b bool
}

var equalFunc = func(a, b testStruct) bool {
	return a.s == b.s && a.i == b.i
}

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
		})
	}
}

func TestRemoveDuplicatesFunc(t *testing.T) {
	tc1 := []testStruct{
		{s: "A", i: 1, b: true},
		{s: "B", i: 1, b: true},
		{s: "A", i: 2, b: true},
		{s: "A", i: 1, b: true},
		{s: "B", i: 1, b: true},
		{s: "B", i: 1, b: true},
		{s: "C", i: 3, b: true},
		{s: "C", i: 3, b: false},
	}

	want1 := []testStruct{
		{s: "A", i: 1, b: true},
		{s: "B", i: 1, b: true},
		{s: "A", i: 2, b: true},
		{s: "C", i: 3, b: true},
	}

	tc2 := []testStruct{
		{s: "A", i: 1, b: true},
		{s: "A", i: 1, b: false},
		{s: "A", i: 1, b: true},
		{s: "B", i: 2, b: true},
	}

	want2 := []testStruct{
		{s: "A", i: 1, b: true},
		{s: "B", i: 2, b: true},
	}

	type args struct {
		s     []testStruct
		equal func(a, b testStruct) bool
	}
	tests := []struct {
		name string
		args args
		want []testStruct
	}{
		{
			name: "Test #1",
			args: args{
				s:     tc1,
				equal: equalFunc,
			},
			want: want1,
		},
		{
			name: "Test #2",
			args: args{
				s:     tc2,
				equal: equalFunc,
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

func TestContainsDuplicates(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s: []int{1, 2, 2, 3},
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				s: []int{1, 2, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicates(tt.args.s); got != tt.want {
				t.Errorf("ContainsDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsDuplicatesFunc(t *testing.T) {
	tc1 := []testStruct{
		{s: "A", i: 1, b: true},
		{s: "B", i: 1, b: true},
		{s: "A", i: 1, b: true},
		{s: "A", i: 4, b: true},
	}

	tc2 := []testStruct{
		{s: "A", i: 1, b: true},
		{s: "B", i: 1, b: true},
		{s: "A", i: 4, b: true},
	}

	type args struct {
		s     []testStruct
		equal func(a, b testStruct) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s:     tc1,
				equal: equalFunc,
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				s:     tc2,
				equal: equalFunc,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicatesFunc(tt.args.s, tt.args.equal); got != tt.want {
				t.Errorf("ContainsDuplicatesFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDuplicatesIndexes(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want map[int][]int
	}{
		{
			name: "Test 1",
			args: args{
				s: []int{1, 2, 1, 1, 3, 2},
			},
			want: map[int][]int{0: {2, 3}, 2: {0, 3}, 3: {0, 2}, 1: {5}, 5: {1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDuplicatesIndexes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDuplicatesIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDuplicatesIndexesFunc(t *testing.T) {
	tc1 := []testStruct{
		{s: "A", i: 0, b: true},
		{s: "B", i: 1, b: true},
		{s: "A", i: 2, b: true},
		{s: "A", i: 3, b: true},
		{s: "B", i: 4, b: true},
		{s: "C", i: 5, b: true},
		{s: "D", i: 76, b: true},
	}

	type args struct {
		s     []testStruct
		equal func(a, b testStruct) bool
	}
	tests := []struct {
		name string
		args args
		want map[int][]int
	}{
		{
			name: "Test 1",
			args: args{
				s:     tc1,
				equal: func(a, b testStruct) bool { return a.s == b.s },
			},
			want: map[int][]int{0: {2, 3}, 2: {0, 3}, 3: {0, 2}, 1: {4}, 4: {1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDuplicatesIndexesFunc(tt.args.s, tt.args.equal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDuplicatesIndexesFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicatesByComparableKey(t *testing.T) {
	tc1 := []testStruct{
		{s: "A", i: 1, b: true},
		{s: "B", i: 2, b: true},
		{s: "A", i: 3, b: true},
		{s: "A", i: 4, b: true},
		{s: "B", i: 5, b: true},
		{s: "B", i: 6, b: true},
		{s: "C", i: 7, b: true},
		{s: "C", i: 8, b: false},
	}

	want1 := []testStruct{
		{s: "A", i: 1, b: true},
		{s: "B", i: 2, b: true},
		{s: "C", i: 7, b: true},
	}

	tc2 := []testStruct{
		{s: "A", i: 1, b: true},
		{s: "A", i: 1, b: false},
		{s: "A", i: 1, b: true},
		{s: "B", i: 2, b: true},
	}

	want2 := []testStruct{
		{s: "A", i: 1, b: true},
		{s: "B", i: 2, b: true},
	}

	type args struct {
		slice []testStruct
		fn    func(testStruct) string
	}
	tests := []struct {
		name string
		args args
		want []testStruct
	}{
		{
			name: "Test #1",
			args: args{
				slice: tc1,
				fn:    func(s testStruct) string { return s.s },
			},
			want: want1,
		},
		{
			name: "Test #2",
			args: args{
				slice: tc2,
				fn:    func(s testStruct) string { return s.s },
			},
			want: want2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicatesByComparableKey(tt.args.slice, tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicatesByComparableKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
