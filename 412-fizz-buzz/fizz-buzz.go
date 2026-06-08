func fizzBuzz(n int) []string {
    
    result := make([]string, n) 

    rulse := []struct{
        divisor int
        word string
    }{
        {3, "Fizz"},
        {5, "Buzz"},
    }

    for i:=1; i<=n; i++{
        for _, rule := range rulse{
            if i % rule.divisor == 0{
                result[i-1] += rule.word
            }
        }
        if result[i-1] == ""{
            result[i-1] = strconv.Itoa(i)
        }
    }

    return result

}