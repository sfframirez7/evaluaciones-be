package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
	"github.com/gorilla/mux"
)

func GetEvaluacionPorColaborador(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]
	idEvaluacionAnual := vars["idEvaluacionAnual"]

	var Resultados, erro = services.GetEvaluacionPorColaborador(idColaborador, idEvaluacionAnual)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetEvaluacionAnual(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetEvaluacionAnual(idColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetEvaluacionesCompletas(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetEvaluacionesCompletas(idColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func EvaluacionCompletadaHandler(w http.ResponseWriter, r *http.Request) {

	var evaluacionCompletada models.EvaluacionCompletada
	err := json.NewDecoder(r.Body).Decode(&evaluacionCompletada)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Datos incorrectos")
		return
	}

	var respuesta, erro = services.GuardarEvaluacionCompletada(evaluacionCompletada)

	if erro != nil || respuesta == false {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&respuesta)

	fmt.Fprintf(w, string(response))
}

func NuevaEvaluacionPorMetaHandler(w http.ResponseWriter, r *http.Request) {

	var evaluacionPorMeta models.NuevaEvaluacionPorMeta
	err := json.NewDecoder(r.Body).Decode(&evaluacionPorMeta)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	var respuesta, erro = services.NuevaEvaluacionPorMeta(evaluacionPorMeta)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&respuesta)

	fmt.Fprintf(w, string(response))
}

func GetEvaluacionMetaPorColaborador(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]
	idPadre := vars["idPadre"]

	var Resultados, erro = services.GetEvaluacionMetaPorColaborador(idColaborador, idPadre)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)
	fmt.Fprint(w, string(response))
}

func GetEvaluacionesTodasHanlder(w http.ResponseWriter, r *http.Request) {

	var Resultados, erro = services.GetEvaluacionsTodas()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)
	fmt.Fprint(w, string(response))
}

func AceptarEvaluacionHanlder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idEvaluacion := vars["idEvaluacion"]

	var Resultados, erro = services.AceptarEvaluacion(idEvaluacion)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)
	fmt.Fprint(w, string(response))
}

func NuevaEvaluacionAnualHandler(w http.ResponseWriter, r *http.Request) {

	var evaluacion models.EvaluacionAnual
	err := json.NewDecoder(r.Body).Decode(&evaluacion)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Datos incorrectos")
		return
	}

	var respuesta, erro = services.NuevaEvaluacionAnual(evaluacion)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&respuesta)

	fmt.Fprintf(w, string(response))
}

func EliminarEvaluacionPorMetaHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	IdEvaluacionMeta := vars["IdEvaluacionMeta"]

	var respuesta, erro = services.EliminarEvaluacionPorMetaService(IdEvaluacionMeta)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&respuesta)

	fmt.Fprintf(w, string(response))
}

func ResetearNotaEvaluacionPorMetaHandler(w http.ResponseWriter, r *http.Request) {

	var resetearNota models.ResetearNotaEvaluacion
	err := json.NewDecoder(r.Body).Decode(&resetearNota)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	token, _ := services.GetToken(r)
	resetearNota.EliminadaPor = token.Usuario.IdColaborador

	var respuesta, erro = services.ResetearNotaEvaluacionPorMetaService(resetearNota)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&respuesta)

	fmt.Fprintf(w, string(response))
}

func ResetearNotaEvaluacionGeneralHandler(w http.ResponseWriter, r *http.Request) {

	var resetearNota models.ResetearNotaEvaluacion
	err := json.NewDecoder(r.Body).Decode(&resetearNota)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	token, _ := services.GetToken(r)
	resetearNota.EliminadaPor = token.Usuario.IdColaborador

	var respuesta, erro = services.ResetearNotaEvaluacionGeneralService(resetearNota)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&respuesta)

	fmt.Fprintf(w, string(response))
}
