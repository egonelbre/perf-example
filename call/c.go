package example

// int nop() { return 0; }
import "C"

func CNop() int { return int(C.nop()) }
