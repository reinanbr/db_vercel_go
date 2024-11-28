package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"time"
	"routes_api_go/db"
	"routes_api_go/models"
)

type ResponseIndex struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type Response struct {
	Saved bool `json:"saved"`
}

type RequestData struct {
	IP        string `json:"ip"`
	SiteName  string `json:"site_name"`
	Hostname  string `json:"hostname"`
	Datetime  string `json:"datetime"`
	Country   string `json:"country"`
	State     string `json:"state"`
	City      string `json:"city"`
	Provedor  string `json:"provedor"`
}

// Função para validar IP
func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// Função para validar site_name (exemplo simples de URL)
func isValidSiteName(siteName string) bool {
	re := regexp.MustCompile(`^(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)
	return re.MatchString(siteName)
}

// Função para validar hostname (não permite caracteres especiais além de "-")
func isValidHostname(hostname string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9-]{1,63}$`)
	return re.MatchString(hostname)
}

// Função para validar datetime (usando RFC3339 para formato ISO 8601)
func isValidDatetime(datetime string) bool {
	_, err := time.Parse(time.RFC3339, datetime)
	return err == nil
}

// Função para validar campos de texto (Country, State, City e Provedor)
func isValidTextField(field string) bool {
	return len(field) > 0 && len(field) <= 100
}




func handler(w http.ResponseWriter, r *http.Request) {
	 pool :=psql_vercel.ConnectDB()
         defer pool.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	// Validações
	if !isValidIP(data.IP) {
		http.Error(w, "IP inválido", http.StatusBadRequest)
		return
	}
	if !isValidSiteName(data.SiteName) {
		http.Error(w, "Site name inválido", http.StatusBadRequest)
		return
	}
	if !isValidHostname(data.Hostname) {
		http.Error(w, "Hostname inválido", http.StatusBadRequest)
		return
	}
	if !isValidDatetime(data.Datetime) {
		http.Error(w, "Datetime inválido, use formato ISO 8601 (RFC3339)", http.StatusBadRequest)
		return
	}
	if !isValidTextField(data.Country) {
		http.Error(w, "Country inválido", http.StatusBadRequest)
		return
	}
	if !isValidTextField(data.State) {
		http.Error(w, "State inválido", http.StatusBadRequest)
		return
	}
	if !isValidTextField(data.City) {
		http.Error(w, "City inválido", http.StatusBadRequest)
		return
	}
	if !isValidTextField(data.Provedor) {
		http.Error(w, "Provedor inválido", http.StatusBadRequest)
		return
	}

	res := access_site_model.CreateAccessSite(pool,data.SiteName, 
							data.IP, 
							data.Hostname, 
							data.Datetime, 
							data.Provedor,
							data.City, 
							data.State, 
								data.Country);
	response := Response{
		Saved:false,
	}
	if(res.Success==1){
		response = Response{
			Saved:true,
		}
	}
	fmt.Fprintf(w, "Dados recebidos com sucesso: %+v\n", response)
}



func read(w http.ResponseWriter,r*http.Request){
	pool :=psql_vercel.ConnectDB()
        defer pool.Close()
	infoAccess,err := access_site_model.ReadAccessSites(pool)
	if err == nil{
		w.Header().Set("Content-Type", "application/json")
		if errJson := json.NewEncoder(w).Encode(infoAccess); err != nil {
			http.Error(w, "Erro ao gerar o JSON", http.StatusInternalServerError)
			log.Printf("Erro ao codificar JSON: %v", errJson)
	}
	}else{
		fmt.Fprintf(w,"error: %v\n",err)
	}
}



func index(w http.ResponseWriter, r *http.Request) {
	response := ResponseIndex{
		Message:   "Estamos online",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func main() {
	http.HandleFunc("/receive-data", handler)
	http.HandleFunc("/read",read)
	http.HandleFunc("/",index)
	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

