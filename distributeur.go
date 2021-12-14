package main

import (
	"bufio"
	_ "context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

const name = "Le distributeur"
const bonjour = "Bonjour ! Je suis "

type Distributeur struct {
	boissons map[string]int
	Name     string
	Bonjour  string
	GetInput string
	Cmds     []Cmd
	Mtnce    bool
}

type Cmd struct {
	Drink    string
	Quantity int
	Date     time.Time
}

func main() {

	distributeur := Distributeur{
		Mtnce: false,
		boissons: map[string]int{
			"Riri": 100,
			"Café": 5,
			"Thé":  3,
			"Eau":  100},
		Name:     name,
		Bonjour:  bonjour,
		Cmds:     []Cmd{},
		GetInput: "Quel boisson souhaitez vous ?",
	}

	distributeur.Greetings()

	for {
		if distributeur.Mtnce {
			distributeur.maintenance()
			continue
		}
		distributeur.offer()
		boisson := distributeur.getUserInput()
		if boisson == "quit" {
			fmt.Println("Good bye!")
			break
		}
		if boisson == "historique" {
			for index, cmd := range distributeur.Cmds {
				fmt.Printf("#%d: %s %d %v", index, cmd.Drink, cmd.Quantity, cmd.Date)
			}
			continue
		}
		if boisson == "maintenance" {
			fmt.Println("Maintenance mode activated")
			distributeur.Mtnce = true
			continue
		}
		distributeur.serve(boisson)

	}
}

func (d Distributeur) Greetings() {
	message := d.Bonjour + " " + d.Name
	fmt.Printf("%s \n", message)
}

func (d Distributeur) offer() {
	fmt.Printf("Voici la liste des boissons\n")
	for boisson, stock := range d.boissons {
		fmt.Printf("* %s: (%d)\n", boisson, stock)
	}
}

func (d Distributeur) getUserInput() string {
	fmt.Printf(d.GetInput + " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
func (d *Distributeur) serve(boisson string) {
	number, ok := d.boissons[boisson]
	if ok == false {
		fmt.Printf("Cette boisson n'existe pas\n")

		return
	}

	if number <= 0 {
		fmt.Printf("Il n'y a plus de %s disponible\n", boisson)

		return
	}
	d.boissons[boisson]--
	d.addCmd(boisson, 1)

	fmt.Printf("Voici votre %s, il reste %d boisson(s) \n", boisson, d.boissons[boisson])
}
func (d *Distributeur) addCmd(boisson string, quantity int) {
	d.Cmds = append(d.Cmds, Cmd{
		Drink:    boisson,
		Quantity: quantity,
		Date:     time.Now(),
	})
}

func (d *Distributeur) maintenance() {
	for {
		fmt.Println("Souhaitez vous export les données ou faire le réassort ? ")
		boisson := d.getUserInput()
		if boisson == "export" {
			file, err := os.Create("export.csv")
			if err != nil {
				log.Fatalf("Error %v", err)
			}

			defer file.Close()

			writer := csv.NewWriter(file)

			defer writer.Flush()
			break
		}

		fmt.Println("Modifier une boisson:")
		for boisson, stock := range d.boissons {
			fmt.Printf("* %s: (%d)\n", boisson, stock)
		}
		boisson = d.getUserInput()
		if boisson == "running" {
			d.Mtnce = false
			break
		}

		_, ok := d.boissons[boisson]
		if !ok {
			continue
		}
		addStock := 0
		fmt.Println("Combien:")
		fmt.Scanln(&addStock)
		d.boissons[boisson] = d.boissons[boisson] + addStock

	}
}
}
