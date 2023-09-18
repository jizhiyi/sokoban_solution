package mapdata

import (
	"sort"
	"testing"
)

func TestPointSlice_Sort(t *testing.T) {
	tests := []struct {
		name string
		ps   PointSlice
		want PointSlice
	}{
		{name: "test", ps: PointSlice{{1, 2}, {3, 4}}, want: PointSlice{{1, 2}, {3, 4}}},
		{name: "test", ps: PointSlice{{3, 4}, {1, 2}}, want: PointSlice{{1, 2}, {3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.ps)
			for i := range tt.ps {
				if tt.ps[i].X != tt.want[i].X && tt.ps[i].Y != tt.want[i].Y {
					t.Errorf("PointSlice.Sort() = %v, want %v", tt.ps, tt.want)
				}
			}
		})
	}
}
