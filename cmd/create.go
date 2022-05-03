package cmd

import (
	"context"
	"fmt"
	"github.com/KawashiroNitori/lootbot/internal/dao"
	"github.com/KawashiroNitori/lootbot/internal/service"
	"github.com/spf13/cobra"
	"os"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "创建新数据",
	Args:  cobra.MinimumNArgs(1),
	Run:   run,
}

var (
	channelID string
	partyID   string
	csvPath   string
	needClear bool
)

func init() {
	RootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&channelID, "channel", "c", "", "小队播报所在的开黑啦频道 ID")
	createCmd.Flags().StringVarP(&partyID, "party", "p", "", "要导入的小队 ID")
	createCmd.Flags().StringVarP(&csvPath, "input", "i", "", "要导入的 CSV 文件路径")
	createCmd.Flags().BoolVar(&needClear, "clear", false, "导入前是否清除原有记录")
}

func run(cmd *cobra.Command, args []string) {
	arg := args[0]
	switch arg {
	case "party":
		createParty(cmd, args)
	case "loot":
		createLoot(cmd, args)
	default:
		fmt.Printf("invalid subcommand: %s", arg)
		os.Exit(1)
	}
}

func createParty(cmd *cobra.Command, args []string) {
	if channelID == "" {
		fmt.Println("需要频道 ID")
		os.Exit(1)
	}
	ctx := context.Background()
	partyDAO := dao.DefaultPartyDAO
	pt, err := partyDAO.CreateParty(ctx, channelID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("小队 ID: %s", pt.ID)
}

func createLoot(cmd *cobra.Command, args []string) {
	if partyID == "" || csvPath == "" {
		fmt.Println("需要小队 ID 或 CSV 导入文件")
		os.Exit(1)
	}
	service.ImportLootsFromCSV(context.Background(), csvPath, partyID, needClear)
}
