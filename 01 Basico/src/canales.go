package src

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text
}
func Test() {
	c := make(chan string, 2)
	c <- "text"

	fmt.Println(len(c), cap(c))
	//range and close
	close(c) // cierra el canal

	for message := range c {
		fmt.Println(message)
	}

	// select
	email1 := make(chan string)
	email2 := make(chan string)
	email3 := make(chan string)
	go message("mensaje", email1)
	go message("mensaje2", email2)
	go message("mensaje3", email3)
	for i := 0; i < 3; i++ {
		select { //para manejar mulpiples canales
		case m1 := <-email1:
			fmt.Println("email de 1", m1)
		case m2 := <-email2:
			fmt.Println("email de 2", m2)
		case m3 := <-email3:
			fmt.Println("email de 3", m3)
		}

	}
}
