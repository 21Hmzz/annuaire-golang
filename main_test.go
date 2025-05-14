package main

func init() {
	annuaire = make(map[string]Contact)
	annuaire["Hamza"] = Contact{Nom: "Hamza", Tel: "0601020303"}
	annuaire["Valentin"] = Contact{Nom: "Valentin", Tel: "0603040506"}
	annuaire["Serhat"] = Contact{Nom: "Serhat", Tel: "0602340406"}
}
