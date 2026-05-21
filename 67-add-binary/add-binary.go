func addBinary(a string, b string) string {
    
    result := ""
    increment := false

    i,j := len(a)-1, len(b)-1

    for i >=0 || j >=0 || increment{
        
        digitA := byte('0')
        if i >= 0 {
            digitA = a[i]
        }
        
        digitB := byte('0')
        if j >= 0 {
            digitB = b[j]    
        }
    
        if digitA == '1' && digitB == '1' && !increment{
            
            result += "0"
            increment = true

        }else if digitA == '1' && digitB == '1' && increment{
            
            result += "1"
            increment = true
        
        }else if digitA == '1' || digitB == '1'{

            if increment {
                result += "0"
            }else {
                result += "1"
                increment = false
            }
            
        }else{

            if increment{
                result += "1"
            }else{
                result += "0"
            }
            
            increment = false
        }
        j--
        i--
    } 

    runes := []byte(result)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)         

}