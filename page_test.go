package ctool

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPage(t *testing.T) {
	args := []struct {
		data []string
		page int
		size int
		want []string
	}{
		{
			data: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"},
			page: 3,
			size: 9,
			want: []string{"19", "20"},
		},
		{
			data: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"},
			page: 1,
			size: 10,
			want: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		},
	}
	for _, arg := range args {
		page := Page(arg.data, arg.page, arg.size)
		if reflect.DeepEqual(page, arg.want) {
			fmt.Println("pass")
		} else {
			fmt.Println("fail")
		}
	}
}

func TestPageInfo(t *testing.T) {
	args := []struct {
		page            int
		size            int
		total           int
		wantTotalPage   int
		wantCurrentPage int
		wantPageSize    int
	}{
		{
			page:            3,
			size:            9,
			total:           20,
			wantTotalPage:   3,
			wantCurrentPage: 3,
			wantPageSize:    9,
		},
		{
			page:            1,
			size:            10,
			total:           20,
			wantTotalPage:   2,
			wantCurrentPage: 1,
			wantPageSize:    10,
		},
	}
	for _, arg := range args {
		totalPage, currentPage, pageSize := PageInfo(arg.page, arg.size, arg.total)
		if totalPage == arg.wantTotalPage && currentPage == arg.wantCurrentPage && pageSize == arg.wantPageSize {
			fmt.Println("pass")
		} else {
			fmt.Println("fail")
		}
	}
}
