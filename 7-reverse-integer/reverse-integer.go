func reverse(x int) int {
    
    res := 0
    negative := false

    if x < 0 {
       negative = true
       x = -1 * x 
    }

    for x >= 10 {

        res = res * 10 + x % 10
        x = x / 10

    }
    
    res = res * 10 + x
    if negative{
        res *= -1
    }
    if res > 2147483647 || res < -2147483648 {
        return 0
    }

    return res

}