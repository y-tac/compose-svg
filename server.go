package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Config 構造体。サブパッケージのconfigも設定する
type Config struct {
	Server ServerConfig `json:"server"`
}

// ServerConfig サーバconfigを設定
type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	fmt.Println(config)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/gen", "gen")

	// サーバー起動
	e.Start(":" + config.Server.Port)
}
