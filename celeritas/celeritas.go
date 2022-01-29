package celeritas

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const pp = 0
const version = "1.0.1"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	Errorlog *log.Logger
	Infolog  *log.Logger
	RootPath string
	config   config
	Routes   *chi.Mux
}
type config struct {
	port     string
	renderer string
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := c.Init(pathConfig)
	if err != nil {
		return err
	}
	err = c.checkDotEnv(rootPath)

	if err != nil {
		return err
	}
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}
	infoLog, errorLog := c.startLoggers()
	c.Infolog = infoLog
	c.Errorlog = errorLog
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
	c.Routes = c.routes().(*chi.Mux)
	c.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}
	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folderNames
		if _, err := os.Stat(root + "/" + path); os.IsNotExist(err) {
			err := os.Mkdir(root+"/"+path, 0755)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (c *Celeritas) checkDotEnv(path string) error {
	if _, err := os.Stat(path + "/.env"); os.IsNotExist(err) {
		a, err := os.Create(path + "/.env")
		if err != nil {
			return err
		}
		defer func(a *os.File) {
			err := a.Close()
			if err != nil {
				return
			}
		}(a)
	}
	return nil
}
func (c *Celeritas) startLoggers() (*log.Logger, *log.Logger) {
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}
func (c *Celeritas) ListenAndServe() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     c.Errorlog,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 100 * time.Second,
		Handler:      c.routes(),
	}
	c.Infolog.Printf("Starting %s on port %s", c.AppName, os.Getenv("PORT"))
	err := srv.ListenAndServe()
	if err != nil {
		c.Errorlog.Fatal(err)
	}
	return nil
}
