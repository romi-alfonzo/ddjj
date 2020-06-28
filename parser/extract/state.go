package extract

import (
	"bufio"
	"ddjj/parser/declaration"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var stateUses = []string{
	"VIVIENDA",
	"BALDIO",
	"TERRENO BALDIO",
	"ALQUILER",
	"GRANJA",
	"EXPLOTACION",
	"CANTERA",
}

// RealStates returns the real states owned by the official.
func RealStates(scanner *bufio.Scanner) []*declaration.RealState {

	var skip = []string{
		"#",
		"Nº FINCA",
		"DATOS PROTEGIDOS",
		"PAÍS:",
		"CTA. CTE. CTRAL. O PADRON",
		"USO",
		"DISTRITO:",
		"SUP. M2",
		"AÑO DE ADQ.",
		"VALOR CONST. G.",
		"CONST.",
		"VALOR TERRENO G.",
		"TIPO DE ADQ.:",
		"IMPORTE",
	}

	scanner = moveUntil(scanner, "1.4 INMUEBLES", true)

	var states []*declaration.RealState
	opts := &stateOpts{
		state:   &declaration.RealState{},
		counter: 0,
		typ:     1,
		scanner: scanner,
	}

	index := 1
	skip = append(skip, strconv.Itoa(index))
	var total int64
	for scanner.Scan() {
		line := scanner.Text()

		// Stop looking for real state when this is found.
		if line == "TOTAL INMUEBLES:" {
			total = getTotalInCategory(scanner)

			// Next page or end.
			scanner = moveUntil(scanner, "Nº FINCA", true)
			line = scanner.Text()
			if line == "" {
				break
			}

			opts.next = nil
			opts.previous = nil
			index = 1
			skip[len(skip)-1] = strconv.Itoa(index)
		}

		if strings.Contains(line, "OBS:") || strings.Contains(line, "RECEPCIONADO EL:") ||
			isDate(line) || isBarCode(line) {

			continue
		}
		if contains(skip, line) || line == "" {
			continue
		}

		// Ver el comentario en getRealState4.
		if opts.typ == 1 && opts.state.Padron == "" && contains(stateUses, line) {
			opts.typ = 4
			opts.counter = 0
		}

		s := getRealState(opts, line)
		if s != nil {
			states = append(states, s)
			opts.counter = -1

			if opts.next != nil {
				opts.state = opts.next
			} else {
				opts.state = &declaration.RealState{}
			}

			opts.previous = s

			// Skip the following item #.
			index++
			skip[len(skip)-1] = strconv.Itoa(index)
		}

		opts.counter++
	}

	totalState := addRealState(states)
	if totalState != total {
		fmt.Println(total, totalState)
		log.Fatal("The amounts in real state do not match")
	}

	return states
}

type stateOpts struct {
	state    *declaration.RealState
	next     *declaration.RealState
	previous *declaration.RealState
	scanner  *bufio.Scanner
	counter  int
	typ      int
}

func getRealState(opts *stateOpts, line string) *declaration.RealState {

	switch opts.typ {
	case 1:
		return getRealState1(opts, line)
	case 2:
		return getRealState2(opts, line)
	case 3:
		return getRealState3(opts, line)
	case 4:
		return getRealState4(opts, line)
	}

	return nil
}

// Este es el caso de la mayoría de los items. Los valores se extraen en este
// orden.
func getRealState1(opts *stateOpts, line string) *declaration.RealState {
	switch opts.counter {
	case 0:
		opts.state.Pais = line
		break
	case 1:
		opts.state.Padron = line
		break
	case 2:
		// Usos que empiezan con "EXPLOTACION" tienen dos líneas.
		if line == "EXPLOTACION" {
			opts.scanner.Scan()
			nextLine := opts.scanner.Text()
			opts.state.Uso = line + " " + nextLine
			break
		}

		opts.state.Uso = line
		break
	case 3:
		opts.state.Distrito = line
		break
	case 4:
		value, _ := strconv.ParseInt(line, 10, 64)
		opts.state.SuperficieTerreno = value
		break
	case 5:
		opts.state.ValorTerreno = stringToInt64(line)
		break
	case 6:
		opts.state.Adquisicion = stringToYear(line)
		break
	case 7:
		value, _ := strconv.ParseInt(line, 10, 64)
		opts.state.SuperficieConstruccion = value
		break
	case 8:
		opts.state.ValorConstruccion = stringToInt64(line)
		break
	case 9:
		opts.state.Importe = stringToInt64(line)
		break
	case 10:
		if isNumber(line) {
			s := opts.state
			// Ver el comentario en getRealState2.
			opts.next = &declaration.RealState{}
			opts.next.ValorConstruccion = stringToInt64(line)
			opts.typ = 2

			// Ver comentario en getRealState3.
			if s.Importe != s.ValorTerreno+s.ValorConstruccion {
				opts.next.Adquisicion = int(s.SuperficieConstruccion)
				s.SuperficieConstruccion = s.ValorConstruccion
				s.ValorConstruccion = s.Importe
				s.Importe = stringToInt64(line)
				opts.typ = 3
			}
		} else {
			opts.state.TipoAdquisicion = line
		}

		return opts.state
	}

	return nil
}

// Este caso sucede cuando el valor de la construcción de un item i + 1 aparece
// antes que el tipo de adquisición de i.
// Esto sucedía con, por ejemplo, el immueble con padrón 9412 de Oscar González
// Daher del 2016.
func getRealState2(opts *stateOpts, line string) *declaration.RealState {
	switch opts.counter {
	case 0:
		opts.state.Importe = stringToInt64(line)
		break
	case 1:
		opts.previous.TipoAdquisicion = line
		break
	case 2:
		opts.state.Pais = line
		break
	case 3:
		opts.state.Padron = line
		break
	case 4:
		// Usos que empiezan con "EXPLOTACION" tienen dos líneas.
		if line == "EXPLOTACION" {
			opts.scanner.Scan()
			nextLine := opts.scanner.Text()
			opts.state.Uso = line + " " + nextLine
			break
		}

		opts.state.Uso = line
		break
	case 5:
		opts.state.Distrito = line
		break
	case 6:
		value, _ := strconv.ParseInt(line, 10, 64)
		opts.state.SuperficieTerreno = value
		break
	case 7:
		opts.state.ValorTerreno = stringToInt64(line)
		break
	case 8:
		opts.state.Adquisicion = stringToYear(line)
		break
	case 9:
		value, _ := strconv.ParseInt(line, 10, 64)
		opts.state.SuperficieConstruccion = value
		break
	case 10:
		opts.next = nil
		opts.state.TipoAdquisicion = line
		opts.typ = 1
		return opts.state
	}

	return nil
}

// Este es el caso en el que el valor del terreno de i+1 aparece inmediatamente
// después del valor del terreno de i.
// Esto sucedía con, por ejemplo, el immueble con padrón 27-0026.24 de Oscar González
// Daher del 2016.
func getRealState3(opts *stateOpts, line string) *declaration.RealState {
	switch opts.counter {
	case 0:
		opts.previous.TipoAdquisicion = line
	case 1:
		opts.state.Pais = line
		break
	case 2:
		opts.state.Padron = line
		break
	case 3:
		// Usos que empiezan con "EXPLOTACION" tienen dos líneas.
		if line == "EXPLOTACION" {
			opts.scanner.Scan()
			nextLine := opts.scanner.Text()
			opts.state.Uso = line + " " + nextLine
			break
		}

		opts.state.Uso = line
		break
	case 4:
		opts.state.Distrito = line
		break
	case 5:
		value, _ := strconv.ParseInt(line, 10, 64)
		opts.state.SuperficieTerreno = value
		break
	case 6:
		opts.state.ValorTerreno = stringToInt64(line)
		break
	case 7:
		value, _ := strconv.ParseInt(line, 10, 64)
		opts.state.SuperficieConstruccion = value
		break
	case 8:
		opts.state.ValorConstruccion = stringToInt64(line)
		break
	case 9:
		opts.state.Importe = stringToInt64(line)
	case 10:
		opts.next = nil
		opts.state.TipoAdquisicion = line
		opts.typ = 1
		return opts.state
	}

	return nil
}

// Este es el caso cuando el el padrón y el uso aparecen antes que el pais.
// Esto sucece, por ejemplo, en la declaración de Juan Eudes Afara Maciel del
// 2014.
func getRealState4(opts *stateOpts, line string) *declaration.RealState {
	opts.state.Padron = opts.state.Pais

	switch opts.counter {
	case 0:
		// Usos que empiezan con "EXPLOTACION" tienen dos líneas.
		if line == "EXPLOTACION" {
			opts.scanner.Scan()
			nextLine := opts.scanner.Text()
			opts.state.Uso = line + " " + nextLine
			break
		}
		opts.state.Uso = line
		break
	case 1:
		opts.state.Pais = line
		break
	case 2:
		opts.state.Distrito = line
		break
	case 3:
		value, _ := strconv.ParseInt(line, 10, 64)
		opts.state.SuperficieTerreno = value
		break
	case 4:
		opts.state.ValorTerreno = stringToInt64(line)
		break
	case 5:
		opts.state.Adquisicion = stringToYear(line)
		break
	case 6:
		value, _ := strconv.ParseInt(line, 10, 64)
		opts.state.SuperficieConstruccion = value
		break
	case 7:
		opts.state.ValorConstruccion = stringToInt64(line)
		break
	case 8:
		opts.state.Importe = stringToInt64(line)
		break
	case 9:
		opts.typ = 1
		opts.state.TipoAdquisicion = line

		return opts.state
	}

	return nil
}

func addRealState(states []*declaration.RealState) int64 {
	var total int64
	for _, d := range states {
		total += d.Importe
	}

	return total
}
