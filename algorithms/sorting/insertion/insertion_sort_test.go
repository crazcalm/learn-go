package insertion

import "testing"

func TestInsertion(t *testing.T) {
    var tests = []struct {
        input []int
        expected []int
        result bool
    }{
        {
            []int{1,2,5,6,3,4},
            []int{1,2,3,4,5,5},
            false,
        },
        {
            []int{6,5,4,3,2,1},
            []int{1,2,3,4,5,6},
            true,
        },
    }

    for _, test :=range tests {
        answer := Insertion(test.input)
        result := isEqual(test.expected, answer)
        if result != test.result {
            t.Errorf("%v =! %v", answer, test.expected)
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
