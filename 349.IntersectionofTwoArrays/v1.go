func intersection(nums1 []int, nums2 []int) (r []int) {
	m1 := toMap(nums1)
	m2 := toMap(nums2)
	for k, _ := range m1 {
		if m2[k] {
			r = append(r, k)
		}
	}
	return r
}

func toMap(t []int) map[int]bool {
	m := map[int]bool{}
	for _, v := range t {
		if !m[v] {
			m[v] = true
		}
	}
	return m
}

