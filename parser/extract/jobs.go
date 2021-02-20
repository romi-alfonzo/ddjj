package extract

import (
	"fmt"
	"strings"
	"github.com/InstIDEA/ddjj/parser/declaration"	
)

// TODO
// create a function to try sort inst-title pairs

/*
analysis data

raw data:

linea:  57 ; INTITUCION:
linea:  64 ; CARGO:   
linea:  66 ; i->POLICIA NACIONAL
linea:  70 ; INTITUCION:  
linea:  71 ; c->RESGUARDO POLICIAL
linea:  79 ; CARGO: 
linea:  79 ; c->SUB OFICIAL SEGUNDO
linea:  82 ; i->MINISTERIO DE TRABAJO, EMPLEO Y SEGURIDAD SOCIAL

parser output:

		"instituciones": [
			{
				"cargo": "RESGUARDO POLICIAL",
				"institucion": "POLICIA NACIONAL"
			},
			{
				"cargo": "SUB OFICIAL SEGUNDO",
				"institucion": "MINISTERIO DE TRABAJO, EMPLEO Y SEGURIDAD SOCIAL"
			}
		]

expected order:

(SUB OFICIAL SEGUNDO, POLICIA NACIONAL)
(RESGUARDO POLICIAL, MINISTERIO DE TRABAJO, EMPLEO Y SEGURIDAD SOCIAL)
*/

var instsCache = map[string]bool{}

func Jobs(e *Extractor, parser *ParserData) []*declaration.Job {

	var instituciones []*declaration.Job
	var resultsPositions []int // for valid and invalid results
	var counter = countJobs(e)
	var successful int

	instsCache = getInstsCache()

	e.Rewind()
	e.BindFlag(EXTRACTOR_FLAG_1)

	job := &declaration.Job{ }

	if counter > 0 &&
	e.MoveUntilStartWith(CurrToken, "DATOS LABORALES") {
		e.SaveLine()

		for e.Scan() {
			if counter == successful {
				break
			}

			if job.Cargo == "" {
				value := getJobTitle(e, &resultsPositions)

				if isValidJobTitle(value) {
					job.Cargo = value
				}
			}

			if job.Cargo != "" &&
			job.Institucion == "" {
				value := getJobInst(e, &resultsPositions)

				if isValidJobInst(value) {
					job.Institucion = value
					instsCache[value] = true
				}
			}

			if job.Cargo != "" && job.Institucion != "" {
				successful++
				instituciones = append(instituciones, job)
				job = &declaration.Job{ }
				e.MoveUntilSavedLine()
			}
		}
	}

	if successful != counter {
		parser.addMessage(fmt.Sprintf("ignored jobs: %d/%d", counter - successful, counter))
	}

	if instituciones == nil {
		parser.addError(fmt.Errorf("failed when extracting jobs"))
		return nil
	}

	return instituciones
}

func getJobTitle(e *Extractor, pos *[]int) string {
	// control positions are based on first matching tokens

	var value string

	if strings.Contains(e.CurrToken, "CARGO") &&
	!ContainsIntItem(*pos, e.CurrLineNum()) {
		val, check := isKeyValuePair(e.CurrToken, "CARGO")
		if check {
			*pos = append(*pos, e.CurrLineNum())
			e.MoveUntilSavedLine()
			return val
		}
	}

	if isCurrLine(e.CurrToken, "INSTITUCIÓN") {
		if !ContainsIntItem(*pos, e.PrevLineNum()) {
			value = e.PrevToken
			*pos = append(*pos, e.PrevLineNum())
			e.MoveUntilSavedLine()
			return value
		}

		if !ContainsIntItem(*pos, e.NextLineNum()) {
			value = e.NextToken
			*pos = append(*pos, e.NextLineNum())
			e.MoveUntilSavedLine()
			return value
		}
	}

	if isCurrLine(e.PrevToken, "TELÉFONO") &&
	!isCurrLine(e.NextToken, "CARGO") &&
	!ContainsIntItem(*pos, e.CurrLineNum()) {
		value = e.CurrToken
		*pos = append(*pos, e.CurrLineNum())
		e.MoveUntilSavedLine()
		return value
	}

	return ""
}

func getJobInst(e *Extractor, pos *[]int) string {
	var value string

	if instsCache[removeAccents(e.CurrToken)] &&
	!ContainsIntItem(*pos, e.CurrLineNum()) {
		value = e.CurrToken
		*pos = append(*pos, e.CurrLineNum())
		e.MoveUntilSavedLine()
		return value
	}

	// control positions are based on first matching tokens

	if isCurrLine(e.PrevToken, "INSTITUCIÓN") &&
	isCurrLine(e.CurrToken, "TIPO") &&
	!ContainsIntItem(*pos, e.NextLineNum()) {
		value = e.NextToken
		*pos = append(*pos, e.NextLineNum())
		e.MoveUntilSavedLine()
		return value
	}

	if isCurrLine(e.PrevToken, "DIRECCIÓN") {
		_, check := isKeyValuePair(e.PrevToken, "DIRECCIÓN")

		if check &&
		!ContainsIntItem(*pos, e.NextLineNum()) {
			value = e.NextToken
			*pos = append(*pos, e.NextLineNum())
			e.MoveUntilSavedLine()
			return value
		}

		if !ContainsIntItem(*pos, e.CurrLineNum()) {
			value = e.CurrToken
			*pos = append(*pos, e.CurrLineNum())
			e.MoveUntilSavedLine()
			return value
		}
	}

	if isCurrLine(e.CurrToken, "FECHA ASUNC") &&
	!ContainsIntItem(*pos, e.NextLineNum()) {
		value = e.NextToken
		*pos = append(*pos, e.NextLineNum())
		e.MoveUntilSavedLine()
		return value
	}

	if isCurrLine(e.PrevToken, "COMISIONADO") &&
	isCurrLine(e.NextToken, "DEPENDENCIA") &&
	!ContainsIntItem(*pos, e.CurrLineNum()) {
		value = e.CurrToken
		*pos = append(*pos, e.CurrLineNum())
		e.MoveUntilSavedLine()	
		return value
	}

	if isCurrLine(e.NextToken, "CARGO") &&
	isCurrLine(e.CurrToken, "NOMBRADO/CONTRATADO") &&
	!ContainsIntItem(*pos, e.PrevLineNum()) {
		value = e.PrevToken
		*pos = append(*pos, e.PrevLineNum())
		e.MoveUntilSavedLine()		
		return value
	}

	if isCurrLine(e.NextToken, "COMISIONADO") &&
	!ContainsIntItem(*pos, e.PrevLineNum()) {
		value = e.PrevToken
		*pos = append(*pos, e.PrevLineNum())
		e.MoveUntilSavedLine()	
		return value
	}	

	return ""
}

func countJobs(e *Extractor) int {
	var counter int

	for e.Scan() {
		// first position
		if isCurrLine(e.CurrToken, "CARGO") {
			counter++
			continue
		}

		// middle position
		if hasLeadingSpaces(e.CurrToken, "CARGO") &&
		!endsWith(e.CurrToken, "CARGO") {
			counter++
		}
	}
	return counter
}

func isValidJobTitle(s string) bool {
	if len(s) < 5 { // minimum length for job title
		return false
	}

	if isDate(s) {
		return false
	}

	if isNumber(s) {
		return false
	}

	if isPhoneNumber(s) {
		return false
	}

	if isJobFormField(s) {
		return false
	}

	if isJobFormCommonAnswer(s) {
		return false
	}
	
	return true
}

func isValidJobInst(s string) bool {

	if containsKWOfInsts(s) {
		return true
	}

	if len(s) < 5 { // minimum length for institute
		return false
	}

	if isDate(s) {
		return false
	}

	if isNumber(s) {
		return false
	}

	if isPhoneNumber(s) {
		return false
	}

	if isAddressStreet(s) {
		return false
	}

	if isJobFormField(s) {
		return false
	}

	if isJobFormCommonAnswer(s) {
		return false
	}

	return true
}

func getInstsCache() map[string]bool {
	// last update 18/02/2020
	// from https://datos.sfp.gov.py/data/oee and parser results
	institutions := [200]string{ "ADMINISTRACION NACIONAL DE ELECTRICIDAD", "ADMINISTRACION NACIONAL DE NAVEGACION Y PUERTOS", "AGENCIA FINANCIERA DE DESARROLO", 
	"AGENCIA NACIONAL DE EVALUACION Y ACREDITACION DE LA EDUCACION", "AGENCIA NACIONAL DE TRANSITO Y SEGURIDAD VIAL", "ARMADA NACIONAL",
	"AUDITORIA GENERAL DEL PODER EJECUTIVO", "AUTORIDAD REGULADORA RADIOLOGICA Y NUCLEAR", "BANCO CENTRAL DEL PARAGUAY", "BANCO NACIONAL DE FOMENTO",
	"CAJA DE JUBILACIONES Y PENSIONES DE EMPLEADOS BANCARIOS", "CAJA DE JUBILACIONES Y PENSIONES DEL PENSONAL DE BANCOS Y AFINES", 
	"CAJA DE JUBILACIONES Y PENSIONES DEL PERSONAL DE LA ANDE", "CAJA DE JUBILACIONES Y PENSIONES DEL PERSONAL MUNICIPAL", 
	"CAJA DE PRESTAMOS DEL MINISTERIO DE DEFENSA NACIONAL", "CAJA DE SEGURIDAD SOCIAL DE EMPLEADOS Y OBREROS FERROVIARIOS", "CAMARA DE DIPUTADOS", 
	"CAMARA DE SENADORES", "CAÑAS PARAGUAYAS", "COMANDO DE LAS FUERZAS MILITARES", "COMISION NACIONAL DE LA COMPETENCIA", 
	"COMISION NACIONAL DE TELECOMUNICACIONES", "COMISION NACIONAL DE VALORES", "COMPAÑIA PARAGUAYA DE COMUNICACIONES", "CONGRESO NACIONAL", 
	"CONSEJO DE LA MAGISTRATURA", "CONSEJO NACIONAL DE CIENCIA Y TECNOLOGIA", "CONSEJO NACIONAL DE EDUCACION SUPERIOR", "CONTRALORIA GENERAL DE LA REPUBLICA",
	"CORTE SUPREMA DE JUSTICIA", "CREDITO AGRICOLA DE HABILITACION", "DEFENSORIA DEL PUEBLO", "DIRECCION DE BENEFICENCIA Y AYUDA SOCIAL", "DIRECCION DE CONTRATACIONES",
	"DIRECCION GENERAL DE ESTADISTICA, ENCUESTA Y CENSO", "DIRECCION NACIONAL DE ADUANAS", "DIRECCION NACIONAL DE AERONAUTICA CIVIL", "DIRECCION NACIONAL DE BENEFICENCIA",
	"DIRECCION NACIONAL DE CONTRATACIONES PUBLICAS", "DIRECCION NACIONAL DE CORREOS DEL PARAGUAY", "DIRECCION NACIONAL DE DEFENSA, SALUD Y BIENESTAR ANIMAL",
	"DIRECCION NACIONAL DEL REGISTRO DEL ESTADO CIVIL DE LAS PERSONAS", "DIRECCION NACIONAL DE PROPIEDAD INTELECTUAL", "DIRECCION NACIONAL DE TRANSPORTE",
	"EMPRESA DE SERVICIOS SANITARIOS DEL PARAGUAY", "ENTE REGULADOR DE SERVICIOS SANITARIOS", "FERROCARRILES DEL PARAGUAY", "FONDO GANADERO", 
	"FONDO NACIONAL DE LA CULTURA Y LAS ARTES", "GOBERNACION DE ALTO PARAGUAY", "GOBERNACION DE ALTO PARANA", "GOBERNACION DE AMAMBAY", "GOBERNACION DE BOQUERON",
 	"GOBERNACION DE CAAGUAZU", "GOBERNACION DE CAAZAPA", "GOBERNACION DE CANINDEYU", "GOBERNACION DE CONCEPCION", "GOBERNACION DE CORDILLERA", "GOBERNACION DE GUAIRA", 
 	"GOBERNACION DE ITAPUA", "GOBERNACION DEL GUAIRA", "GOBERNACION DE MISIONES", "GOBERNACION DE ÑEEMBUCU", "GOBERNACION DE PARAGUARI", "GOBERNACION DE PRESIDENTE HAYES",
 	"GOBERNACION DE SAN PEDRO", "GOBIERNO DEPARTAMENTAL DE CENTRAL", "GOBIERNO DEPARTAMENTAL DE PARAGUARI", "HOSPITAL DE EMERGENCIAS MEDICAS",
 	"HOSPITAL REGIONAL SALTO DEL GUAIRA", "INDUSTRIA NACIONAL DEL CEMENTO", "INSTITUTO DE PREVISION SOCIAL", "INSTITUTO FORESTAL NACIONAL",
 	"INSTITUTO NACIONAL DE ALIMENTACION Y NUTRICION", "INSTITUTO NACIONAL DE COOPERATIVISMO", "INSTITUTO NACIONAL DE DESARROLLO RURAL Y DE LA TIERRA",
 	"INSTITUTO NACIONAL DE EDUCACION SUPERIOR DR. RAUL PEÑA", "INSTITUTO NACIONAL DEL COOPERATIVISMO", "INSTITUTO NACIONAL DEL INDIGENA", "INSTITUTO NACIONAL DE TECNOLOGIA",
 	"INSTITUTO PARAGUAYO DE ARTESANI", "INSTITUTO PARAGUAYO DEL INDIGENA", "INSTITUTO PARAGUAYO DE TECNOLOGIA AGRARIA", "INSTITUTO SUPERIOR DE BELLAS ARTES", "ITAIPU",
 	"JURADO DE ENJUICIAMIENTO DE MAGISTRADOS", "MECANISMO NACIONAL DE PREVENSION CONTRA LA TORTURA", "MINISTERIO DE AGRICULTURA Y GANADERIA", "MINISTERIO DE DEFENSA NACIONAL",
 	"MINISTERIO DE DESARROLLO SOCIAL", "MINISTERIO DE EDUCACION Y CULTURA", "MINISTERIO DE HACIENDA", "MINISTERIO DE INDUSTRIA Y COMERCIO", "MINISTERIO DE JUSTICIA",
 	"MINISTERIO DE LA DEFENSA PUBLICA", "MINISTERIO DEL AMBIENTE Y DESARROLLO SOSTENIBLE", "MINISTERIO DE LA MUJER", "MINISTERIO DE LA NIÑEZ Y ADOLESCENCIA",
 	"MINISTERIO DEL INTERIOR", "MINISTERIO DEL TRABAJO, EMPLEO Y SEGURIDAD SOCIAL", "MINISTERIO DE OBRAS PUBLICAS Y COMUNICACIONES", "MINISTERIO DE RELACIONES EXTERIORES",
 	"MINISTERIO DE SALUD PUBLICA Y BIENESTAR SOCIAL", "MINISTERIO DE TECNOLOGIAS DE LA INFORMACION Y COMUNICACION", "MINISTERIO DE TRABAJO, EMPLEO Y SEGURIDAD SOCIAL",
 	"MINISTERIO DE URBANISMO, VIVIENDA Y HABITAT", "MINISTERIO PUBLICO", "MINISTERIO PUBLICO FISCALIA GENERAL DEL ESTADO", "MUNICIPALIDAD DE 3 DE MAYO",
	"MUNICIPALIDAD DE ABAI", "MUNICIPALIDAD DE ASUNCION", "MUNICIPALIDAD DE CAACUPE", "MUNICIPALIDAD DE CAAPUCU", "MUNICIPALIDAD DE CAMBYRETA", "MUNICIPALIDAD DE CAPIATA",
	"MUNICIPALIDAD DE CAPITAN BADO", "MUNICIPALIDAD DE CARAPEGUA", "MUNICIPALIDAD DE CIUDAD DEL ESTE", "MUNICIPALIDAD DE CORONEL BOGADO", "MUNICIPALIDAD DE DR. JUAN LEON MALLORQUIN",
	"MUNICIPALIDAD DE EDELIRA", "MUNICIPALIDAD DE ENCARNACION", "MUNICIPALIDAD DE FERNANDO DE LA MORA", "MUNICIPALIDAD DE FILADELFIA", "MUNICIPALIDAD DE FUERTE OLIMPO",
	"MUNICIPALIDAD DE FULGENCIO YEGROS", "MUNICIPALIDAD DE GENERAL ARTIGAS", "MUNICIPALIDAD DE HERNANDARIAS", "MUNICIPALIDAD DE HOHENAU", "MUNICIPALIDAD DE HORQUETA",
	"MUNICIPALIDAD DE ITAKYRY", "MUNICIPALIDAD DE ITAUGUA", "MUNICIPALIDAD DE JUAN EMILIO O’LEARY", "MUNICIPALIDAD DE KATUETE", "MUNICIPALIDAD DE LAMBARE", "MUNICIPALIDAD DE LA PAZ",
	"MUNICIPALIDAD DE LOMA PLATA", "MUNICIPALIDAD DE LUQUE", "MUNICIPALIDAD DE MCAL. FRANCISCO SOLANO LOPEZ", "MUNICIPALIDAD DE NATALIO", "MUNICIPALIDAD DE NUEVA TOLEDO",
	"MUNICIPALIDAD DE PARAGUARI", "MUNICIPALIDAD DE PASO BARRETO", "MUNICIPALIDAD DE PEDRO JUAN CABALLERO", "MUNICIPALIDAD DE QUYQUYHO", "MUNICIPALIDAD DE SAN COSME Y DAMIAN",
	"MUNICIPALIDAD DE SAN JUAN DEL PARANA", "MUNICIPALIDAD DE SAN JUAN NEPOMUCENO", "MUNICIPALIDAD DE SAN LORENZO", "MUNICIPALIDAD DE SAN PEDRO DEL YCUAMANDIYU",
	"MUNICIPALIDAD DE SAN RAFAEL DEL PARANA", "MUNICIPALIDAD DE SAPUCAI", "MUNICIPALIDAD DE SARGENTO JOSE FELIX LOPEZ", "MUNICIPALIDAD DE TEMBIAPORA", "MUNICIPALIDAD DE TOBATI",
	"MUNICIPALIDAD DE TTE.1º MANUEL IRALA FERNANDEZ", "MUNICIPALIDAD DE VILLA ELISA", "MUNICIPALIDAD DE VILLETA", "MUNICIPALIDAD DE YBY PYTA", "MUNICIPALIDAD DE YBY YAU",
	"MUNICIPALIDAD DE YRYBUCUA", "PETROLEOS PARAGUAYOS", "PODER JUDICIAL", "POLICIA NACIONAL", "PRESIDENCIA DE LA REPUBLICA", "PROCURADURIA GENERAL DE LA REPUBLICA", "SECRETARIA DE ACCION SOCIAL",
	"SECRETARIA DE DEFENSA AL CONSUMIDOR Y AL USUARIO", "SECRETARIA DE EMERGENCIA NACIONAL", "SECRETARIA DE INFORMACION Y COMUNICACION", "SECRETARIA DEL AMBIENTE",
	"SECRETARIA DE LA NIÑEZ Y DE LA ADOLESCENCIA", "SECRETARIA DE PREVENCION DE LAVADO DE DINERO O BIENES", "SECRETARIA NACIONAL ANTICORRUPCION", "SECRETARIA NACIONAL ANTIDROGAS",
	"SECRETARIA NACIONAL DE DEPORTES", "SECRETARIA NACIONAL DE INTELIGENCIA", "SECRETARIA NACIONAL DE LA CULTURA", "SECRETARIA NACIONAL DE LA JUVENTUD",
	"SECRETARIA NACIONAL DE LA VIVIENDA Y EL HABITAT", "SECRETARIA NACIONAL DE TURISMO", "SECRETARIA TECNICA DE PLANIFICACION", "SERVICIO NACIONAL DE CALIDAD VEGETAL Y DE SEMILLA",
	"SERVICIO NACIONAL DE CALIDAD Y SALUD ANIMAL", "SERVICIO NACIONAL DE CALIDAD Y SANIDAD VEGETAL Y DE SEMILLAS", "SERVICIO NACIONAL DE SANEAMIENTO AMBIENTAL",
	"SERVICIO NACIONAL DE PROMOCION PROFESIONAL", "SINDICATURA GENERAL DE QUIEBRAS", "TRIBUNAL SUPERIOR DE JUSTICIA ELECTORAL", "UNIVERSIDAD NACIONAL DE ASUNCION",
	"UNIVERSIDAD NACIONAL DE CAAGUAZU", "UNIVERSIDAD NACIONAL DE CANINDEYU", "UNIVERSIDAD NACIONAL DE CONCEPCION", "UNIVERSIDAD NACIONAL DE ITAPUA", "UNIVERSIDAD NACIONAL DEL ESTE",
	"UNIVERSIDAD NACIONAL DE PILAR", "UNIVERSIDAD NACIONAL DE VILLARRICA DEL ESPIRITU SANTO", "VICEPRESIDENCIA DE LA REPUBLICA", "INSTITUTO DE PREVISION SOCIAL", "YACYRETA", }

	cache := map[string]bool{}

    for _, item := range institutions {  
        cache[item] = true
    }
	return cache
}

func containsKWOfInsts(s string) bool {
	kw := [40]string { "ENTIDAD","MUNICIPALIDAD","CAJA DE","CONSEJO","UNIVERSIDAD",
	"CREDITO","EMPRESA","AGENCIA","SECRETARIA","POLICIA","ADMINISTRACION","MINISTERIO",
	"CONGRESO","CAMARA","PODER","TRIBUNAL","INSTITUTO","COMANDO","GOBERNACION","CORTE",
	"SERVICIO","DIRECCION","COMPAÑIA","HOSPITAL","INDUSTRIA","JURADO","SINDICATURA",
	"MECANISMO","BANCO","FONDO","ENTE","FACULTAD","INDUSTRIA","AUDITORIA","PROCURADURIA",
	"COMISION","HONORABLE","ARMADA","NACIONAL","COLEGIO", }
	
	s = removeAccents(s)
	for _, value := range kw {
		// if kw is first word of current token
		if hasTrailingSpaces(s, value) {
			return true
		}
	}
	return false
}

func isJobFormField(s string) bool {
	formField := []string {
		"TIPO",
		"INSTITUCION",
		"DIRECCION",
		"DEPENDENCIA",
		"CATEGORIA",
		"NOMBRADO/CONTRATADO",
		"CARGO",
		"FECHA ASUNC./CESE/OTROS",
		"ACTO ADMINIST",
		"FECHA ACT. ADM",
		"TELEFONO",
		"COMISIONADO",
		"FECHA INGRESO",
		"FECHA EGRESO",
	}

	s = removeAccents(s)
	for _, value := range formField {
		if isCurrLine(s, value) {
			return true
		}
	}

	return false
}

func isJobFormCommonAnswer(s string) bool {
	commonAnswer := []string{ "SI", "NO", "PERSONAL DE BLANCO", "RECEPCIONADO" }

	s = removeAccents(s)
	for _, value := range commonAnswer {
		if s == value {
			return true
		}
	}

	return false
}
