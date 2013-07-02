package merge

func Sort(data []int) {
	if len(data) < 2 {
		return
	}
	Sort(data[:len(data)/2])
	Sort(data[len(data)/2:])
	copy(data, merge(data[:len(data)/2], data[len(data)/2:]))
}

func merge(a []int, b []int) []int {
	res := make([]int, len(a)+len(b))
	ai, bi, i := 0, 0, 0
	for ; i < len(res); i++ {
		if a[ai] < b[bi] {
			res[i] = a[ai]
			ai++
			if ai == len(a) {
				break
			}
		} else {
			res[i] = b[bi]
			bi++
			if bi == len(b) {
				break
			}
		}
	}
	copy(res[i+1:], a[ai:])
	copy(res[i+1:], b[bi:])
	return res
}
