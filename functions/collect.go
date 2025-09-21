package functions

func CollectStuff(s []int) []int {
	for i := 0; i < 10; i += 1 {
		s = append(s, i)
	}
	return s
}
