package config

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	SecretKey          []byte
	Version            = "0.0.1"
)

// Inicializar as variaveis de ambiente
func Carregar() {
	setEnvVariable()

	if erro := godotenv.Load(".env"); erro != nil {
		log.Fatal(erro)
	}

	var erro error

	Porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil {
		Porta = 8181
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

func setEnvVariable() {
	chave := make([]byte, 64)

	if _, erro := godotenv.Read(".env"); erro != nil {
		if _, erro := rand.Read(chave); erro != nil {
			log.Fatal(erro)
		}

		mapValues := make(map[string]string)

		stringBase64 := base64.StdEncoding.EncodeToString(chave)

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Informar usuário do banco de dados: ")
		scanner.Scan()
		mapValues["DB_USUARIO"] = scanner.Text()

		fmt.Print("Informar senha do usuário do banco de dados: ")
		scanner.Scan()
		mapValues["DB_SENHA"] = scanner.Text()

		fmt.Print("Informar nome da base no banco de dados: ")
		scanner.Scan()
		mapValues["DB_NOME"] = scanner.Text()

		fmt.Print("Informar porta que a API ficará executado: ")
		scanner.Scan()
		mapValues["API_PORTA"] = scanner.Text()

		mapValues["SECRET_KEY"] = stringBase64

		newDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		godotenv.Write(mapValues, newDir+"/.env")
	}
}
