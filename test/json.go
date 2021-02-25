//	{
//	"name": {"first": "Tom", "last": "Anderson"},
//	"age":37,
//	"children": ["Sara","Alex","Jack"],
//	"fav.movie": "Deer Hunter",
//	"friends": [
//				{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
//				{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
//				{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
//				]
//	}

//	"name.last"          >> "Anderson"
//	"age"                >> 37
//	"children"           >> ["Sara","Alex","Jack"]
//	"children.#"         >> 3
//	"children.1"         >> "Alex"
//	"child*.2"           >> "Jack"
//	"c?ildren.0"         >> "Sara"
//	"fav\.movie"         >> "Deer Hunter"
//	"friends.#.first"    >> ["Dale","Roger","Jane"]
//	"friends.1.last"     >> "Craig"

package main

import "rule_engine_by_go/utils/json"

const json_ = `{"name":{"first":"Janet","last.name":"Prichard"},"age":47}`

func main() {
	value := json.Get(json_, "name.last\\.name")
	println(value.String())
}
