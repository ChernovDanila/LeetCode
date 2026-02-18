func isValid(s string) bool {

   if len(s) % 2 != 0 {
    return false
   }

   stack := []rune{}
   pairs := map[rune]rune{
        ')':'(',
        '}':'{',
        ']':'[',
   }
   
    for _,ch := range(s){
        
        if openBracket, exists := pairs[ch]; exists{

            if len(stack) == 0 || stack[len(stack)-1] != openBracket{
                return false
            }

            stack = stack[:len(stack) -1]

        }else{
            stack = append(stack, ch)
        }
    }

     return len(stack) == 0

}