package save_access_site_controller




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





		
