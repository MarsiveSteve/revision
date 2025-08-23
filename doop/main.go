package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	res := doop(args[0], args[1], args[2])
	os.Stdout.WriteString(res)
}

func doop(s0, o, s2 string) string {
	if o != "+" && o != "-" && o != "/" && o != "*" && o != "%" {
		return ""
	}

	n1, ok1 := atoi64(s0)
	n2, ok2 := atoi64(s2)
	if !ok1 || !ok2 {
		return ""
	}

	switch o {
	case "+":
		if addOverflow(n1, n2) {
			return ""
		}
		return itoa64(n1+n2) + "\n"
	case "-":
		if subOverflow(n1, n2) {
			return ""
		}
		return itoa64(n1-n2) + "\n"
	case "*":
		if mulOverflow(n1, n2) {
			return ""
		}
		return itoa64(n1*n2) + "\n"
	case "/":
		if n2 == 0 {
			return "No division by 0\n"
		}
		return itoa64(n1/n2) + "\n"
	case "%":
		if n2 == 0 {
			return "No modulo by 0\n"
		}
		return itoa64(n1%n2) + "\n"
	}
	return ""
}

func atoi64(s string) (int64, bool) {
	var n int64
	var sign int64 = 1
	start := 0
	if len(s) == 0 {
		return 0, false
	}
	if s[0] == '+' {
		start = 1
	} else if s[0] == '-' {
		sign = -1
		start = 1
	}
	for i := start; i < len(s); i++ {
		d := s[i]
		if d < '0' || d > '9' {
			return 0, false
		}
		digit := int64(d - '0')
		if willOverflow(n, digit) {
			return 0, false
		}
		n = n*10 + digit
	}
	return n * sign, true
}

func itoa64(n int64) string {
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	absN := n
	if n < 0 {
		absN = -n
	}
	for absN > 0 {
		i--
		buf[i] = byte(absN%10) + '0'
		absN /= 10
	}
	if n < 0 {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}

func willOverflow(a, b int64) bool {
	if b > 0 && a > (mathMaxInt64()-b) {
		return true
	}
	if b < 0 && a < (mathMinInt64()-b) {
		return true
	}
	return false
}

func addOverflow(a, b int64) bool {
	return willOverflow(a, b)
}

func subOverflow(a, b int64) bool {
	return willOverflow(a, -b)
}

func mulOverflow(a, b int64) bool {
	if a == 0 || b == 0 {
		return false
	}
	if a == mathMinInt64() && b == -1 {
		return true
	}
	if b == mathMinInt64() && a == -1 {
		return true
	}
	result := a * b
	if result/b != a {
		return true
	}
	return false
}

func mathMaxInt64() int64 {
	return 9223372036854775807
}

func mathMinInt64() int64 {
	return -9223372036854775808
}
