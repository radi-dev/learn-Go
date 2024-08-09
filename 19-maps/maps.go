package main

import "fmt"

type vehicle struct {
	canMove, canFly bool
}

func main() {
	var map1 map[string]vehicle
	map1 = make(map[string]vehicle)
	map1["car"] = vehicle{canMove: true}
	map1["helicopter"] = vehicle{canFly: true}
	fmt.Printf("\nmap 1 is: %v\nmap1[car].canFly is: %v\n", map1, map1["car"].canFly)

	map_literal := map[string]vehicle{
		"truck":   {canMove: true, canFly: false},
		"biplane": {true, true}}
	fmt.Printf("\nmap literal is: %v\nmap_literal[biplane].canFly is: %v\nmap_literal[car].canFly is: %v\n\n", map_literal, map_literal["biplane"].canFly, map_literal["car"])

	delete(map_literal, "truck")
	delete(map_literal, "humvee") //key doesn't exist. Op does not fail
	fmt.Println(map_literal["truck"], "\n", map_literal)

	map_literal["triplane"] = vehicle{true, true}
	fmt.Println(map_literal)

	//check if key exists
	key, ok := map_literal["triplane"]
	fmt.Printf("\nVale = %v, is present? = %v\n", key, ok)

	key, ok = map_literal["humvee"]
	fmt.Printf("Vale = %v, is present? = %v\n\n", key, ok)

}

/*
- The zero value of a map is nil.
- A nil map has no keys, nor can keys be added.
- Use make function to create and initialize a map of a given type.
- accessing undefined keys does returns a map whose fields are the default values of their corresponding type
*/

//cd 19-maps
//go run maps.go
