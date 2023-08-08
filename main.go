package main

import (
	"errors"
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const defaultPort int = 8080

type (
	config struct {
		Port int
	}

	database struct {
		m     sync.RWMutex
		tasks map[string]interface{}
	}
)

func (d *database) Add(task string) bool {
	if task == "" {
		return false
	}
	d.m.Lock()
	defer d.m.Unlock()
	if _, ok := d.tasks[task]; !ok {
		d.tasks[task] = nil
		return true
	}
	return false
}

func (d *database) Del(task string) bool {
	d.m.Lock()
	defer d.m.Unlock()
	if _, ok := d.tasks[task]; ok {
		delete(d.tasks, task)
		return true
	}
	return false
}

func (d *database) List() []string {
	d.m.RLock()
	defer d.m.RUnlock()
	return keys(d.tasks)
}

func main() {
	var (
		cfg = config{defaultPort}
		db  = database{tasks: map[string]interface{}{}}
	)
	parseCmd(&cfg)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cfg.Port), setupServer(&db)))
}

func parseCmd(cfg *config) {
	flag.Func(
		"port",
		"specify the port where you want to run the application",
		func(value string) error {
			p, err := strconv.Atoi(value)
			if err != nil {
				return err
			}

			if p < 1023 || p > 30000 {
				return errors.New("port must be between 1023 and 30000")
			}

			cfg.Port = p

			return nil
		},
	)

	flag.Parse()
}

func setupServer(db *database) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(CustomSystem{fs: http.Dir("./web/static/")})),
	)
	mux.HandleFunc("/", index(db))
	return mux
}

func index(db *database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ts, err := template.ParseFiles("./web/tmpl/index.html.gotmpl")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		if r.Method == http.MethodGet {
			err = ts.Execute(w, db.List())
			if err != nil {
				log.Println(err.Error())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		} else if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				log.Println(err.Error())
				http.Error(w, "Error processing request", http.StatusBadRequest)
			}

			if v, ok := r.PostForm["task"]; ok && len(v) == 1 {
				db.Add(strings.TrimSpace(v[0]))
				http.Redirect(w, r, r.URL.String(), http.StatusSeeOther)
			}
		} else if r.Method == http.MethodDelete {
			sanitized, _, ok := strings.Cut(strings.TrimPrefix(r.URL.Path, "/"), "/")
			if ok {
				log.Println("error")
				http.Error(w, "bad request", http.StatusBadRequest)
			}

			db.Del(sanitized)
			w.WriteHeader(http.StatusOK)
		} else {
			w.Header().Set("Allow", strings.Join([]string{http.MethodGet, http.MethodPost, http.MethodDelete}, ", "))
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("method not allowed"))
		}
	}
}

func keys[T comparable, K any](input map[T]K) []T {
	var (
		result = make([]T, len(input))
		i      = 0
	)

	for k := range input {
		result[i] = k
		i++
	}

	return result
}
