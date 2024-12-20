package main

import (
	"flag"
	"log"

	"github.com/ValeryBMSTU/web-11/internal/auth/api"
	"github.com/ValeryBMSTU/web-11/internal/auth/config"
	"github.com/ValeryBMSTU/web-11/internal/auth/provider"
	"github.com/ValeryBMSTU/web-11/internal/auth/usecase"
	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "../../configs/auth_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(prv)
	srv := api.NewServer(cfg.IP, cfg.Port, use)

	srv.Run()
}
