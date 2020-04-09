package main

import (
	"fmt"
)

// func main(){
// 	var creature = "shark"
// 	fmt.Println("creature 1 : ",creature)
// 	printValue(&creature)
// 	fmt.Printf("updated creature T : %T V : %v\n",creature,creature)

// }
// func printValue(creature *string){
// 	*creature = "jelly fish"
// 	fmt.Printf("Creature 2 :%+v\n",creature)
// 	fmt.Printf("Creature 2 T:%T\n",creature)

// }

type creature struct {
	species string
}

func main() {
	creature := creature{
		species: "shark",
	}
	fmt.Println(creature)
	changeCreature(&creature)
	fmt.Println("after : ", creature)

}

func changeCreature(creature *creature) {
	fmt.Println("change creature before", *creature)
	creature.species = "Jelly fish"
}
