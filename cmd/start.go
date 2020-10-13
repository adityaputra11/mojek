package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/adityaputra11/mojek/internal/config"
	"github.com/gin-gonic/gin"
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
	// dsn := config.Viper().GetString("database.url")
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer fmt.Println(db)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.Run(":" + config.Get("port"))
}
