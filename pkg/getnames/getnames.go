package getnames

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/anditakaesar/uwa-logic/constants"
	"github.com/tidwall/buntdb"
)

const (
	NameKey      string = "agent:%s:name"
	NextKey      string = "NEXTKEY"
	NextKeyValue string = "%d:%d:%d"
	LenNames     int    = 10
	CounterKey   string = "COUNTER"
)

func Run() {
	fmt.Println("[main] initialize db")
	db, err := setupDB()
	if err != nil {
		log.Fatalf("[main] buntdb.Open err %s", err.Error())
	}
	defer db.Close()

	fmt.Println("[main] initialize main key")
	err = InitializeData(db)
	if err != nil {
		log.Fatalf("[main] SetupDB err %s", err.Error())
	}

	log.Fatal(srv(db))
}

func setupDB() (*buntdb.DB, error) {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitializeData(db *buntdb.DB) error {
	names := constants.GetNames(LenNames)
	err := db.Update(func(tx *buntdb.Tx) error {
		for i, name := range names {
			_, _, err := tx.Set(fmt.Sprintf(NameKey, strconv.Itoa(i)), fmt.Sprintf("%s:%d", name, 0), nil)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Println("[main][SetupDB] init names err")
		return err
	}

	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(NextKey, fmt.Sprintf(NextKeyValue, 0, len(names)-1, 0), nil)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("[main][SetupDB] init first key err")
		return err
	}

	return nil
}

func GetCurrentName(db *buntdb.DB) (string, error) {
	currentName := ""
	err := db.Update(func(tx *buntdb.Tx) error {
		nextKey, err := tx.Get(NextKey)
		if err != nil {
			log.Println("[main][GetNextValue] get nextkey err")
			return err
		}
		currentAndMax := strings.Split(nextKey, ":")
		currentKey := fmt.Sprintf(NameKey, currentAndMax[0])
		currentNameAndCountData, err := tx.Get(currentKey)
		if err != nil {
			log.Println("[main][GetNextValue] get nextname err")
			return err
		}
		currentNameAndCount := strings.Split(currentNameAndCountData, ":")
		currentName = currentNameAndCount[0]
		currentCount, _ := strconv.Atoi(currentNameAndCount[1])

		_, _, err = tx.Set(currentKey, fmt.Sprintf("%s:%d", currentName, currentCount+1), nil)
		if err != nil {
			log.Println("[main][GetNextValue] set count err")
			return err
		}

		currentRobin, _ := strconv.Atoi(currentAndMax[0])
		maxRobin, _ := strconv.Atoi(currentAndMax[1])
		totalCount, _ := strconv.Atoi(currentAndMax[2])

		nextRobin := currentRobin + 1
		if nextRobin > maxRobin {
			nextRobin = 0
		}
		_, _, err = tx.Set(NextKey, fmt.Sprintf(NextKeyValue, nextRobin, maxRobin, totalCount+1), nil)
		if err != nil {
			log.Println("[main][GetNextValue] set nextkey err")
			return err
		}

		return nil
	})
	if err != nil {
		return currentName, err
	}

	return currentName, nil
}

func srv(db *buntdb.DB) error {
	myHandler := MyHandler{
		DB: db,
	}
	http.HandleFunc("/", myHandler.GetNameRobin)
	http.HandleFunc("/counts", myHandler.GetNameCounts)

	return http.ListenAndServe(":8888", nil)
}

type MyHandler struct {
	DB *buntdb.DB
}

func (h *MyHandler) GetNameRobin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "close")
	if r.Method != http.MethodGet {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	requestNumber := r.URL.Query().Get("number")

	nextName, err := GetCurrentName(h.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Printf("[main][srv] got connection: %s: %s", requestNumber, nextName)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"number": requestNumber,
		"result": nextName,
	})
}

func (h *MyHandler) GetNameCounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "close")
	if r.Method != http.MethodGet {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	result := map[string]string{}

	err := h.DB.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			result[key] = value
			return true // continue iteration
		})
		return err
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
