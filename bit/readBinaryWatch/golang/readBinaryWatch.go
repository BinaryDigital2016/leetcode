package bit

import "strconv"

/*
二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。

每个 LED 代表一个 0 或 1，最低位在右侧。

给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。

案例:

输入: n = 1
返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
*/

func readBinaryWatch(num int) []string {
	ret := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if count1(i)+count1(j) == num { //遍历时分中1的和
				hour := strconv.Itoa(i)
				min := strconv.Itoa(j)
				if j < 10 {
					min = "0" + min
				}
				ret = append(ret, hour+":"+min)
			}
		}
	}
	return ret
}

func count1(n int) int {
	ret := 0
	for n > 0 {
		ret += 1
		n = n & (n - 1)
	}
	return ret
}
