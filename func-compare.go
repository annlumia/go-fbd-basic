package main

// EQR ...
var EQR fc = func(en bool, args ...interface{}) []interface{} {
	return []interface{}{en && (args[0].(float64) == args[1].(float64))}
}

// GTR ...
var GTR fc = func(en bool, args ...interface{}) []interface{} {
	return []interface{}{en && (args[0].(float64) > args[1].(float64))}
}

// LTR ...
var LTR = func(en bool, args ...interface{}) []interface{} {
	return []interface{}{en && (args[0].(float64) < args[1].(float64))}
}

// GER ...
var GER = func(en bool, args ...interface{}) []interface{} {
	return []interface{}{en && (args[0].(float64) >= args[1].(float64))}
}

// LER ...
var LER = func(en bool, args ...interface{}) []interface{} {
	return []interface{}{en && (args[0].(float64) <= args[1].(float64))}
}
