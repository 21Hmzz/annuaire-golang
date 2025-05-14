package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

type Contact struct {
	Nom string
	Tel string
}

var annuaire = make(map[string]Contact)

func ChargerAnnuaire(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	var contacts []Contact
	err = json.Unmarshal(file, &contacts)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	for _, contact := range contacts {
		annuaire[contact.Nom] = contact
	}
	fmt.Println("Contacts chargés depuis le fichier", filename)
}
func SauvegarderAnnuaire(filename string) {
	contacts := make([]Contact, 0, len(annuaire))
	for _, contact := range annuaire {
		contacts = append(contacts, contact)
	}

	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de l'encodage JSON :", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
		return
	}

	fmt.Println("Annuaire sauvegardé dans le fichier", filename)
}

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
	SauvegarderAnnuaire("contacts.json")

	fmt.Println("Contact ajouté :", nom)
}

func SupprimerContact(nom string) {
	if _, ok := annuaire[nom]; ok {
		delete(annuaire, nom)
		fmt.Println("Contact supprimé :", nom)
	} else {
		fmt.Println("Contact introuvable.")
	}

	SauvegarderAnnuaire("contacts.json")

}

func ModifierContact(nom, nouveauTel string) {
	if _, ok := annuaire[nom]; ok {
		annuaire[nom] = Contact{Nom: nom, Tel: nouveauTel}
		fmt.Println("Contact modifié :", nom)
	} else {
		fmt.Println("Contact introuvable.")
	}
	SauvegarderAnnuaire("contacts.json")

}

func main() {
	file := "contacts.json"
	ChargerAnnuaire(file)

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
	case "modifier":
		if *nom == "" || *tel == "" {
			fmt.Println("Nom et numéro de téléphone requis pour modifier un contact.")
			return
		}
		ModifierContact(*nom, *tel)
		fmt.Println("Liste des contacts :")
		ListerContacts()
	default:
		fmt.Println("Action non reconnue. Utilisez --action avec : ajouter, rechercher, lister, supprimer, modifier")
	}
}
