package insertion

import "testing"

func TestInsertion(t *testing.T) {
    var tests = []struct {
        input []int
        expected []int
    }{
        {
            []int{1,2,5,6,3,4},
            []int{1,2,3,4,5,5},
        },
    }

    for _, test :=range tests {
        answer := Insertion(test.input)
        result := isEqual(test.expected, answer)
        if result == false {
            t.Errorf("%s =! %s", answer, test.expected)
        }
    }
}


func isEqual(list1[]int, list2[]int) bool {
    if len(list1) != len(list2) {
        return false
    }

    for index, item :=range list1 {
        if item != list2[index] {
            return false
        }
    }
    return true
}
