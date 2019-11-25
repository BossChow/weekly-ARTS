package main

import "testing"

func TestFindRemovedIndex(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "空数组",
			args: args{},
			want: -1,
		},
		{
			name: "正值数组",
			args: args{[]int{7, 2, 3, 4, 5}},
			want: 1,
		},
		{
			name: "正值含0数组",
			args: args{[]int{0, 1, 2, 3, 4}},
			want: 0,
		},
		{
			name: "含奇数负值数组",
			args: args{[]int{1, 2, -1, 3, 4}},
			want: 2,
		},
		{
			name: "含偶数负值数据",
			args: args{[]int{1, -1, -2, 3, 4}},
			want: 0,
		},
		{
			name: "全部奇数负值数组",
			args: args{[]int{-1, -2, -3, -4, -5}},
			want: 0,
		},
		{
			name: "全部偶数负值数组",
			args: args{[]int{-1, -2, -3, -4}},
			want: 3,
		},
		{
			name: "负值含0数组",
			args: args{[]int{-1, -2, -3, 0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindRemovedIndex(tt.args.arr); got != tt.want {
				t.Errorf("FindRemovedIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
