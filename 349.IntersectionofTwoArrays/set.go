type Set struct {
	m map[int]bool
}

func (s Set) Intersection(set *Set) *Set {
	r := NewSet([]int{})
	for k, _ := range set.m {
		if s.m[k] {
			r.m[k] = true
		}
	}
	return r
}

func NewSet(num []int) *Set {
	set := &Set{}
	set.m = map[int]bool{}
	for _, v := range num {
		set.m[v] = true
	}
	return set
}

func intersection(nums1 []int, nums2 []int) (r []int) {
	s1 := NewSet(nums1)
	s2 := NewSet(nums2)
	s3 := s1.Intersection(s2)

	for k, _ := range s3.m {
		r = append(r, k)
	}

	return r
}



