/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the server using the UDP protocol",
	Long:  `You can send messages to the UDP server, and the server responds back.`,
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

	sendMessage(connection)
	defer connection.Close()

}

func sendMessage(connection *net.UDPConn) {
	for {
		//Write a message
		fmt.Println("Type a message: ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Failed to retreive user input: ", err.Error())
		}
		_, err = connection.Write([]byte(line))
		if err != nil {
			log.Fatal("Error: ", err.Error())
		}

		//Make a receiver and get the message handler server response
		recv := make([]byte, 1024)
		n, _, err := connection.ReadFromUDP(recv)
		if err != nil {
			log.Fatal("Error: ", err.Error())
		}

		fmt.Println(string(recv[:n]))
	}
}
