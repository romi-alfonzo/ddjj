package extract

import (
	"bufio"
	"ddjj/parser/declaration"
	"log"
	"strconv"
	"strings"
)

// Deposits returns the deposits at financial institutions.
func Deposits(scanner *bufio.Scanner) []*declaration.Deposit {
	var skip = []string{
		"#",
		"TIPO ENTIDAD",
		"NOMBRE DE ENTIDAD",
		"TIPO DE CUENTA",
		"Nº DE CUENTA",
		"PAÍS",
		"IMPORTE",
		"DATOS PROTEGIDOS",
	}

	scanner = moveUntil(scanner, "1.2 DEPÓSITOS", true)

	var deposits []*declaration.Deposit
	opts := &depositOpts{
		deposit: &declaration.Deposit{},
		counter: 0,
	}

	index := 1
	skip = append(skip, strconv.Itoa(index))
	var total int64
	for scanner.Scan() {
		line := scanner.Text()

		// Stop looking for deposits in the page when this is found.
		if line == "TOTAL DEPÓSITOS:" {
			total = getTotalInCategory(scanner)

			// Next page or end.
			scanner = moveUntil(scanner, "TIPO ENTIDAD", true)
			line = scanner.Text()
			if line == "" {
				break
			}

			index = 1
		}

		if strings.Contains(line, "OBS:") {
			continue
		}
		if contains(skip, line) || line == "" {
			continue
		}

		d := getDeposit(opts, line)
		if d != nil {
			deposits = append(deposits, d)
			opts.counter = -1
			opts.deposit = &declaration.Deposit{}

			// Skip the following item #.
			index++
			skip = append(skip, strconv.Itoa(index))
		}

		opts.counter++
	}

	totalDeposits := addDeposits(deposits)
	if totalDeposits != total {
		log.Fatal("Deposits do not match")
	}

	return deposits
}

type depositOpts struct {
	deposit *declaration.Deposit
	counter int
}

func getDeposit(opts *depositOpts, line string) *declaration.Deposit {

	switch opts.counter {
	case 0:
		opts.deposit.TipoEntidad = line
		break
	case 1:
		opts.deposit.Entidad = line
		break
	case 2:
		opts.deposit.Tipo = line
		break
	case 3:
		opts.deposit.Pais = line
		break
	case 4:
		value := strings.ReplaceAll(line, ".", "")
		i, _ := strconv.ParseInt(value, 10, 64)
		opts.deposit.Importe = i
		return opts.deposit
	}

	return nil
}

func addDeposits(deposits []*declaration.Deposit) int64 {
	var total int64
	for _, d := range deposits {
		total += d.Importe
	}

	return total
}
