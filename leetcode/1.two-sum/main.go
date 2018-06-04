package main

import "fmt"
import "strings"
import "strconv"


func str2Array(in string) []int {
    strings := strings.Split(in, ",")
    ints := make([]int, len(strings))
    for i, s := range strings {
        ints[i], _ = strconv.Atoi(s)
    }
    return ints
}

func add2Num(in []int, target int) []int {
    length := len(in)
    ret := make([]int, 2)
    for i := 0; i < length; i++ {
        for j := i; j < length; j++ {
            if in[i] + in[j] == target {
                ret[0] = in[i]
                ret[1] = in[j]
                return ret
            }
        }
    }
    return ret
}

func main() {
    input := "2,7,11,15"
    target := 9

    fmt.Printf("input: %v\ntarget: %v\n", input, target)

    ret := add2Num(str2Array(input), target)
    fmt.Printf("%v\n", ret)
}
