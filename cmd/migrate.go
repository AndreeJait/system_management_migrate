package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
	"system_mananagement_migrate/commons/utils"
	"system_mananagement_migrate/config"
	"system_mananagement_migrate/shared/constant"
)

const (
	ModeFlag   = "mode"
	ModuleFlag = "module"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("use -h to show available commands")
	},
}

var migrateDownCmd = &cobra.Command{
	Use: "down",
	Run: func(cmd *cobra.Command, args []string) {
		processMigrate(cmd, constant.MigrateDown)
	},
}

var migrateUpCmd = &cobra.Command{
	Use: "up",
	Run: func(cmd *cobra.Command, args []string) {
		processMigrate(cmd, constant.MigrateUp)
	},
}

var migrateFreshCmd = &cobra.Command{
	Use: "fresh",
	Run: func(cmd *cobra.Command, args []string) {
		processMigrate(cmd, constant.MigrateFresh)
	},
}

func processMigrate(cmd *cobra.Command, migrateType constant.MigrateType) {
	modeApp, _ := cmd.Flags().GetString(ModeFlag)
	module, _ := cmd.Flags().GetString(ModuleFlag)
	envMode := config.GetEnvMode(modeApp)

	cfg := config.LoadConfig(envMode)
	moduleApp := constant.GetMigrateDB(module)

	var err error
	switch migrateType {
	case constant.MigrateDown:
		err = utils.MigrateDown(*cfg, moduleApp)
		break
	case constant.MigrateFresh:
		err = utils.MigrateFresh(*cfg, moduleApp)
		break
	default:
		err = utils.MigrateUp(*cfg, moduleApp)
		break
	}
	if err != nil {
		logrus.Fatal(err)
	}
}

func Run() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	migrateUpCmd.Flags().String(ModeFlag, "dev", "to define mode app")
	migrateUpCmd.Flags().String(ModuleFlag, "", "to define module app")
	migrateCmd.AddCommand(migrateUpCmd)

	migrateFreshCmd.Flags().String(ModeFlag, "dev", "to define mode app")
	migrateFreshCmd.Flags().String(ModuleFlag, "", "to define module app")
	migrateCmd.AddCommand(migrateFreshCmd)

	migrateDownCmd.Flags().String(ModeFlag, "dev", "to define mode app")
	migrateDownCmd.Flags().String(ModuleFlag, "", "to define module app")
	migrateCmd.AddCommand(migrateDownCmd)

	err := migrateCmd.Execute()
	if err != nil {
		logrus.Error(err)
	}
}
