package postcode

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

// HStatApi - выдача статистики по системе.
func HStatApi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	Rec, err = GetStat()
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}
	Rec.Version		= Version
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(Rec); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}

func HGetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	gist			:= strings.ToLower(GetVal(w, ps, "name")) 
	if isValueInTable(gist) < 0 {
	    w.WriteHeader(404)
	    return
	}
	Out, err		:= GetGist(gist, 0)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(Out); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}

func HGetAllIndexes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	gist			:= strings.ToLower(GetVal(w, ps, "name")) 
	tLevel			:= isValueInTable(gist)
	if  tLevel < 0 {
	    w.WriteHeader(404)
	    return
	}
	id			:= GetVal(w, ps, "id")
	_, err 			= strconv.Atoi(id)
	if  err != nil {
	    w.WriteHeader(500)
	    return
	}
	Out, err		:= GetAllIndexes(gist, tLevel, id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(Out); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}

func HGetAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	id			:= GetVal(w, ps, "id")
	_, err 			= strconv.Atoi(id)
	if  err != nil {
	    w.WriteHeader(500)
	    return
	}

	Out, err		:= GetAddress(id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(Out); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}

func HGetAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    var err error
    var indexes Indexes

    err = json.NewDecoder(r.Body).Decode(&indexes)
    if err != nil {
	w.WriteHeader(500)
	fmt.Println(err)
	return
    }
    for i:=0; i<len(indexes.Indexes); i++{
	_, err 			:= strconv.Atoi(indexes.Indexes[i])
	if  err != nil {
	    w.WriteHeader(500)
	    return
	}
    }
    Out, err		:= GetAddresses(indexes)
    if err != nil {
	if err == sql.ErrNoRows {
	    fmt.Println(err)
	    w.WriteHeader(404)
	    return
	}
	w.WriteHeader(500)
	fmt.Println(err)
	return
    }
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    if err = json.NewEncoder(w).Encode(Out); err != nil {
	fmt.Println(err)
	w.WriteHeader(500)
    }
}

func HGetNamesByTopID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	gist			:= strings.ToLower(GetVal(w, ps, "name")) 
	if isValueInTable(gist) < 0 {
	    w.WriteHeader(404)
	    return
	}
	id			:= GetVal(w, ps, "id")
	k, err 			:= strconv.Atoi(id)
	if  err != nil {
	    w.WriteHeader(500)
	    return
	}
	Out, err		:= GetGist(gist, k)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(Out); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}

func isValueInTable(value string) int {
    for k, v := range table {
        if v == value {
            return k
        }
    }
    return -1
}