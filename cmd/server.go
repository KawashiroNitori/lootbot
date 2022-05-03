package cmd

import (
	"github.com/KawashiroNitori/lootbot/internal/bot"
	"github.com/KawashiroNitori/lootbot/internal/http/router"
	"github.com/KawashiroNitori/lootbot/internal/service"
	"github.com/KawashiroNitori/lootbot/internal/util"
	"github.com/lonelyevil/khl"
	"github.com/lonelyevil/khl/log_adapter/plog"
	"github.com/phuslu/log"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动服务器",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runServer(cmd *cobra.Command, args []string) {
	logger := log.Logger{
		Level:  log.TraceLevel,
		Writer: &log.ConsoleWriter{},
	}
	logger.Info().Msg("Starting KHL connection...")
	s := khl.New("1/MTEwMTk=/keM97tT/TxI1BschdnZONA==", plog.NewLogger(&logger))

	lootChatHandler := bot.NewLootChatHandler()
	s.AddHandler(util.Recover(lootChatHandler.MessageIn))

	if err := s.Open(); err != nil {
		logger.Err(err).Msg("Open KHL connection failed")
		os.Exit(1)
	}
	service.DefaultKHLSession = s
	defer func() {
		_ = s.Close()
	}()

	logger.Info().Msg("Starting HTTP Server...")
	r := router.NewRouter()
	_ = http.ListenAndServe(":5000", r)
}
