package functions 

func CollectStuff() []int {
	s := []int{}
	for i := 0; i < 12; i += 2 {
		s = append(s, i)
	}
	return s
}
