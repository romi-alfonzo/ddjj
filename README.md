# Parser simple para declaraciones juradas de la CGR

La contraloria general de la república publica de manera regular todas las
declaraciones juradas de los funcionarios públicos y autoridades electas
(https://portaldjbr.contraloria.gov.py/portal-djbr/).

Este proyecto provee de un modulo GO que recibe como entrada una declaración
y produce un JSON con los datos extraídos.


## Uso

```
cd parser
go run main.go test_declarations/267948_MARIO_ABDO_BENITEZ.pdf
```


