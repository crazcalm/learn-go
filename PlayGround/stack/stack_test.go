package stack

import "testing"

func TestStackEmpty(t *testing.T){
    var s Stack

    if !s.Empty() == true{
        t.Error(`s.Empty() != true`)
    }
}
