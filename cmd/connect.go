/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("connect called")
		connectToServer()
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func connectToServer() {
	server, err := net.ResolveUDPAddr("udp", ":8088")
	if err != nil {
		log.Panic("Failed to connect", err.Error())

	}

	connection, err := net.DialUDP("udp", nil, server)
	if err != nil {
		log.Panic("Listen Failed: ", err.Error())
	}

	defer connection.Close()

	_, err = connection.Write([]byte("This is a UDP Message, you fag"))
	if err != nil {
		log.Panic("Error: ", err.Error())
	}

	//Get data to buffer
	recv := make([]byte, 1024)
	_, err = connection.Read(recv)
	if err != nil {
		log.Panic("Error: ", err.Error())
	}

	fmt.Println(string(recv))
}
