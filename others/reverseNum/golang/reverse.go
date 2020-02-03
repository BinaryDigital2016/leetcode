/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1](-2147483648,2147483647)。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

const MAX = 2<<30 -1
const MIN = -2<<30
func reverse(x int) int {
    result := 0
    for x!=0 {
        a := x%10
        x = x/10
        if result > MAX/10 || (result == MAX/10 && a > 7){
            return 0
        }
        if result < MIN/10 || (result == MIN/10 && a < -8){
            return 0
        }
        tmp := result*10 + a 
        result = tmp

    }
    return result
}
