package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/pretty"
	bolt "go.etcd.io/bbolt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type Paciente struct {
	NroPaciente     int
	Nombre          string
	Apellido        string
	DniPaciente     int
	FechaNacimiento string
	NroObraSocial   int
	NroAfiliade     int
	Domicilio       string
	Telefono        string
	Email           string
}

type Medique struct {
	DniMedique           int
	Nombre               string
	Apellido             string
	Especialidad         string
	MontoConsultaPrivada float64
	Telefono             string
}

type Consultorio struct {
	NroConsultorio int
	Nombre         string
	Domicilio      string
	CodigoPostal   string
	Telefono       string
}

type ObraSocial struct {
	NroObraSocial    int
	Nombre           string
	ContactoNombre   string
	ContactoApellido string
	ContactoTelefono string
	ContactoEmail    string
}

type Turno struct {
	NroTurno              int
	Fecha                 string
	NroConsultorio        int
	DniMedique            int
	NroPaciente           int
	NroObraSocialConsulta int
	NroAfiliadeConsulta   int
	MontoPaciente         float64
	MontoObraSocial       float64
	FechaReserva          string
	Estado                string
}

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
	// abre transacción de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

	err = b.Put(key, val)
	if err != nil {
		return err
	}
	// cierra transacción
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
	var buf []byte
	// abre una transacción de lectura
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})
	return buf, err
}

func cargarBolt(db *bolt.DB) {
	cargarPacientes(db)
	cargarMediques(db)
	cargarConsultorios(db)
	cargarObrasSociales(db)
	cargarTurnos(db)
}

func cargarPacientes(db *bolt.DB) {
	pacientes := []Paciente{
		Paciente{NroPaciente: 1, Nombre: "Mariana", Apellido: "Lopez", DniPaciente: 48914850, FechaNacimiento: "2008-08-25", NroObraSocial: 3, NroAfiliade: 729563014, Domicilio: "Paunero 1135", Telefono: "11-6166-3670", Email: "mlopez3@email.com"},
		Paciente{NroPaciente: 2, Nombre: "Marcela", Apellido: "Trupia", DniPaciente: 39176511, FechaNacimiento: "1990-06-26", NroObraSocial: 2, NroAfiliade: 485970623, Domicilio: "Entre Ríos 698", Telefono: "11-1054-7634", Email: "mtrupia9@email.com"},
		Paciente{NroPaciente: 3, Nombre: "Pablo", Apellido: "Agüero", DniPaciente: 46851930, FechaNacimiento: "2007-08-19", NroObraSocial: 1, NroAfiliade: 862109437, Domicilio: "Hipólito Yrigoyen 2111", Telefono: "11-8220-0162", Email: "pagüero0@email.com"},
	}

	for _, paciente := range pacientes {
		cargar(db, paciente, "paciente", paciente.NroPaciente)
	}
}

func cargarMediques(db *bolt.DB) {
	mediques := []Medique{
		Medique{DniMedique: 43461574, Nombre: "Stephen", Apellido: "Strange", Especialidad: "Clínico", MontoConsultaPrivada: 3181.39, Telefono: "11-4033-1622"},
		Medique{DniMedique: 49194249, Nombre: "Peter", Apellido: "Parker", Especialidad: "Clínico", MontoConsultaPrivada: 2774.51, Telefono: "11-4045-2400"},
		Medique{DniMedique: 41724061, Nombre: "Steve", Apellido: "Rogers", Especialidad: "Clínico", MontoConsultaPrivada: 2087.39, Telefono: "11-7597-2544"},
	}

	for _, medique := range mediques {
		cargar(db, medique, "medique", medique.DniMedique)
	}
}

func cargarConsultorios(db *bolt.DB) {
	consultorios := []Consultorio{
		Consultorio{NroConsultorio: 1, Nombre: "Anexus Consultorios Médicos", Domicilio: "Av. Pres. Hipólito Yrigoyen 2375", CodigoPostal: "B1666GQH", Telefono: "11-2351-7877"},
		Consultorio{NroConsultorio: 2, Nombre: "Consultorio San Miguel", Domicilio: "Italia 1213", CodigoPostal: "B1663NXY", Telefono: "11-4131-7095"},
		Consultorio{NroConsultorio: 3, Nombre: "Consultorio Médico Lourdes", Domicilio: "Padre Stoppler 1121", CodigoPostal: "B1615KCW", Telefono: "11-4801-1221"},
	}

	for _, consultorio := range consultorios {
		cargar(db, consultorio, "consultorio", consultorio.NroConsultorio)
	}
}

func cargarObrasSociales(db *bolt.DB) {
	obrasSociales := []ObraSocial{
		ObraSocial{NroObraSocial: 1, Nombre: "Obra Social de panaderos, pasteleros y factureros de Entre Rios", ContactoNombre: "Pedro", ContactoApellido: "Alarcon", ContactoTelefono: "11-2759-3533", ContactoEmail: "pedroalarcon@email.com"},
		ObraSocial{NroObraSocial: 2, Nombre: "Obra Social del personal de la actividad azucarera tucumana", ContactoNombre: "Catalina", ContactoApellido: "Ramirez", ContactoTelefono: "11-3868-6462", ContactoEmail: "catalinaramirez@email.com"},
		ObraSocial{NroObraSocial: 3, Nombre: "Obra Social del personal de panaderias", ContactoNombre: "Marisa", ContactoApellido: "Rodriguez", ContactoTelefono: "11-4648-6649", ContactoEmail: "marisarodriguez@email.com"},
	}

	for _, obraSocial := range obrasSociales {
		cargar(db, obraSocial, "obraSocial", obraSocial.NroObraSocial)
	}
}

func cargarTurnos(db *bolt.DB) {
	turnos := []Turno{
		Turno{NroTurno: 1, Fecha: "2023-06-13 08:00:00", NroConsultorio: 2, DniMedique: 43461574, NroPaciente: 3, NroObraSocialConsulta: 1, NroAfiliadeConsulta: 862109437, MontoPaciente: 1740.00, MontoObraSocial: 4060.00, FechaReserva: "2023-06-01 09:48:50", Estado: "atendido"},
		Turno{NroTurno: 2, Fecha: "2023-06-28 11:15:00", NroConsultorio: 1, DniMedique: 39176511, NroPaciente: 2, NroObraSocialConsulta: 2, NroAfiliadeConsulta: 485970623, MontoPaciente: 1550.00, MontoObraSocial: 3900.00, FechaReserva: "2023-06-02 10:34:12", Estado: "reservado"},
		Turno{NroTurno: 3, Fecha: "2023-06-10 12:30:00", NroConsultorio: 2, DniMedique: 48914850, NroPaciente: 1, NroObraSocialConsulta: 3, NroAfiliadeConsulta: 729563014, MontoPaciente: 1400.00, MontoObraSocial: 3780.00, FechaReserva: "2023-06-01 11:12:35", Estado: "cancelado"},
	}

	for _, turno := range turnos {
		cargar(db, turno, "turno", turno.NroTurno)
	}
}

func cargar(db *bolt.DB, v interface{}, bucketName string, key int) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, bucketName, []byte(strconv.Itoa(key)), data)

	_, err = ReadUnique(db, bucketName, []byte(strconv.Itoa(key)))
	if err != nil {
		log.Fatal(err)
	}
}

func mostrarDatos(db *bolt.DB, bucketName string) {
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("Bucket no encontrado")
		}
		// Iterar sobre todas las keys
		err := bucket.ForEach(func(key, value []byte) error {
			// Obtener el valor utilizando la key
			value = bucket.Get([]byte(key))
			fmt.Printf("Key: %s\n Value:\n %s \n", string(key), string(pretty.Pretty(value)))
			return nil
		})
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}

func clearConsole() {
	term := exec.Command("clear")
	output, err := term.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(output))
}

func eligeOpcionBolt(db *bolt.DB, opcion int) {
	clearConsole()
	switch opcion {
	case 1:
		fmt.Println("Base de datos creada exitosamente.")
	case 2:
		cargarBolt(db)
		fmt.Println("Datos cargados exitosamente.")
		menuBolt(db)
	case 3:
		fmt.Println("•••••••••••••• Pacientes ••••••••••••••")
		mostrarDatos(db, "paciente")
		menuBolt(db)
	case 4:
		fmt.Println("•••••••••••••• Mediques ••••••••••••••")
		mostrarDatos(db, "medique")
		menuBolt(db)
	case 5:
		fmt.Println("•••••••••••••• Consultorios ••••••••••••••")
		mostrarDatos(db, "consultorio")
		menuBolt(db)
	case 6:
		fmt.Println("•••••••••••••• Obras sociales ••••••••••••••")
		mostrarDatos(db, "obraSocial")
		menuBolt(db)
	case 7:
		fmt.Println("•••••••••••••• Turnos ••••••••••••••")
		mostrarDatos(db, "turno")
		menuBolt(db)
	case 0:
		os.Exit(0)
	default:
		clearConsole()
		fmt.Println("Opcion no existente...")
		menuBolt(db)
	}
}

func mostrarMenuBolt() {
	fmt.Println(`
 __                                            
/\ \__                                         
\ \ ,_\  __  __  _ __    ___     ___     ____  
 \ \ \/ /\ \/\ \/\  __\/  _  \  / __ \  / ,__\ 
  \ \ \_\ \ \_\ \ \ \/ /\ \/\ \/\ \L\ \/\__,  \
   \ \__\\ \____/\ \_\ \ \_\ \_\ \____/\/\____/
    \/__/ \/___/  \/_/  \/_/\/_/\/___/  \/___/ 
•••••••••••••••••••••••••••••••••••••••••••••••••
•   1 → Crear base de datos noSQL.           •
•   2 → Cargar datos.                        •
•   3 → Ver pacientes.                       •
•   4 → Ver mediques.                        •
•   5 → Ver consultorios.                    •
•   6 → Ver obras sociales.                  •
•   7 → Ver turnos.                          •
•   0 → Salir.                               •
•••••••••••••••••••••••••••••••••••••••••••••••••`)
}

func menuBolt(db *bolt.DB) {
	mostrarMenuBolt()

	fmt.Println("Seleccione una opcion... ")
	var input int
	fmt.Scan(&input)
	eligeOpcionBolt(db, input)
}

func iniciarBd() {
	for true {
		mostrarMenuBolt()
		fmt.Println("Seleccione una opcion... ")
		var input int
		fmt.Scan(&input)
		if input == 1 {
			db, err := bolt.Open("turnos.db", 0600, nil)
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()
			clearConsole()
			fmt.Println("Base de datos creada.")
			menuBolt(db)
		}
		if input == 0 {
			os.Exit(0)
		} else {
			clearConsole()
			fmt.Println("Base de datos no creada.")
		}
	}
}

func main() {
	clearConsole()
	iniciarBd()
}
