//go:build !duckdb_use_lib && alpine

package duckdb

/*
#cgo LDFLAGS: -lduckdb
#cgo LDFLAGS: -lstdc++ -lm -ldl -L${SRCDIR}/deps/alpine_amd64
#include <duckdb.h>
*/
import "C"
