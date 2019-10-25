package main

import "github.com/spf13/cobra"

// client flag
var(
	DeckAddr = "http://127.0.0.1:8080"
)

func main(){
	RootCMD := &cobra.Command{Short: "deck command line tool."}
	RootCMD.PersistentFlags().StringVarP(&DeckAddr, "deck-addr", "d", DeckAddr, "deck server addr")
	RootCMD.AddCommand(NewComponentCommand())
	RootCMD.Execute()
}