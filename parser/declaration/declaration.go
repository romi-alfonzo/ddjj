package declaration

// Declaration is the data on a public official's declaraion
type Declaration struct {
	Ano         int
	Cedula      int
	Nombre      string
	Apellido    string
	Funcion     string
	Institucion string
}

// Debtor describe a person that owns money to the official.
type Debtor struct {
	Nombre  string
	Clase   string
	Plazo   int
	Importe int64
	Obs     string
}
