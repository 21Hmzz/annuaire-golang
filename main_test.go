package main

import (
	"fmt"
	"testing"
)

func initAnnuaire() {
	annuaire = make(map[string]Contact)
	filename := "contacts.json"
	ChargerAnnuaire(filename)
}

func TestAjoutDoublon(t *testing.T) {
	initAnnuaire()
	AjouterContact("Lukhas", "0707070707")
	AjouterContact("Lukhas", "0808080808")
	fmt.Println("Liste des contacts après ajout :")
	ListerContacts()

	contact, ok := annuaire["Lukhas"]
	if !ok {
		t.Error("Lukhas n'a pas été ajouté")
	}
	if contact.Tel != "0707070707" {
		t.Errorf("Numéro incorrect pour Lukhas. Attendu: 0707070707, Obtenu: %s", contact.Tel)
	}
}

func TestAjouterContact(t *testing.T) {
	initAnnuaire()
	AjouterContact("Lukhas", "0707070707")
	fmt.Println("Liste des contacts après ajout :")
	ListerContacts()

	contact, ok := annuaire["Lukhas"]
	if !ok {
		t.Error("Lukhas n'a pas été ajouté")
	}
	if contact.Tel != "0707070707" {
		t.Errorf("Numéro incorrect pour Lukhas. Attendu: 0707070707, Obtenu: %s", contact.Tel)
	}
}

func TestModifierContact(t *testing.T) {
	initAnnuaire()

	AjouterContact("Serhat", "0505050505")

	ModifierContact("Serhat", "0600000000")
	fmt.Println("Liste des contacts après modification :")
	ListerContacts()
	contact := annuaire["Serhat"]
	if contact.Tel != "0600000000" {
		t.Errorf("Modification échouée. Attendu: 0600000000, Obtenu: %s", contact.Tel)
	}
}

func TestSupprimerContact(t *testing.T) {
	initAnnuaire()

	AjouterContact("Valentin", "0808080808")

	SupprimerContact("Valentin")
	fmt.Println("Liste des contacts après suppression :")
	ListerContacts()
	if _, ok := annuaire["Valentin"]; ok {
		t.Error("Valentin aurait dû être supprimée")
	}
}
