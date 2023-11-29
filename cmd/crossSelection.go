/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	calculation "github.com/victor8titov/wireCalc/internal"
)

// crossSelectionCmd represents the crossSelection command
var crossSelectionCmd = &cobra.Command{
	Use:   "cs [diameter]",
	Short: "Вычисление сечения провода по его диаметру (wire cross selection)",
	Long: `Вычисление провода как для одножильных так
	 и для многожильных медных проводов
	`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		diameter := args[0]
		if len(diameter) == 0 {
			fmt.Println("Нужен диаметр для вычислений")
			return
		}
		d, err := strconv.ParseFloat(diameter, 64)
		if err != nil {
			fmt.Println("Error during parse")
		}

		cs := calculation.CrossSectionByDiameter(d)
		fmt.Println("Сечение провода: ", cs)
	},
}

func init() {
	rootCmd.AddCommand(crossSelectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crossSelectionCmd.PersistentFlags().Float64P("diameter", "d", 0.0, "Диаметр провода")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crossSelectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
