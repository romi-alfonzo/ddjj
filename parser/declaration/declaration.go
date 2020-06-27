package declaration

import (
	"fmt"
)

// Declaration is the data on a public official's declaraion
type Declaration struct {
	Ano         int
	Cedula      int
	Nombre      string
	Apellido    string
	Funcion     string
	Institucion string

	// Activos
	Deposits []*Deposit
	Debtors  []*Debtor
}

// Deposit describes money at a financial institution.
type Deposit struct {
	TipoEntidad string
	Entidad     string
	Tipo        string
	Pais        string
	Importe     int64
}

// Debtor describe a person that owns money to the official.
type Debtor struct {
	Nombre  string
	Clase   string
	Plazo   int
	Importe int64
}

func (d *Deposit) String() string {
	return fmt.Sprintf("Tipo Entidad: %s\n"+
		"Entidad: %s\n"+
		"Tipo: %s\n"+
		"Pais: %s\n"+
		"Importe: %d\n",
		d.TipoEntidad, d.Entidad, d.Tipo, d.Pais, d.Importe)
}

func (d *Debtor) String() string {
	return fmt.Sprintf("Nombre: %s\n"+
		"Clase: %s\n"+
		"Plazo: %d\n"+
		"Importe: %d\n",
		d.Nombre, d.Clase, d.Plazo, d.Importe)
}
