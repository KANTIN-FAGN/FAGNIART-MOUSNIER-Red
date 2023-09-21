package PROJETRED

import (
	"fmt"
	"time"
)

func (p Perso) MenuDemarrage() {

	fmt.Println("                                                                                   * ")
	fmt.Println("                                                                                                  * ")
	fmt.Println("                                                                   *                                             *  ")
	fmt.Println("                                                                                                        *  ")
	fmt.Println("                                                                                            * ")
	fmt.Println("                                                                                                                     *  ")
	fmt.Println("                                                                 *  ")
	fmt.Println("                                                                                                         *  ")
	fmt.Println("                                                                                   *             * ")
	fmt.Println("                                                                 *                                                               * ")
	fmt.Println("                                                                                          (             ) ")
	fmt.Println("                                                                                 )      (*)           (*)      ( ")
	fmt.Println("                                                                         *       (*)      |             |      (*) ")
	fmt.Println("                                                                                  |      |~|           |~|      |          * ")
	fmt.Println("                                                                                 |~|     | |           | |     |~| ")
	fmt.Println("                                                                                 | |     | |           | |     | | ")
	fmt.Println("                                                                                ,| |a@@@@| |@@@@@@@@@@@| |@@@@a| |. ")
	fmt.Println("                                                                           .,a@@@| |@@@@@| |@@@@@@@@@@@| |@@@@@| |@@@@a,. ")
	fmt.Println("                                                                         ,a@@@@@@| |@@@@@@@@@@@@.@@@@@@@@@@@@@@| |@@@@@@@a, ")
	fmt.Println("                                                                       a@@@@@@@@@@@@@@@@@@@@@' . `@@@@@@@@@@@@@@@@@@@@@@@@a  ")
	fmt.Println("                                                                        ;`@@@@@@@@@@@@@@@@@@'   .   `@@@@@@@@@@@@@@@@@@@@@'; ")
	fmt.Println("                                                                        ;@@@`@@@@@@@@@@@@@'     .     `@@@@@@@@@@@@@@@@'@@@; ")
	fmt.Println("                                                                       ;@@@;,.aaaaaaaaaa       .       aaaaa,,aaaaaaa,;@@@; ")
	fmt.Println("                                                                        ;;@;;;;@@@@@@@@;@      @.@      ;@@@;;;@@@@@@;;;;@@; ")
	fmt.Println("                                                                       ;;;;;;;@@@@;@@;;@    @@ . @@    ;;@;;;;@@;@@@;;;;;;;  ")
	fmt.Println("                                                                       ;;;;;;;;@@;;;;;;;  @@   .   @@  ;;;;;;;;;;;@@;;;;@;; ")
	fmt.Println("                                                                       ;;;;;;;;;;;;;;;;;@@     .     @@;;;;;;;;;;;;;;;;@@a; ")
	fmt.Println("                                                                       %;;;;;;;;@;;;;;;;;       .       ;;;;;;;;;;;;;;;;@@;; ")
	fmt.Println("                                                                       %;;;;;;;a@;;;;;;;;     ,   ,     ;;;;;;;;;;;;;;;;;;;;")
	fmt.Println("                                                                       %;;;;;;;@@;;;;;;;;   ,       ,   ;;;;;;;;;;;;;;;;;;;;")
	fmt.Println("                                                                       `;;;;;;;;;;;;;;;;                ;;;;;;;;;;;;;;;;;;;'")
	fmt.Println("                                                                          `;;;;;;;;;;;;,               ,;;;;;;;;;;;;;;;'")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(3 * time.Second)

	introduction := `
	En plein cœur d'un royaume gourmand nommé Pâtissia, les pâtisseries ont pris vie et évolué en créatures magiques et délicieusement périlleuses. Chaque recoin de ce monde est une symphonie de douceurs où les gâteaux rient, les tartes dansent et les cupcakes gambadent. Cependant, l'équilibre sucré de Pâtissia est en péril.

	La recette ancestrale du Gâteau d'Harmonie, gardienne de l'équilibre entre le monde des hommes et celui des pâtisseries, a été dérobée par le maléfique Chef de l'Ombre. Sans cette recette, le chaos sucré menace d'engloutir tout Pâtissia, et même le monde des hommes.
	
	Vous incarnez une courageuse pâtisserie aventurière, armée de sa pâte à choux enchantée et d'un moule protecteur, prête à sauver Pâtissia et le monde des hommes. Votre quête : retrouver les ingrédients perdus du Gâteau d'Harmonie et affronter les sbires sucrés du Chef de l'Ombre pour restaurer la paix et l'équilibre.
	
	Explorez des contrées sucrées, affrontez des monstres de confiserie, résolvez des énigmes pâtissières et découvrez des recettes secrètes pour relever le défi ultime : Dough or Die. Préparez-vous à un festin d'aventures où le destin de deux mondes repose entre vos couches crémeuses.
	`
	for _, letter := range introduction {
		fmt.Print(string(letter))
		time.Sleep(15 * time.Millisecond)

	}
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(5 * time.Second)

	fmt.Println(" ·▄▄▄▄        ▄• ▄▌ ▄▄ •  ▄ .▄          ▄▄▄      ·▄▄▄▄  ▪  ▄▄▄ .")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("██▪ ██ ▪     █▪██▌▐█ ▀ ▪██▪▐█    ▪     ▀▄ █·    ██▪ ██ ██ ▀▄.▀·")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("▐█· ▐█▌ ▄█▀▄ █▌▐█▌▄█ ▀█▄██▀▐█     ▄█▀▄ ▐▀▀▄     ▐█· ▐█▌▐█·▐▀▀▪▄")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("██. ██ ▐█▌.▐▌▐█▄█▌▐█▄▪▐███▌▐▀    ▐█▌.▐▌▐█•█▌    ██. ██ ▐█▌▐█▄▄▌")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("▀▀▀▀▀•  ▀█▄▀▪ ▀▀▀ ·▀▀▀▀ ▀▀▀ ·     ▀█▄▀▪.▀  ▀    ▀▀▀▀▀• ▀▀▀ ▀▀▀ ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("1 : Pour Créer Ta Patisserie Préférée !")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		p.CharCreation()
	default:
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		fmt.Scan(&choice)
	}
}
