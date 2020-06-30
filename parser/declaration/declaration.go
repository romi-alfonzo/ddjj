package declaration

import (
	"fmt"
	"time"
)

// Declaration is the data on a public official's declaraion
type Declaration struct {
	Date        time.Time
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
	Furniture    []*Furniture
	OtherAssets  []*OtherAsset

	Debts []*Debt
}

// Net returns the patrimony's net value.
func (d *Declaration) Net() int64 {
	var debts int64
	for _, v := range d.Debts {
		debts += v.Saldo
	}

	return d.AddAssets() - debts
}

// AddAssets adds all the assets.
func (d *Declaration) AddAssets() int64 {
	var total int64

	for _, v := range d.Deposits {
		total += v.Importe
	}
	for _, v := range d.Debtors {
		total += v.Importe
	}
	for _, v := range d.RealStates {
		total += v.Importe
	}
	for _, v := range d.Vehicles {
		total += v.Importe
	}
	for _, v := range d.Agricultural {
		total += v.Importe
	}
	for _, v := range d.Furniture {
		total += v.Importe
	}
	for _, v := range d.OtherAssets {
		total += v.Importe
	}

	return total
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

// Furniture is a furniture owned by the official.
type Furniture struct {
	Tipo    string
	Importe int64
}

// OtherAsset is another asset not included in other fields.
type OtherAsset struct {
	Descripcion string
	Empresa     string
	RUC         string
	Pais        string
	Cantidad    int64
	Precio      int64
	Importe     int64
}

// Debt is money the official owes to others.
type Debt struct {
	Tipo    string
	Empresa string
	Plazo   int
	Cuota   int64
	Total   int64
	Saldo   int64
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

func (f *Furniture) String() string {
	return fmt.Sprintf("Tipo: %s\nImporte:%d\n", f.Tipo, f.Importe)
}

func (a *OtherAsset) String() string {
	return fmt.Sprintf("Descripcion: %s\n"+
		"Empresa: %s\n"+
		"RUC: %s\n"+
		"Pais: %s\n"+
		"Cantidad: %d\n"+
		"Precio: %d\n"+
		"Importe: %d\n",
		a.Descripcion, a.Empresa, a.RUC, a.Pais, a.Cantidad, a.Precio, a.Importe)
}

func (d *Debt) String() string {
	return fmt.Sprintf("Tipo: %s\n"+
		"Empresa: %s\n"+
		"Plazo: %d\n"+
		"Cuota: %d\n"+
		"Total: %d\n"+
		"Saldo: %d\n",
		d.Tipo, d.Empresa, d.Plazo, d.Cuota, d.Total, d.Saldo)
}
