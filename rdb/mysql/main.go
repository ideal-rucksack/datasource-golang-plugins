package main

import (
	"github.com/ideal-rucksack/datasource-golang-plugin/rdb"
	"github.com/ideal-rucksack/datasource-golang-plugin/rdb/command"
	"github.com/ideal-rucksack/datasource-golang-plugins/rdb/mysql/client"
	"net/http"
)

var registry rdb.CommandFactoryRegistry

func init() {
	factoryRegistry := rdb.NewCommandFactoryRegistry()
	registry = *factoryRegistry
	registry.Register("databases", func(client rdb.Client) rdb.Command {
		return command.DatabaseCommand{Client: client}
	})
}

func main() {
	exec, err := rdb.Run()
	if err != nil {
		panic(err)
	}
	c := client.Client{}
	cmd, hasInstance := registry.GetCommand(exec.Action, c)
	if hasInstance {
		execute, err := cmd.Execute()
		if err != nil {
			panic(err)
		}
		err = cmd.Notify(*exec, rdb.NotifyRequest{Status: http.StatusOK, Payload: execute})
		if err != nil {
			panic(err)
		}
	}
}
