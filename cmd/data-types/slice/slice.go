package slice

func BasicUse() {
	// slices的创建 可以通过数组的创建和转slice的方式

	likeArr := []int{1, 2, 3}
	println("数组不设置大小就是切片")
	println(likeArr)

	// 也可以通过make
	oneSlice := make([]int, 0, 5)
	println("oneSlice 的length是: ", len(oneSlice), ", 容量是:", cap(oneSlice))

	println("添加一个元素后:")
	oneSlice = append(oneSlice, 1)
	println("oneSlice 的length是: ", len(oneSlice), ", 容量是:", cap(oneSlice))

	println("清空slice后:")
	oneSlice = oneSlice[:0]
	println("oneSlice 的length是: ", len(oneSlice), ", 容量是:", cap(oneSlice))

	for i := range 10 {
		println("添加", i+1, "个元素后:")
		oneSlice = append(oneSlice, i)
		println("oneSlice 的length是: ", len(oneSlice), ", 容量是:", cap(oneSlice))
	}

	s1 := []int{1, 2, 3}
	s2 := make([]int, 3)

	copy(s2, s1)
	println("修改s1的值:")
	s1[0] = 0
	println("查看s2的值:")
	for _, item := range s2 {
		print(item, " ")
	}
}

func AdvancedUse() {

}
