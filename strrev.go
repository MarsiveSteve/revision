package piscine

func StrRev(s string) string {
	var i string

	for _, n := range s {
		i = string(n) + i
	}
	return string(i)
}
