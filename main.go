package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"os"
)

var (
	info    = color.New(color.FgCyan).SprintFunc()
	success = color.New(color.FgGreen).SprintFunc()
	warning = color.New(color.FgYellow).SprintFunc()
	errcol  = color.New(color.FgRed).SprintFunc()
)

// os exit ? return marche pas...
func menu() {
	items := []string{"Ajouter", "Lister", "Rechercher", "Modifier", "Supprimer", "Quitter"}
	for {
		prompt := promptui.Select{
			Label: "Choisissez une action",
			Items: items,
		}
		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Println(color.RedString("✘ Erreur de menu:"), err)
			return
		}

		switch items[idx] {
		case "Ajouter":
			var nom, tel string
			fmt.Print("Nom : ")
			fmt.Scanln(&nom)
			fmt.Print("Tel : ")
			fmt.Scanln(&tel)
			AjouterContact(nom, tel)
			os.Exit(0)

		case "Lister":
			ListerContacts()
			os.Exit(0)
		case "Rechercher":
			var nom string
			fmt.Print("Nom à rechercher : ")
			fmt.Scanln(&nom)
			RechercherContact(nom)
			os.Exit(0)

		case "Modifier":
			var nom, tel string
			fmt.Print("Nom à modifier : ")
			fmt.Scanln(&nom)
			fmt.Print("Nouveau tel : ")
			fmt.Scanln(&tel)
			ModifierContact(nom, tel)
			os.Exit(0)

		case "Supprimer":
			var nom string
			fmt.Print("Nom à supprimer : ")
			fmt.Scanln(&nom)
			SupprimerContact(nom)
			os.Exit(0)

		case "Quitter":
			fmt.Println(color.CyanString("👋 Au revoir !"))
			os.Exit(0)
		}
		fmt.Println()
	}
}

type Contact struct {
	Nom string
	Tel string
}

var annuaire = make(map[string]Contact)

func ChargerAnnuaire(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(errcol("Erreur lors de la lecture du fichier:"), err)
		return
	}

	var contacts []Contact
	if err := json.Unmarshal(file, &contacts); err != nil {
		fmt.Println(errcol("Erreur lors du décodage JSON:"), err)
		return
	}

	for _, contact := range contacts {
		annuaire[contact.Nom] = contact
	}
	fmt.Println(success("✔ Contacts chargés depuis le fichier"), info(filename))
}
func SauvegarderAnnuaire(filename string) {
	contacts := make([]Contact, 0, len(annuaire))
	for _, contact := range annuaire {
		contacts = append(contacts, contact)
	}

	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		fmt.Println(errcol("Erreur lors de l'encodage JSON:"), err)
		return
	}

	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		fmt.Println(errcol("Erreur lors de l'écriture du fichier:"), err)
		return
	}

	fmt.Println(success("✔ Annuaire sauvegardé dans"), info(filename))
}

func ListerContacts() {
	if len(annuaire) == 0 {
		fmt.Println(warning("⚠ Annuaire vide."))
		return
	}
	for _, c := range annuaire {
		fmt.Printf("  - %s : %s\n", color.New(color.Bold).Sprint(c.Nom), color.New(color.Italic).Sprint(c.Tel))
	}
}

func RechercherContact(nom string) {
	if contact, ok := annuaire[nom]; ok {
		fmt.Printf("%s Contact trouvé : %s - %s\n",
			success("✔"),
			color.New(color.Bold).Sprint(contact.Nom),
			color.New(color.Underline).Sprint(contact.Tel),
		)
	} else {
		fmt.Println(errcol("✘ Contact non trouvé."))
	}
}

func AjouterContact(nom, tel string) {
	if _, existe := annuaire[nom]; existe {
		fmt.Println(errcol("✘ Ce contact existe déjà."))
		return
	}
	annuaire[nom] = Contact{Nom: nom, Tel: tel}
	SauvegarderAnnuaire("contacts.json")

	fmt.Println(success("✔ Contact ajouté :"), info(nom))
}

func SupprimerContact(nom string) {
	if _, ok := annuaire[nom]; ok {
		delete(annuaire, nom)
		fmt.Println(success("✔ Contact supprimé :"), info(nom))
	} else {
		fmt.Println(errcol("✘ Contact introuvable."))
	}

	SauvegarderAnnuaire("contacts.json")

}

func ModifierContact(nom, nouveauTel string) {
	if _, ok := annuaire[nom]; ok {
		annuaire[nom] = Contact{Nom: nom, Tel: nouveauTel}
		fmt.Println(success("✔ Contact modifié :"), info(nom))
	} else {
		fmt.Println(errcol("✘ Contact introuvable."))
	}
	SauvegarderAnnuaire("contacts.json")

}

func Help() {
	fmt.Println(warning("⚠ Aide :"))
	fmt.Println(info("  - ajouter : Ajouter un contact"))
	fmt.Println(info("  - rechercher : Rechercher un contact"))
	fmt.Println(info("  - lister : Lister tous les contacts"))
	fmt.Println(errcol("  - supprimer : Supprimer un contact"))
	fmt.Println(warning("  - modifier : Modifier un contact"))
}

func main() {

	action := flag.String("action", "", "actions possible : ajouter, rechercher, lister, supprimer, modifier")
	nom := flag.String("nom", "", "Nom du contact")
	tel := flag.String("tel", "", "Numéro de téléphone du contact")
	help := flag.Bool("help", false, "Afficher l'aide")
	flag.Parse()

	if *help {
		Help()
		return
	}

	file := "contacts.json"
	ChargerAnnuaire(file)
	switch *action {
	case "ajouter":
		if *nom == "" || *tel == "" {
			fmt.Println(info("Nom et numéro de téléphone requis pour ajouter un contact."))
			return
		}
		AjouterContact(*nom, *tel)
	case "lister":
		ListerContacts()
	case "rechercher":
		if *nom == "" {
			fmt.Println(errcol("Nom requis pour rechercher un contact."))
			return
		}
		RechercherContact(*nom)
	case "supprimer":
		if *nom == "" {
			fmt.Println(errcol("Nom requis pour supprimer un contact."))
			return
		}
		SupprimerContact(*nom)
	case "modifier":
		if *nom == "" || *tel == "" {
			fmt.Println(errcol("Nom et numéro de téléphone requis pour modifier un contact."))
			return
		}
		ModifierContact(*nom, *tel)

	default:
		fmt.Println(errcol("✘ Aucune action spécifiée."))
		menu()
		return
	}
}
