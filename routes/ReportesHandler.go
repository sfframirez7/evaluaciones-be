package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
	"github.com/gorilla/mux"
)

func GetResultadosEvaluacionesHistoricasHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetResultadosEvaluacionesHistoricas(idColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	fmt.Fprint(w, string(response))
}

func ColaboradoresPendientesCompletarEvaluacionHanlder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idEvaluacionAnual := vars["idEvaluacionAnual"]

	var Resultados, erro = services.ColaboradoresPendientesCompletarEvaluacion(idEvaluacionAnual)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	fmt.Fprint(w, string(response))
}

func GetRptCompetencias(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetRptCompetencias(idColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	fmt.Fprint(w, string(response))
}

func GetRptResumenGeneralHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idEvaluacionAnual := vars["idEvaluacionAnual"]

	var Resultados, erro = services.RptResumenGeneralService(idEvaluacionAnual)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	fmt.Fprint(w, string(response))
}

func GetRptResumenGeneralPorEquipoHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idEvaluacionAnual := vars["idEvaluacionAnual"]

	token, _ := services.GetToken(r)
	var Resultados, erro = services.RptResumenGeneralPorEquipoService(idEvaluacionAnual, token.Usuario.IdColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	fmt.Fprint(w, string(response))
}

func GetRptReseteoDeNotaHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pagina := vars["pagina"]

	var Resultados, erro = services.RptReseteoNotaService(pagina)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	fmt.Fprint(w, string(response))
}
