package word

func IsPalindrome(s string) bool {
    for i:= range s{
        if s[i] != s[len(s) -1] {
            return false
        }
    }
    return true
}
