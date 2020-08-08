package main

type fc func(en bool, args ...interface{}) []interface{}

// Functions ...
var Functions = map[string]fc{
	"EQ_R": EQR,
	"GT_R": GTR,
	"LT_R": LTR,
	"GE_R": GER,
	"LE_R": LER,

	// Mathematics
	"ADD_R": ADDR,
}
