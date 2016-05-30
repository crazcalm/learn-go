package insertion

import "fmt"

func Insertion(list[] int) [] int {
    for i := 1; i < len(list); i++{
        key := list[i]
        // inser list[i] in sorted sequence list[0...i - 1]
        k := i -1
        for k >= 0 {
            if list[k] <= key{
                break
            }
            list[k + 1] = list[k]
            k = k - 1
        }
        list[k + 1] = key
    }

    return list
}

func main(){
    test := []int{3,2,5,6,2,1}
    answer := Insertion(test)
    for _, item := range answer {
        fmt.Println(item)
    }
}
