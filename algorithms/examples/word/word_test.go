package word

import "testing"

func TestPalindrome(t *testing.T) {
    if !IsPalindrome("ddddddd") {
        t.Error("IsPalindrome{'dddddd'} = false")
    }
}
