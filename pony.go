package pony

import "strconv"

// Text makes phrase colorful. Offset adjusts the rainbow effect.
func Text(phrase string, offset int) string {
	var parts [][]byte
	appendc := func(x int) {
		var p []byte
		p = strconv.AppendInt(append(p, "\x1B[38;5;"...), int64(x), 10)
		p = append(append(p, 'm', ' '), "\x1B[0m"...)
		parts = append(parts, p)
	}
	x := 196
	for i := 0; i < 2; i++ {
		for j := 0; j < 7; j++ {
			for k := 0; k < 2; k++ {
				for l := 0; l < 6; l++ {
					if l > 0 {
						if k == 0 {
							if i == 0 {
								x = x - 36 + 1
							} else {
								x = x - 36
							}
						} else {
							if i == 0 {
								x = x + 36
							} else {
								x = x - 1 + 36
							}
						}
					}
					appendc(x)
				}
			}
		}
	}
	var s []byte
	for i := 0; i < len(parts) && i < len(phrase); i++ {
		j := (i + offset) % len(parts)
		if j < 0 {
			j += len(parts)
		}
		p := parts[j]
		p[len(p)-5] = phrase[i%len(phrase)]
		s = append(s, p...)
	}
	return string(s)
}
