package main

var toMap = map[int]int{
	595: 595,
}

func mapTiles(from [][]int) [][]int {

	to := make([][]int, len(from))
	for x := range from {
		for y := range from[x] {
			to[x] = make([]int, len(from[x]))
			t, ok := toMap[y]
			if !ok {
				t = from[x][y]
			}
			to[x][y] = t
		}
	}
	return to
}

type Level struct {
	Tiles [][]int
}

var tiles = [][]int{
	{
		595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595,
		595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595,
		595, 595, 595, 595, 595, 145, 145, 145, 145, 145, 145, 145, 595, 595, 595, 595, 595,
		595, 595, 615, 595, 595, 145, 297, 297, 297, 297, 297, 145, 595, 615, 595, 595, 595,
		595, 595, 595, 595, 595, 145, 297, 297, 297, 297, 297, 145, 595, 595, 595, 595, 595,

		595, 595, 595, 595, 595, 145, 297, 297, 297, 297, 297, 145, 595, 595, 595, 595, 595,
		595, 595, 595, 615, 595, 145, 297, 297, 297, 297, 297, 145, 595, 595, 595, 595, 595,
		595, 595, 615, 595, 595, 145, 297, 297, 297, 297, 297, 145, 595, 595, 595, 595, 595,
		595, 595, 595, 595, 595, 145, 297, 297, 297, 297, 297, 145, 595, 615, 595, 595, 595,
		595, 595, 595, 595, 595, 145, 145, 145, 703, 145, 145, 145, 595, 595, 595, 595, 595,

		595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595,
		595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 615, 595, 595,
		595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595,
		595, 615, 595, 595, 595, 595, 595, 595, 595, 595, 595, 615, 595, 595, 595, 595, 595,
		595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595, 595,
	},
	// {
	// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 26, 27, 28, 29, 30, 31, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 51, 52, 53, 54, 55, 56, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 76, 77, 78, 79, 80, 81, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 101, 102, 103, 104, 105, 106, 0, 0, 0, 0,
	//
	// 	0, 0, 0, 0, 0, 126, 127, 842, 842, 130, 131, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 303, 303, 245, 242, 303, 303, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
	//
	// 	0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
	// },
}