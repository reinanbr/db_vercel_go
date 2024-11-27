package access_site_model


import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/jackc/pgx/v4/pgxpool"
)




type AccessSite struct {
	ID       int
	Site     string
	IP       string
	Hostname string
	Date     time.Time
	Provedor string
	City     string
	State    string
	Country  string
}


type AccessSiteJson struct {
	ID       int       `json:"id"`
	Site     string    `json:"site"`
	IP       string    `json:"ip"`
	Hostname string    `json:"hostname"`
	Date     time.Time `json:"date"`
	Provedor string    `json:"provedor"`
	City     string    `json:"city"`
	State    string    `json:"state"`
	Country  string    `json:"country"`
}

type ResponseModel struct {
	Success int
	Message string
}



// Criar um novo registro
func CreateAccessSite(pool *pgxpool.Pool, site, ip, hostname, date, provedor, city, state, country string)(ResponseModel){
	parsedTime, errTime := time.Parse(time.RFC3339, date)
	parsTime := parsedTime.UTC()
	if errTime != nil {
		fmt.Println("Erro ao parsear a data:", errTime)
	}
	fmt.Printf("date RFC: %s| date UTC: %s\n",date,parsedTime)
	var responseModel ResponseModel
	responseModel.Message = "null"
	responseModel.Success = 0


	sql := `INSERT INTO table_access_sites (site, ip, hostname, date, provedor, city, state, country)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	
	_, err := pool.Exec(context.Background(), sql, site, ip, hostname, parsTime, provedor, city, state, country)
	if err != nil {
		log.Fatal("Erro ao inserir o registro: ", err)
		responseModel.Message = "error: psql not sucefull connection"
		return responseModel
	}
	fmt.Println("Novo registro criado com sucesso:", site)
		responseModel.Success = 1
		responseModel.Message = "access site info saved"
		return responseModel
}


/*
// Ler todos os registros
func ReadAccessSites(pool *pgxpool.Pool) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM table_access_sites")
	if err != nil {
		log.Fatal("Erro ao buscar registros: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var site AccessSite
		err := rows.Scan(&site.ID, &site.Site, &site.IP, &site.Hostname, &site.Date, &site.Provedor, &site.City, &site.State, &site.Country)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Site: %s, IP: %s, Hostname: %s, Date: %s, Provedor: %s, City: %s, State: %s, Country: %s\n",
			site.ID, site.Site, site.IP, site.Hostname, site.Date.Format("2006-01-02 15:04:05"), site.Provedor, site.City, site.State, site.Country)
	}
}
*/

func ReadAccessSites(pool *pgxpool.Pool) ([]AccessSiteJson, error) {
    rows, err := pool.Query(context.Background(), "SELECT * FROM table_access_sites")
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar registros: %v", err)
    }
    defer rows.Close()

    var sites []AccessSiteJson
    for rows.Next() {
        var site AccessSiteJson
        err := rows.Scan(&site.ID, &site.Site, &site.IP, &site.Hostname, &site.Date, &site.Provedor, &site.City, &site.State, &site.Country)
        if err != nil {
            return nil, fmt.Errorf("erro ao escanear registro: %v", err)
        }
        sites = append(sites, site)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("erro ao iterar registros: %v", err)
    }

    return sites, nil
}






// Procurar registro por IP
func FindAccessSiteByIP(pool *pgxpool.Pool, ip string) {
	var site AccessSite
	sql := `SELECT * FROM table_access_sites WHERE ip = $1`

	err := pool.QueryRow(context.Background(), sql, ip).Scan(&site.ID, &site.Site, &site.IP, &site.Hostname, &site.Date, &site.Provedor, &site.City, &site.State, &site.Country)
	if err != nil {
		fmt.Println("Nenhum registro encontrado para o IP:", ip)
		return
	}
	fmt.Printf("Encontrado: ID: %d, Site: %s, IP: %s, Hostname: %s, Date: %s, Provedor: %s, City: %s, State: %s, Country: %s\n",
		site.ID, site.Site, site.IP, site.Hostname, site.Date.Format("2006-01-02 15:04:05"), site.Provedor, site.City, site.State, site.Country)
}

// Atualizar um registro
func UpdateAccessSite(pool *pgxpool.Pool, id int, site, ip, hostname, provedor, city, state, country string) {
	sql := `UPDATE table_access_sites
            SET site = $1, ip = $2, hostname = $3, provedor = $4, city = $5, state = $6, country = $7
            WHERE id = $8`

	_, err := pool.Exec(context.Background(), sql, site, ip, hostname, provedor, city, state, country, id)
	if err != nil {
		log.Fatal("Erro ao atualizar o registro: ", err)
	}
	fmt.Println("Registro atualizado com sucesso:", id)
}

// Deletar um registro
func DeleteAccessSite(pool *pgxpool.Pool, id int) {
	sql := `DELETE FROM table_access_sites WHERE id = $1`

	_, err := pool.Exec(context.Background(), sql, id)
	if err != nil {
		log.Fatal("Erro ao deletar o registro: ", err)
	}
	fmt.Println("Registro deletado com sucesso:", id)
}
