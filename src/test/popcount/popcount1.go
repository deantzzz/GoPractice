package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// practice 2.3
// 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。 （11.4节将展示如何系统地比较两个不同实现的性能。）
func PopCountWithForLoop(x uint64) int {
	var counts int
	for i := 0; i < 8; i++ {
		counts += int(pc[x>>(i*8)])
	}
	return counts
}

// practice 2.4
// 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。
// 比较和查表算法的性能差异。
func PopCountWithBitMove(x uint64) int {
	var counts int
	for i := 0; i < 64; i++ {
		// & 运算，&1，只有最右边第1位是1时&1结果才为1
		if x&1 != 0 {
			counts++
		}
		//右移一位
		x = x >> 1
	}
	return counts
}

// practice 2.5
// 表达式 x&(x-1) 用于将x的最低的一个非零的bit位清零。使用这个算法重写 PopCount函数，
// 然后比较性能
func PopCountWithCleanLowest(x uint64) int {
	// 3	0011
	// 3-1	0010
	// &	0010
	// 2-1  0001
	// &	0000
	var counts int
	for x != 0 {
		x = x & (x - 1)
		counts++
	}
	return counts
}

// 比较所有方法的性能 11.4
func ComparePerformence(x uint64) {
	//待补充
}
