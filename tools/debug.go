package tools

// var (
// 	goroutinePrefix = []byte("goroutine ")
// 	errBadStack     = errors.New("invalid runtime.Stack output")
// )

// // This is terrible, slow, and should never be used.
// func GoID() (int, error) {
// 	buf := make([]byte, 32)
// 	n := runtime.Stack(buf, false)
// 	buf = buf[:n]
// 	// goroutine 1 [running]: ...

// 	buf, ok := bytes.CutPrefix(buf, goroutinePrefix)
// 	if !ok {
// 		return 0, errBadStack
// 	}

// 	i := bytes.IndexByte(buf, ' ')
// 	if i < 0 {
// 		return 0, errBadStack
// 	}

// 	return strconv.Atoi(string(buf[:i]))
// }
