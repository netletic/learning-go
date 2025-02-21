package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"snippetbox.netletic.com/internal/models"
)

type application struct {
	logger         *slog.Logger
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "", "MySQL data source name")
	flag.Parse()

	// logging
	slogHandlerOpts := &slog.HandlerOptions{
		// AddSource: true, // adds line number where error occurred
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, slogHandlerOpts))

	// set dsn based on env vars if dsn not set via CLI flag
	if *dsn == "" {
		dbUser := os.Getenv("SNIPPETBOX_DB_USER")
		if dbUser == "" {
			logger.Error("SNIPPETBOX_DB_USER environment variable is not set")
			os.Exit(1)
		}
		dbPassword := os.Getenv("SNIPPETBOX_DB_PASSWORD")
		if dbPassword == "" {
			logger.Error("SNIPPETBOX_DB_PASSWORD environment variable is not set")
			os.Exit(1)
		}
		*dsn = fmt.Sprintf("%s:%s@/snippetbox?parseTime=true", dbUser, dbPassword)
	}

	// db
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		logger:         logger,
		snippets:       &models.SnippetModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	// mux
	mux := app.Routes()
	logger.Info("starting server", slog.String("addr", *addr))
	err = http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
