func romanToInt(s string) int {

    romanDictionary := map[rune]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }

    result := 0

    for i:=0; i < len(s); i++{

        current := romanDictionary[rune(s[i])]
        if i+1 < len(s){
            if current < romanDictionary[rune(s[i+1])]{
                result -= current
            }else{
                result += current    
            }
        }
        if i+1 == len(s){
            result += current    
        }

    }
    return result

}