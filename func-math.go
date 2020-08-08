package main

// ADDR ...
var ADDR fc = func(en bool, args ...interface{}) []interface{} {
	var res float64
	for _, in := range args {
		res += in.(float64)
	}

	return []interface{}{res}
}
