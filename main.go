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
func RechercherContact(nom string) {
	if contact, ok := annuaire[nom]; ok {
		fmt.Printf("Contact trouvé : %s - %s\n", contact.Nom, contact.Tel)
	} else {
		fmt.Println("Contact non trouvé.")
	}
}

func AjouterContact(nom, tel string) {
	if _, existe := annuaire[nom]; existe {
		fmt.Println("Contact déjà existant.")
		return
	}
	annuaire[nom] = Contact{Nom: nom, Tel: tel}
	fmt.Println("Contact ajouté :", nom)
}

func SupprimerContact(nom string) {
	if _, ok := annuaire[nom]; ok {
		delete(annuaire, nom)
		fmt.Println("Contact supprimé :", nom)
	} else {
		fmt.Println("Contact introuvable.")
	}
}

func main() {
	annuaire["Hamza"] = Contact{Nom: "Hamza", Tel: "0601020303"}
	annuaire["Valentin"] = Contact{Nom: "Valentin", Tel: "0603040506"}
	annuaire["Serhat"] = Contact{Nom: "Serhat", Tel: "0602340406"}

	action := flag.String("action", "", "actions possible : ajouter, rechercher, lister, supprimer, modifier")
	nom := flag.String("nom", "", "Nom du contact")
	tel := flag.String("tel", "", "Numéro de téléphone du contact")
	flag.Parse()
	switch *action {
	case "ajouter":
		if *nom == "" || *tel == "" {
			fmt.Println("Nom et numéro de téléphone requis pour ajouter un contact.")
			return
		}
		AjouterContact(*nom, *tel)
		fmt.Println("Liste des contacts :")
		ListerContacts()
	case "lister":
		ListerContacts()
	case "rechercher":
		if *nom == "" {
			fmt.Println("Nom requis pour rechercher un contact.")
			return
		}
		RechercherContact(*nom)
	case "supprimer":
		if *nom == "" {
			fmt.Println("Nom requis pour supprimer un contact.")
			return
		}
		SupprimerContact(*nom)
		fmt.Println("Liste des contacts :")
		ListerContacts()
	default:
		fmt.Println("Action non reconnue. Utilisez --action avec : ajouter, rechercher, lister, supprimer, modifier")
	}
}
