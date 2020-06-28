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
	Deposits     []*Deposit
	Debtors      []*Debtor
	RealStates   []*RealState
	Vehicles     []*Vehicle
	Agricultural []*Agricultural
}

// Deposit describes money at a financial institution.
type Deposit struct {
	TipoEntidad string
	Entidad     string
	Tipo        string
	Pais        string
	Importe     int64
}

// Debtor describes a person that owns money to the official.
type Debtor struct {
	Nombre  string
	Clase   string
	Plazo   int
	Importe int64
}

// RealState is a real state owned by the official.
type RealState struct {
	Padron                 string
	Uso                    string
	Pais                   string
	Distrito               string
	Adquisicion            int
	TipoAdquisicion        string
	SuperficieTerreno      int64
	ValorTerreno           int64
	SuperficieConstruccion int64
	ValorConstruccion      int64
	Importe                int64
}

// Vehicle is a vehicle owned by the official.
type Vehicle struct {
	Tipo        string
	Marca       string
	Modelo      string
	Importe     int64
	Adquisicion int
	Fabricacion int
}

// Agricultural is an official's agricultural activity.
type Agricultural struct {
	Tipo      string
	Ubicacion string
	Especie   string
	Cantidad  int64
	Precio    int64
	Importe   int64
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

func (s *RealState) String() string {
	return fmt.Sprintf("Padron: %s\n"+
		"Uso: %s\n"+
		"Pais: %s\n"+
		"Distrito: %s\n"+
		"Adquisicion: %d\n"+
		"TipoAdquisicion: %s\n"+
		"SuperficieTerreno: %d\n"+
		"ValorTerreno: %d\n"+
		"SuperficieConstruccion: %d\n"+
		"ValorConstruccion: %d\n"+
		"Importe: %d\n",
		s.Padron, s.Uso, s.Pais, s.Distrito, s.Adquisicion, s.TipoAdquisicion,
		s.SuperficieTerreno, s.ValorTerreno, s.SuperficieConstruccion, s.ValorConstruccion,
		s.Importe)
}

func (v *Vehicle) String() string {
	return fmt.Sprintf("Tipo: %s\n"+
		"Marca: %s\n"+
		"Modelo: %s\n"+
		"Importe: %d\n"+
		"Adquisicion: %d\n"+
		"Fabricacion: %d\n",
		v.Tipo, v.Marca, v.Modelo, v.Importe, v.Adquisicion, v.Fabricacion)
}

func (a *Agricultural) String() string {
	return fmt.Sprintf("Tipo Actividad: %s\n"+
		"Ubicación: %s\n"+
		"Especie: %s\n"+
		"Cantidad: %d\n"+
		"Precio: %d\n"+
		"Importe: %d\n",
		a.Tipo, a.Ubicacion, a.Especie, a.Cantidad, a.Precio, a.Importe)
}
