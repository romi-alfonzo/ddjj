package declaration

import (
	"fmt"
	"time"
)

// Declaration is the data on a public official's declaraion
type Declaration struct {
	Fecha       time.Time `json:"fecha" bson:"fecha"`
	Cedula      int       `json:"cedula" bson:"cedula"`
	Nombre      string    `json:"nombre" bson:"nombre"`
	Apellido    string    `json:"appellido" bson:"appellido"`
	Cargo       string    `json:"cargo" bson:"cargo"`
	Institucion string    `json:"institucion" bson:"institucion"`

	// Activos
	Deposits     []*Deposit      `json:"depositos" bson:"depositos"`
	Debtors      []*Debtor       `json:"deudores" bson:"deudores"`
	RealStates   []*RealState    `json:"inmuebles" bson:"inmuebles"`
	Vehicles     []*Vehicle      `json:"vehiculos" bson:"vehiculos"`
	Agricultural []*Agricultural `json:"actividadesAgropecuarias" bson:"actividadesAgropecuarias"`
	Furniture    []*Furniture    `json:"muebles" bson:"muebles"`
	OtherAssets  []*OtherAsset   `json:"otrosActivos" bson:"otrosActivos"`

	Debts []*Debt `json:"deudas" bson:"deudas"`

	Assets       int64 `json:"activos" bson:"activos"`
	Liabilities  int64 `json:"pasivos" bson:"pasivos"`
	NetPatrimony int64 `json:"patrimonioNeto" bson:"patrimonioNeto"`
}

// CalculatePatrimony adds up assets and debts.
func (d *Declaration) CalculatePatrimony() int64 {

	var debts int64
	for _, v := range d.Debts {
		debts += v.Saldo
	}

	d.Assets = d.AddAssets()
	d.Liabilities = debts
	d.NetPatrimony = d.Assets - d.Liabilities

	return d.NetPatrimony
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
	TipoEntidad string `json:"tipoEntidad" bson:"tipoEntidad"`
	Entidad     string `json:"entidad" bson:"entidad"`
	Tipo        string `json:"tipo" bson:"tipo"`
	Pais        string `json:"pais" bson:"pais"`
	Importe     int64  `json:"importe" bson:"importe"`
}

// Debtor describes a person that owns money to the official.
type Debtor struct {
	Nombre  string `json:"nombre" bson:"nombre"`
	Clase   string `json:"clase" bson:"clase"`
	Plazo   int    `json:"plazo" bson:"plazo"`
	Importe int64  `json:"importe" bson:"importe"`
}

// RealState is a real state owned by the official.
type RealState struct {
	Padron                 string `json:"padron" bson:"padron"`
	Uso                    string `json:"uso" bson:"uso"`
	Pais                   string `json:"pais" bson:"pais"`
	Distrito               string `json:"distrito" bson:"distrito"`
	Adquisicion            int    `json:"adquisicion" bson:"adquisicion"`
	TipoAdquisicion        string `json:"tipoAdquisicion" bson:"tipoAdquisicion"`
	SuperficieTerreno      int64  `json:"superficieTerreno" bson:"superficieTerreno"`
	ValorTerreno           int64  `json:"valorTerreno" bson:"valorTerreno"`
	SuperficieConstruccion int64  `json:"superficieConstruccion" bson:"superficieConstruccion"`
	ValorConstruccion      int64  `json:"valorConstruccion" bson:"valorConstruccion"`
	Importe                int64  `json:"importe" bson:"importe"`
}

// Vehicle is a vehicle owned by the official.
type Vehicle struct {
	Tipo        string `json:"tipo" bson:"tipo"`
	Marca       string `json:"marca" bson:"marca"`
	Modelo      string `json:"modelo" bson:"modelo"`
	Adquisicion int    `json:"adquisicion" bson:"adquisicion"`
	Fabricacion int    `json:"fabricacion" bson:"fabricacion"`
	Importe     int64  `json:"importe" bson:"importe"`
}

// Agricultural is an official's agricultural activity.
type Agricultural struct {
	Tipo      string `json:"tipo" bson:"tipo"`
	Ubicacion string `json:"ubicacion" bson:"ubicacion"`
	Especie   string `json:"especie" bson:"especie"`
	Cantidad  int64  `json:"cantidad" bson:"cantidad"`
	Precio    int64  `json:"precio" bson:"precio"`
	Importe   int64  `json:"importe" bson:"importe"`
}

// Furniture is a furniture owned by the official.
type Furniture struct {
	Tipo    string `json:"tipo" bson:"tipo"`
	Importe int64  `json:"importe" bson:"importe"`
}

// OtherAsset is another asset not included in other fields.
type OtherAsset struct {
	Descripcion string `json:"descripcion" bson:"descripcion"`
	Empresa     string `json:"empresa" bson:"empresa"`
	RUC         string `json:"ruc" bson:"ruc"`
	Pais        string `json:"pais" bson:"pais"`
	Cantidad    int64  `json:"cantidad" bson:"cantidad"`
	Precio      int64  `json:"precio" bson:"precio"`
	Importe     int64  `json:"importe" bson:"importe"`
}

// Debt is money the official owes to others.
type Debt struct {
	Tipo    string `json:"tipo" bson:"tipo"`
	Empresa string `json:"empresa" bson:"empresa"`
	Plazo   int    `json:"plazo" bson:"plazo"`
	Cuota   int64  `json:"cuota" bson:"cuota"`
	Total   int64  `json:"total" bson:"total"`
	Saldo   int64  `json:"saldo" bson:"saldo"`
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
