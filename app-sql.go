package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	limpiarPantalla()
	iniciarBd()
}

func menuPrincipal(db *sql.DB) {
	mostrarMenu()

	fmt.Println("Seleccione una opcion... ")
	var input int
	fmt.Scan(&input)
	eligeOpcionMenu(input, db)
}

func iniciarBd() {
	for true {
		mostrarMenu()
		fmt.Println("Seleccione una opcion... ")
		var input int
		fmt.Scan(&input)

		if input == 1 {
			crearBaseDeDatos()
			//conexion con base de datos.
			db, err := sql.Open("postgres", "user=postgres host=localhost dbname=turnos sslmode=disable")
			if err != nil {
				log.Fatal(err)
			}
			limpiarPantalla()
			fmt.Println("Base de datos creada.")
			menuPrincipal(db)
		}
		if input == 0 {
			os.Exit(0)
		} else {
			limpiarPantalla()
			fmt.Println("Base de datos no creada.")
		}
	}
}

func mostrarMenu() {
	fmt.Println(`
 __                                            
/\ \__                                         
\ \ ,_\  __  __  _ __    ___     ___     ____  
 \ \ \/ /\ \/\ \/\  __\/  _  \  / __ \  / ,__\ 
  \ \ \_\ \ \_\ \ \ \/ /\ \/\ \/\ \L\ \/\__,  \
   \ \__\\ \____/\ \_\ \ \_\ \_\ \____/\/\____/
    \/__/ \/___/  \/_/  \/_/\/_/\/___/  \/___/ 
•••••••••••••••••••••••••••••••••••••••••••••••••
•   1 → Crear base de datos.                 •
•   2 → Crear Tablas.                        •
•   3 → Cargar Datos.                        •
•   4 → Cargar PK's.                         •
•   5 → Cargar FK's.                         •
•   6 → Cargar Stored Procedures y Triggers. •
•   7 → Ingresar Menu de acciones.           •
•   8 → Borrar PK's y FK's                   •
•   0 → Salir.                               •
•••••••••••••••••••••••••••••••••••••••••••••••••`)
}

func eligeOpcionMenu(opcion int, db *sql.DB) {
	limpiarPantalla()

	switch opcion {
	case 1:
		fmt.Println("Base de datos ya creada")
		menuPrincipal(db)
	case 2:
		_, err := db.Exec(leerArchivo("sql-files/create-tables.sql"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Tablas creadas.")
		menuPrincipal(db)
	case 3:
		_, err := db.Exec(leerArchivo("sql-files/add-data.sql"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Datos cargados.")
		menuPrincipal(db)
	case 4:
		_, err := db.Exec(leerArchivo("sql-files/add-primary-keys.sql"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("PK's cargadas.")
		menuPrincipal(db)
	case 5:
		_, err := db.Exec(leerArchivo("sql-files/add-foreign-keys.sql"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("FK's cargadas.")
		menuPrincipal(db)
	case 6:
		_, err := db.Exec(leerArchivo("sql-files/cancelacion-de-turnos.sql"))
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(leerArchivo("sql-files/atencion-de-turno.sql"))
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(leerArchivo("sql-files/generacion-de-turnos.sql"))
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(leerArchivo("sql-files/reserva-de-turnos.sql"))
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(leerArchivo("sql-files/liquidacion-obras-sociales.sql"))
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(leerArchivo("sql-files/generar-emails.sql"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Stored Procedures y Triggers cargados.")
		menuPrincipal(db)
	case 7:
		fmt.Println("Ingresando al menú de acciones...")
		menuAcciones(db)
	case 8:
		_, err := db.Exec(leerArchivo("sql-files/drop-keys.sql"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("PK's y FK's eliminadas.")
		menuPrincipal(db)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Opcion no existente")
		menuPrincipal(db)
	}
}

func menuAcciones(db *sql.DB) {
	mostrarMenuAcciones()

	fmt.Println("Seleccione una opcion... ")
	var input int
	fmt.Scan(&input)
	eligeOpcionMenuAcciones(input, db)
}

func mostrarMenuAcciones() {
	fmt.Println(`
 __                                            
/\ \__                                         
\ \ ,_\  __  __  _ __    ___     ___     ____  
 \ \ \/ /\ \/\ \/\  __\/  _  \  / __ \  / ,__\ 
  \ \ \_\ \ \_\ \ \ \/ /\ \/\ \/\ \L\ \/\__,  \
   \ \__\\ \____/\ \_\ \ \_\ \_\ \____/\/\____/
    \/__/ \/___/  \/_/  \/_/\/_/\/___/  \/___/ 
•••••••••••••••••••••••••••••••••••••••••••••••••••
•   1 → Generar turnos.                        •
•   2 → Reservar turno.                        •
•   3 → Cancelar turnos.                       •
•   4 → Atender turno.                         •
•   5 → Liquidar obras sociales.               •
•   6 → Generar emails de recordatorio.        •
•   7 → Generar emails de turnos perdidos.     •
•   0 → Volver al menu principal.              •
•••••••••••••••••••••••••••••••••••••••••••••••••••`)
}

func eligeOpcionMenuAcciones(opcion int, db *sql.DB) {
	limpiarPantalla()

	switch opcion {
	case 1:
		fmt.Printf("Generando turnos para el mes 06 del año 2023...\n")
		generarTurnos(db)
		fmt.Printf("Turnos generados.")
		menuAcciones(db)
	case 2:
		fmt.Printf("Reservando turnos...\n")
		reservarTurnosDePrueba(db)
		fmt.Printf("Turnos reservados.")
		menuAcciones(db)
	case 3:
		fmt.Printf("Cancelando turnos de médique 49194249...\n")
		cancelarTurnosDePrueba(db)
		fmt.Printf("Turnos cancelados.")
		menuAcciones(db)
	case 4:
		fmt.Printf("Actualizando turnos atendidos...\n")
		atenderTurnosDePrueba(db)
		fmt.Printf("Turnos actualizados.")
		menuAcciones(db)
	case 5:
		fmt.Printf("Liquidando obras sociales...\n")
		liquidarObrasSocialesDePrueba(db)
		fmt.Printf("Obras sociales liquidadas.")
		menuAcciones(db)
	case 6:
		fmt.Printf("Enviando recordatorios...\n")
		enviarEmails(db, "recordatorios")
		fmt.Printf("Recordatorios generados.")
		menuAcciones(db)
	case 7:
		fmt.Printf("Enviando avisos de pérdidas de turnos...\n")
		enviarEmails(db, "perdidas")
		fmt.Printf("Avisos generados.")
		menuAcciones(db)
	case 0:
		fmt.Printf("Volviendo al menu principal...")
		menuPrincipal(db)
	default:
		fmt.Println("Opcion no existente.")
		menuAcciones(db)
	}
}

func generarTurnos(db *sql.DB) {
	_, err := db.Exec(`begin;
				set transaction isolation level serializable;
				select generar_turnos_disponibles(2023, 6);
				commit;`)

	if err != nil {
		log.Fatal(err)
	}
}

func reservarTurnosDePrueba(db *sql.DB) {
	_, err := db.Exec(`begin;
				set transaction isolation level serializable;
				select reservar_turnos();
				commit;`)

	if err != nil {
		log.Fatal(err)
	}
}

func atenderTurnosDePrueba(db *sql.DB) {
	atenderTurno(db, 791)
	atenderTurno(db, 792)
	atenderTurno(db, 795)
	atenderTurno(db, 796)
}

func cancelarTurnosDePrueba(db *sql.DB) {	
	rows, err := db.Query(`begin;
						set transaction isolation level serializable;
						select cancelar_turnos(49194249, '2023-06-01', '2023-06-28');
						commit;`)
	
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var resultado int
	for rows.Next() {
		err = rows.Scan(&resultado)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func liquidarObrasSocialesDePrueba(db *sql.DB) {
	_, err := db.Exec(`begin;
				set transaction isolation level serializable;
				select liquidar_obras_sociales(2023, 6);
				commit;`)

	if err != nil {
		log.Fatal(err)
	}
}

func setearIsolationLevel(db *sql.DB, isolationLevel string) {
	query := `begin;
				set transaction isolation level ` + isolationLevel + ";"
	_, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}
}

func commitTransaccion(db *sql.DB) {
	_, err := db.Query("commit;")
	if err != nil {
		log.Fatal(err)
	}
}

func reservarTurno(db *sql.DB, nro_historia_clinica int, dni_medique_pedide int, fecha_pedida string, hora_pedida string) {
	var resultado bool
	setearIsolationLevel(db, "serializable")

	query := "select reservar_turno($1, $2, $3, $4);"
	rows, err := db.Query(query, strconv.Itoa(nro_historia_clinica), strconv.Itoa(dni_medique_pedide), fecha_pedida, hora_pedida)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	commitTransaccion(db)

	for rows.Next() {
		err = rows.Scan(&resultado)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func cancelarTurnos(db *sql.DB, dni_medique int, fecha_desde string, fecha_hasta string) {
	setearIsolationLevel(db, "serializable")

	query := "select cancelar_turnos($1,$2,$3);"
	rows, err := db.Query(query, strconv.Itoa(dni_medique), fecha_desde, fecha_hasta)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	commitTransaccion(db)

	var resultado int
	for rows.Next() {
		err = rows.Scan(&resultado)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func atenderTurno(db *sql.DB, nro_turno int) {
	query := "select atender_paciente($1);"
	rows, err := db.Query(query, strconv.Itoa(nro_turno))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var resultado bool
	for rows.Next() {
		err = rows.Scan(&resultado)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func liquidarObrasSociales(db *sql.DB, anio int, mes int) {
	setearIsolationLevel(db, "serializable")
	query := "select liquidar_obras_sociales($1, $2);"
	_, err := db.Query(query, strconv.Itoa(anio), strconv.Itoa(mes))
	if err != nil {
		log.Fatal(err)
	}
	commitTransaccion(db)
}

func enviarEmails(db *sql.DB, tipo_de_envio string) {
	if tipo_de_envio == "recordatorios" {
		_, err := db.Query("select generar_recordatorios()")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := db.Query("select generar_avisos_perdida_de_turnos()")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func crearBaseDeDatos() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	_, err = db.Exec(`drop database if exists turnos;`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`create database turnos;`)
	if err != nil {
		log.Fatal(err)
	}
}

func leerArchivo(direccion string) string {
	resultado, err := ioutil.ReadFile(direccion)
	if err != nil {
		log.Fatal(err)
	}
	return string(resultado)
}

func limpiarPantalla() {
	term := exec.Command("clear")
	output, erro := term.Output()
	if erro != nil {
		fmt.Println("Error al ejecutar comando")
		return
	}
	fmt.Print(string(output))
}
