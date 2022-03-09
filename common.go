package flashbench

import (
	"fmt"
	"os"
)

func getKey(n int) []byte {
	return []byte("key_" + fmt.Sprintf("%07d", n))
}

func getMember(n int) []byte {
	return []byte("member_" + fmt.Sprintf("%07d", n))
}

func geyValue64B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}

func geyValue128B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}

func geyValue256B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}

func geyValue512B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv" + "valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}

func isPathOk(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
