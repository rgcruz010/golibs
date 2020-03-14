package filter_test

import (
	"golibs/cmd/filter"
	"reflect"
	"testing"
	"time"
)

func TestParallelFilter(t *testing.T) {
	intPointer := func(num int) *int {
		return &num
	}
	type args struct {
		arr        interface{}
		filterFunc interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    interface{}
	}{
		{
			name: "Success",
			args: args{
				arr: []int{1, 2, 3, 4},
				filterFunc: func(entry int) bool {
					return entry == 1
				},
			},
			wantErr: false,
			want:    []int{1},
		},
		{
			name: "Success",
			args: args{
				arr: []*int{
					intPointer(1),
					intPointer(2),
					intPointer(3),
					intPointer(4),
				},
				filterFunc: func(entry *int) bool {
					return *entry == 1
				}},
			wantErr: false,
			want: []*int{
				intPointer(1),
			},
		},
		{
			name: "Failed",
			args: args{
				arr:        "[]int{1, 2, 3, 4}",
				filterFunc: nil,
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filter.Parallel(tt.args.arr, tt.args.filterFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Simple() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Simple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	intPointer := func(num int) *int {
		return &num
	}
	type args struct {
		arr        interface{}
		filterFunc interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    interface{}
	}{
		{
			name: "Success",
			args: args{
				arr: []int{1, 2, 3, 4},
				filterFunc: func(entry int) bool {
					return entry == 1
				},
			},
			wantErr: false,
			want:    []int{1},
		},
		{
			name: "Success",
			args: args{
				arr: []*int{
					intPointer(1),
					intPointer(2),
					intPointer(3),
					intPointer(4),
				},
				filterFunc: func(entry *int) bool {
					return *entry == 1
				}},
			wantErr: false,
			want: []*int{
				intPointer(1),
			},
		},
		{
			name: "Failed",
			args: args{
				arr:        "[]int{1, 2, 3, 4}",
				filterFunc: nil,
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "Success string slice",
			args: args{
				arr: []string{"1", "2", "3", "4"},
				filterFunc: func(entry string) bool {
					return entry == "3"
				},
			},
			wantErr: false,
			want:    []string{"3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filter.Simple(tt.args.arr, tt.args.filterFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Simple() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Simple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFilterFast(b *testing.B) {
	source := [100]int{}
	for i := 0; i < len(source); i++ {
		source[i] = i + 1
	}
	isMultipliedBy3 := func(num int) bool {
		return num%3 == 0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		filter.Simple(source, isMultipliedBy3)
	}
}

func BenchmarkParallelFilter(b *testing.B) {
	source := [100]int{}
	for i := 0; i < len(source); i++ {
		source[i] = i + 1
	}
	isMultipliedBy3 := func(num int) bool {
		time.Sleep(20 * time.Millisecond)
		return num%3 == 0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		filter.Parallel(source, isMultipliedBy3)
	}
}

func BenchmarkImperative(b *testing.B) {
	source := [100]int{}
	for i := 0; i < len(source); i++ {
		source[i] = i + 1
	}

	isMultipliedBy3 := func(num int) bool {
		time.Sleep(20 * time.Millisecond)
		return num%3 == 0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, num := range source {
			isMultipliedBy3(num)
		}
	}
}

func BenchmarkParallelFilterFast(b *testing.B) {
	source := [100]int{}
	for i := 0; i < len(source); i++ {
		source[i] = i + 1
	}
	isMultipliedBy3 := func(num int) bool {
		return num%3 == 0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		filter.Parallel(source, isMultipliedBy3)
	}
}

func BenchmarkImperativeFast(b *testing.B) {
	source := [100]int{}
	for i := 0; i < len(source); i++ {
		source[i] = i + 1
	}

	isMultipliedBy3 := func(num int) bool {
		return num%3 == 0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, num := range source {
			isMultipliedBy3(num)
		}
	}
}
