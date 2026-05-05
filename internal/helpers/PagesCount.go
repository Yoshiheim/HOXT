package helpers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func PagesCount(w http.ResponseWriter, r *http.Request, id *uint64, preid, nextid *uint32) {
	page := r.FormValue("page")

	if id == nil || preid == nil || nextid == nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	if *id > uint64(5000) {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if *id == 0 {
		*id = 1
	}

	if page != "" {
		var err error

		*id, err = strconv.ParseUint(page, 10, 64)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}
		if *id > 1 {
			*preid = uint32(*id - 1)
		}
	}
	*nextid = uint32(*id + 1)
}

func ParsePageID(r *http.Request) (id uint32, err error) {
	pageStr := r.FormValue("page")
	const defaultPage = 1

	if pageStr == "" {
		id = defaultPage
	} else {
		val, err := strconv.ParseUint(pageStr, 10, 32)
		if err != nil {
			return 0, fmt.Errorf("invalid page")
		}
		id = uint32(val)
	}

	return id, nil
}

func SafeParsePage(r *http.Request) (id, preid, nextid uint32, err error) {
	const defaultPage = 1

	pageStr := r.FormValue("page")

	if pageStr == "" {
		id = defaultPage
	} else {
		val, err := strconv.ParseUint(pageStr, 10, 32)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid page")
		}
		id = uint32(val)
	}

	if id == 0 {
		return 0, 0, 1, nil
	}

	if id > 1 {
		preid = id - 1
	}

	if id < math.MaxUint32 {
		nextid = id + 1
	} else {
		nextid = id
	}

	return id, preid, nextid, nil
}
