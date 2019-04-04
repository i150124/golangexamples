package golangexamples

import "github.com/ehteshamz/greetings"

func ConcatSlice (sliceToConcat []byte) string {
    res := make([]byte,0)
    for i:=1; i<=len(sliceToConcat);i++ {
        res = append(res,sliceToConcat[i-1])
        if (i!=len(sliceToConcat)){
            res = append(res,'-')
        }
    }
    return string(res)
}

func Encrypt (sliceToEncrypt *[]byte, ceaserCount int) {
    for i:=1; i<=len(*sliceToEncrypt);i++ {
        (*sliceToEncrypt)[i-1] = ((*sliceToEncrypt)[i-1]+byte(ceaserCount))%255
    }
}

func EZGreetings(name string) string {
    return greetings.PrintGreetings(name)
}
