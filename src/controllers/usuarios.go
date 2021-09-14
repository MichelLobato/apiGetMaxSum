package controllers

import (
	"api/src/modelos"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//MaxSum Recebe uma lista em formato Json e passa para Struct
func MaxSum(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var num modelos.Resp
	if erro = json.Unmarshal(corpoRequest, &num); erro != nil {
		log.Fatal(erro)
	}

	var RespJSON modelos.RespJSON
	RespJSON.Positions, RespJSON.Sum = KadaneMaxSum(num.Lista)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, RespJSON)
}

//KadaneMaxSum Recebe uma lista e usando o Algoritmo de Kadane retorna o array de maior soma e seu resultaddo
func KadaneMaxSum(tArray []int) ([]int, int) {

	mStartIdx := 0
	mEndIdx := 1

	startIdx := 0
	endIdx := 1

	sum := tArray[0]
	maxiSum := tArray[0]

	for i := 1; i < len(tArray); i++ {
		if tArray[i] > tArray[i]+sum {
			sum = tArray[i]
			startIdx = i
			endIdx = i + 1
		} else {
			sum += tArray[i]
			endIdx = i + 1
		}

		if sum > maxiSum {
			maxiSum = sum
			mStartIdx = startIdx
			mEndIdx = endIdx
		}
	}

	rArray := make([]int, mEndIdx-mStartIdx)
	copy(rArray, tArray[mStartIdx:mEndIdx])
	fmt.Println(mStartIdx, mEndIdx)

	return rArray, maxiSum
}
