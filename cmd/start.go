package cmd

import (
	"fmt"
	"log"
	"strings"

	v1 "github.com/adityaputra11/mojek/internal/api/v1"
	"github.com/adityaputra11/mojek/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/cobra"
)

var (
	configpath string
	startCmd   = &cobra.Command{
		Use:   "start",
		Short: "start server",
		Long:  `start server, default port is 5000`,
		Run:   startServer,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVarP(&configpath, "config", "c", "", "config file (default is $HOME/.serverd/default.yaml)")
	startCmd.PersistentFlags().Int("port", 5000, "Port to run Application server on")
	config.Viper().BindPFlag("port", startCmd.PersistentFlags().Lookup("port"))
}

func initConfig() {
	config.Viper().SafeWriteConfig()
	config.Viper().WriteConfigAs("$HOME/.serverd/.config")
	if len(configpath) != 0 {
		config.Viper().SetConfigFile(configpath)
	} else {
		config.Viper().AddConfigPath("$HOME/.serverd/")
		config.Viper().AddConfigPath("./config")
		config.Viper().SetConfigName("default")
	}
	config.Viper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.Viper().AutomaticEnv()
	if err := config.Viper().ReadInConfig(); err != nil {
		log.Fatalf("Using config file [%s]: %v", config.Viper().ConfigFileUsed(), err)
	}
	fmt.Println("Config paths:", config.Viper().ConfigFileUsed())
	fmt.Println("DBType:", config.Viper().GetString("database.type"), len(config.Viper().GetString("database.url")))
}

func startServer(cmd *cobra.Command, agrs []string) {
	fmt.Println("Start http-server")
	fmt.Println(config.Viper().GetString("database.url"))
	db, err := gorm.Open(config.Get("database.type"), config.Get("database.url"))
	if err != nil {
		panic(fmt.Errorf("Failed to connect database: %w", err))
	}
	defer db.Close()
	router := gin.Default()
	apiV1Router := router.Group("/api/v1")
	v1.RegisterRouterAPIV1(apiV1Router, db)
	router.Run(":" + config.Get("port"))
}
