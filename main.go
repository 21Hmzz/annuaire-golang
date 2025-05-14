package main

import (
	"flag"
	"fmt"
)

type Contact struct {
	Nom string
	Tel string
}

var annuaire = make(map[string]Contact)

func ListerContacts() {
	if len(annuaire) == 0 {
		fmt.Println("Annuaire vide.")
		return
	}
	for _, c := range annuaire {
		fmt.Printf("- %s : %s\n", c.Nom, c.Tel)
	}
}

func main() {
	annuaire["Hamza"] = Contact{Nom: "Hamza", Tel: "0601020303"}
	annuaire["Valentin"] = Contact{Nom: "Valentin", Tel: "0603040506"}
	annuaire["Serhat"] = Contact{Nom: "Serhat", Tel: "0602340406"}

	action := flag.String("action", "", "actions possible : ajouter, rechercher, lister, supprimer, modifier")
	flag.Parse()
	switch *action {
	case "lister":
		ListerContacts()
	default:
		fmt.Println("Action non reconnue. Utilisez --action avec : ajouter, rechercher, lister, supprimer, modifier")
	}
}
