package application

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"encoding/json"
	"net/http"

	"github.com/MrM2025/rpforcalc/tree/master/calc_go/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

// Функция запуска приложения
// тут будем чиать введенную строку и после нажатия ENTER писать результат работы программы на экране
// если пользователь ввел exit - то останаваливаем приложение
func (a *Application) Run() error {
	var (
		calc calculation.TCalc
	)
	for {
		// читаем выражение для вычисления из командной строки
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		// убираем пробелы, чтобы оставить только вычислемое выражение
		text = strings.TrimSpace(text)
		// выходим, если ввели команду "exit"
		if text == "exit" {
			log.Println("aplication was successfully closed")
			return nil
		}
		//вычисляем выражение
		result, err := calc.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed wit error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

type CalcReqJSON struct {
	Expression string `json:"expression"`
}

type CalcResJSON struct{
	Result string `json:"result,omitempty"`
	Error string `json:"error,omitempty"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	var (
		calc calculation.TCalc
		status int
		emsg string
	)

	request := new(CalcReqJSON)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	result, err := calc.Calc(request.Expression)
	if err != nil {

		switch {	
		case errors.Is(err, calculation.EmptyExpressionErr):
			status = http.StatusUnprocessableEntity
			emsg = calculation.EmptyExpressionErr.Error()

		case errors.Is(err, calculation.IncorrectExpressionErr):
			status = http.StatusUnprocessableEntity
			emsg = calculation.IncorrectExpressionErr.Error()
 
		case errors.Is(err, calculation.NumToPopMErr): // numtopop > nums' slise length
			status = http.StatusUnprocessableEntity
			emsg = calculation.NumToPopMErr.Error()

		case errors.Is(err, calculation.NumToPopZeroErr): // numtopop <= 0
			status = http.StatusUnprocessableEntity
			emsg = calculation.NumToPopZeroErr.Error()

		case errors.Is(err, calculation.NthToPopErr): // no operator to pop
			status = http.StatusUnprocessableEntity
			emsg = calculation.NthToPopErr.Error()

		case errors.Is(err, calculation.DvsByZeroErr):
			status = http.StatusUnprocessableEntity
			emsg = calculation.DvsByZeroErr.Error()
		}

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(CalcResJSON{Error: emsg})

	} else if err := recover(); err != nil {
		log.Printf("Panic: %v", err)
		status = http.StatusInternalServerError
		emsg = "unknown error"
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(CalcResJSON{Error: emsg})
	} else {
			log.Printf("Successful calculation: %s = %f", request.Expression, result)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(CalcResJSON{Result: fmt.Sprintf("%f", result)})
		}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}

