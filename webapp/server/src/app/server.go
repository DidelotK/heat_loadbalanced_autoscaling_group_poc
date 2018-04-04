package main

import (
	// Natives dependencies
	"fmt"
	"log"
	"net/http"
	"os/exec"
	// Other dependencies
	"github.com/gorilla/mux"
	"strings"
	"encoding/json"
	"strconv"
	"regexp"
	"io/ioutil"
	"time"
	"os"
)


func main() {
	const port = "8080"

	router := mux.NewRouter()
	router.HandleFunc("/api/cpu", getCpu)
	router.HandleFunc("/api/ip", getIp)
	router.HandleFunc("/api/makeCpuLoad", makeCpuLoad)
  router.PathPrefix("/").Handler(http.FileServer(http.Dir("../../../front/build")))
  
  http.Handle("/", router)
  fmt.Println("Server running on: " + port)
  log.Fatal(http.ListenAndServe(":" + port, nil))
}
 
func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range(lines) {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

func getCpu(w http.ResponseWriter, r *http.Request) {
	idle0, total0 := getCPUSample()
	time.Sleep(3 * time.Second)
	idle1, total1 := getCPUSample()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	// Préparation de la réponse
	resJson, encodeErr := json.Marshal(cpuUsage)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		fmt.Println("GET_CPU: ENCODE ERROR : " + encodeErr.Error())
		return
	}

	// Envoie de la réponse
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}

func getIp(w http.ResponseWriter, r *http.Request) {

	ip:= os.Getenv("SERVER_IP")

	// Si l'ip n'est pas dans les variables d'environnement
	if ip == "" {
		// Execution de la commande pour récuperer l'ip de la machine
		out, commandErr :=  exec.Command("sh", "-c", "ip addr show eth0 | sed '3q;d' | awk '{print $2}'").Output()
		if commandErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			fmt.Println("GET_IP: COMMAND ERROR")
			return
		}

		// Extraction de l'ip du résultat (ip/mask)
		var re = regexp.MustCompile(`(.*)/.*`)
		ip = re.ReplaceAllString(string(out), `$1`)
		ip = strings.TrimSuffix(ip, "\n")
	}

	// Préparation de la réponse
	resJson, encodeErr := json.Marshal(ip)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		fmt.Println("GET_IP: ENCODE ERROR")
		return
	}

	// Envoie de la réponse
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}

func makeCpuLoad(w http.ResponseWriter, r *http.Request) {
	// Récupération du temps à générer la charge
	durationString := r.URL.Query().Get("duration")

	// Check if the duration is given and is a number
	duration, fetchParamErr := strconv.Atoi(durationString)
	if fetchParamErr != nil || duration == 0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		fmt.Println("MAKE_CPU_LOAD: BAD REQUEST")
		return
	}

	// Execution de la commande pour générer de la charge (pendant 5s)
	cmd :=  exec.Command("sh", "-c", "stress-ng --cpu 8 --cpu-ops 900000 --timeout " + durationString)
	commandErr := cmd.Start()
	if commandErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		fmt.Println("MAKE_CPU_LOAD: COMMAND ERROR : " + commandErr.Error())
		return
	}

	// Préparation de la réponse
	resJson, encodeErr := json.Marshal("OK")
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		fmt.Println("MAKE_CPU_LOAD: ENCODE ERROR : " + encodeErr.Error())
		return
	}
	// Envoie de la réponse
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}
