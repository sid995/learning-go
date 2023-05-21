package main

import "fmt"

// func main() {
// 	x := 1
// 	p := &x

// 	fmt.Println(p)
// 	*p = 2
// 	fmt.Println(x)
// }

// func main() {
// 	v := 1
// 	incr(&v)
// 	fmt.Println(incr(&v))
// }

// func incr(p *int) int {
// 	*p++
// 	return *p
// }

// var n = flag.Bool("n", false, "omit trailing newline")
// var sep = flag.String("s", " ", "seperator")

// func main() {
// 	flag.Parse()
// 	fmt.Print(strings.Join(flag.Args(), *sep))
// 	if !*n {
// 		fmt.Println()
// 	}
// }

// func main() {
// 	p := new(int)
// 	fmt.Println(*p)
// 	*p = 2
// 	fmt.Println(*p)
// }

type Player struct {
	health int
}

/*
function becomes part of the struct Player
can be accessed using the definition variable in the main function

usage -> player.takeDamageFromExplosion(10)

the memory location of the player struct is accessed and the data in
that memory is modified and returned in the main func
*/
func (player *Player) takeDamageFromExplosion(explosionDmg int) {
	player.health -= explosionDmg
}

/*
default way the function is declared. this is with dynamic 'dmg' argument
the player struct is an arg of the function is a dereferenced value
*/
func takeDamageFromExplosion(player *Player, explosionDmg int) {
	player.health -= explosionDmg
}

func main() {
	player := &Player{
		health: 100,
	}

	fmt.Println(player)
	takeDamageFromExplosion(player, 80)
	fmt.Println(player)
}
