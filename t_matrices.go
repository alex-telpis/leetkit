package leetkit

// ParseIntsMatrix is a shorthand for Parse[[][]int](s).
func ParseIntsMatrix(s string) [][]int { return Parse[[][]int](s) }

// ParseStringMatrix is a shorthand for Parse[[][]string](s).
func ParseStringMatrix(s string) [][]string { return Parse[[][]string](s) }

// ParseByteMatrix is a shorthand for Parse[[][]byte](s).
func ParseByteMatrix(s string) [][]byte { return Parse[[][]byte](s) }
