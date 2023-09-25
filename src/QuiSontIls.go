package PROJETRED

import (
	"fmt"
	"time"
)

func (p Perso) QuiSontIls() {
	fmt.Println("			 ██████╗ ██╗   ██╗██╗    ███████╗ ██████╗ ███╗   ██╗████████╗    ██╗██╗     ███████╗    ██████╗ ")
	fmt.Println("			██╔═══██╗██║   ██║██║    ██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝    ██║██║     ██╔════╝     ╚════██╗")
	fmt.Println("			██║   ██║██║   ██║██║    ███████╗██║   ██║██╔██╗ ██║   ██║       ██║██║     ███████╗       ▄███╔╝")
	fmt.Println("			██║▄▄ ██║██║   ██║██║    ╚════██║██║   ██║██║╚██╗██║   ██║       ██║██║     ╚════██║       ▀▀══╝ ")
	fmt.Println("			╚██████╔╝╚██████╔╝██║    ███████║╚██████╔╝██║ ╚████║   ██║       ██║███████╗███████║       ██╗  ")
	fmt.Println("			 ╚══▀▀═╝  ╚═════╝ ╚═╝    ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝       ╚═╝╚══════╝╚══════╝       ╚═╝")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(1 * time.Second)
	fmt.Println("							Partie 2 :			ABBA - Gimme! Gimme! Gimme!")
	fmt.Println("							Partie 3 :				Steven Spielberg")
	fmt.Println("							Partie 4 :			Queen - I Want To Break Free")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Scan(&choice)

	switch choice {
	case "1": 
	fmt.Println("\033[H\033[2J")
		p.Menu()
	default : 
	fmt.Println("\033[H\033[2J")
	fmt.Println("Recommence mon donuts sucré au sucre ! ")
	p.QuiSontIls()
	}
}
	  
																								 
 

   
   
  
                                                                                        
