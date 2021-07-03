package main

import "fmt"

func main() {
	idMap := make(map[string]int)
	idMap["david"] = 456
	idMap["hussain"] = 985
	id2Map := map[string]int{"joe": 123}
	fmt.Println(id2Map["joe"])
	fmt.Println(idMap)

	delete(id2Map, "joe")
}
