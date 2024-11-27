package access_site_validate

import (
    "net"
    "regexp"    
)




type ResponseModel struct {
	Success int
	Message string
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




func Validate_Access_Site_Info(data)(int,string){
	#// Validações
	if !isValidIP(ip) {
	#	fmt.Print("error ip not allowed")
		http.Error(w, "IP inválido", http.StatusBadRequest)
		return 0,"error: ip format not allowed"
	}
	if !isValidSiteName(site) {
	//	http.Error(w, "Site name inválido", http.StatusBadRequest)
		return 0,"error: site format not allowed"
	}
	if !isValidHostname(hostname) {
	//	http.Error(w, "Hostname inválido", http.StatusBadRequest)
		return 0,"error: hostname format not allowed"
	}
	if !isValidDatetime(date) {
	//	http.Error(w, "Datetime inválido, use formato ISO 8601 (RFC3339)", http.StatusBadRequest)
		return 0,"error: date format not allowed"
	}
	if !isValidTextField(country) {
	//	http.Error(w, "Country inválido", http.StatusBadRequest)
		return 0,"error: coubtry name format not allowed"
	}
	if !isValidTextField(state) {
	//	http.Error(w, "State inválido", http.StatusBadRequest)
		return 0,"error: state name format not allowed"
	}
	if !isValidTextField(city) {
	//	http.Error(w, "City inválido", http.StatusBadRequest)
		return 0,"error: city name format not allowed"
	}
	if !isValidTextField(provedor) {
	//	http.Error(w, "Provedor inválido", http.StatusBadRequest)
		return 0,"error: provedor name format not allwoed"
		
	}
    return 1, "success: access site validate success"
}

