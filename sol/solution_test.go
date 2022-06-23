package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	for idx := 0; idx < b.N; idx++ {
		minCostConnectPoints(points)
	}
}
func Test_minCostConnectPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[0,0],[2,2],[3,10],[5,2],[7,0]",
			args: args{points: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}},
			want: 20,
		},
		{
			name: "[3,12],[-2,5],[-4,1]",
			args: args{points: [][]int{{3, 12}, {-2, 5}, {-4, 1}}},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPoints(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
