package utils

func ToUpper(s string) string {
    result := ""
    for _, c := range s {
        if c >= 'a' && c <= 'z' {
            c = c - 'a' + 'A'
        }
        result += string(c)
    }
    return result
}
