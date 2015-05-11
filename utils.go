// Utility functions mainly for manipulating and parsing the returned JSON
package git

import (
	"fmt"
)

// Given a slice of maps this function returns a slice of the values for a
// top level key. Eg: The "login" key for a slice of users
func ValuesForKey(key string, js []map[string]interface{}) []interface{} {
	fin := make([]interface{}, 0)
	for _, item := range js {
		fin = append(fin, item[key])
	}

	return fin
}

// Converts a slice of interface{}s to a slice of strings
func StringifyInterfaceSlice(slc []interface{}) []string {
	fin := make([]string, 0)

	for _, x := range slc {
		if val, ok := x.(string); ok {
			fin = append(fin, val)
		} else {
			panic(fmt.Sprintf("%v Failed to convert to string", x))
		}
	}
	return fin
}
