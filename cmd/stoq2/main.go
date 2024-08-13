package main

import (
	"crud/internal/config"
	"encoding/json"
	"fmt"
	"os"
)

// Função para fazer configuração com um arquivo
func main(){

	// Será alocado essa estrutura na memória e ele vai passar para mim o lugar que está alocado
	default_config := &config.Config{}

	// Esse if é diferente pois a variável "file_config" só consegue ser usada dentro do escopo do if
	if file_config := os.Getenv("STOQ2_CONFIG"); file_config != ""{
		file, err := os.ReadFile(file_config)
		if err != nil {
			fmt.Println(err.Error())
			// log.Panicln(err.Error()) faz com que mate o sistema se cair nesse IF
		}

		// Não terá tratamento de erro nesse caso pq o sistema não precisa necessariamente de um arquivo para fazer a configuração
		json.Unmarshal(file, default_config)

	}

	conf := config.NewConfig(default_config)

	fmt.Println(conf)

}