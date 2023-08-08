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

const (
	defaultPort int = 8080

	indexHTML string = `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>TODO list</title>
    <style>
      div {
        margin: auto;
        width: 50%;
        display: grid;
        justify-content: center;
        align-items: center;
        margin-top: 100px;
      }
    </style>
</head>
<body>

  <div>
    <ol>
      {{ range $k := . }}
      <li>{{ . }}</li>
      {{- end }}
    </ol>

    <form action="/" method="POST">
      <input type="text" name="task"/>
      <input type="submit" value="Add"/>
    </form>
  </div>

</body>
</html>
`
)

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

	mux.HandleFunc("/", index(db))
	return mux
}

func index(db *database) http.HandlerFunc {
	ts, err := template.New("index").Parse(indexHTML)
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
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
		} else {
			w.Header().Set("Allow", strings.Join([]string{http.MethodGet, http.MethodPost}, ", "))
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
