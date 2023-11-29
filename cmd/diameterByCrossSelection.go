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

// diameterByCrossSelection represents the d command
var diameterByCrossSelection = &cobra.Command{
	Use:   "d [сечение в мм]",
	Short: "Определение диаметра по сечению  провода",
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		cs := args[0]
		crossSelection, err := strconv.ParseFloat(cs, 64)
		if err != nil {
			fmt.Println("Error during parse")
		}

		diameter := calculation.DiameterByCrossSelection(crossSelection)
		fmt.Println("Диаметр провода: ", diameter)
	},
}

func init() {
	rootCmd.AddCommand(diameterByCrossSelection)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
