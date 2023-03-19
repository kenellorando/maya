// commands.go
// Maya commands and handlers

package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "hello-maya",
		Description: "Test command, says hello to Maya.",
	},
	{
		Name:        "get-instances",
		Description: "Get all instances managable by Maya.",
	},
	{
		Name:        "start-instance",
		Description: "Start a specified instance.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "instance-id",
				Description: "Instance ID",
				Required:    true,
			},
		},
	},
	{
		Name:        "stop-instance",
		Description: "Stop a specified instance.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "instance-id",
				Description: "Instance ID",
				Required:    true,
			},
		},
	},
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"hello-maya": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hello.",
			},
		})
	},
	"get-instances": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		describeInstancesOutput, err := describeInstances()
		if err != nil {
			fmt.Println(err)
			return
		}
		instanceList := describeInstancesOutput.Reservations
		var results string
		results += "```"
		for _, j := range instanceList {
			for _, k := range j.Instances[0].Tags {
				if *k.Key == "Name" {
					results += *k.Value
				}
			}
			results += "\t"
			results += *j.Instances[0].InstanceId
			results += "\t"
			results += *j.Instances[0].State.Name
			results += "\n"
		}
		results += "```"
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: results,
			},
		})
	},
	"start-instance": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		_, err := startInstance(optionMap["instance-id"].StringValue())
		if err != nil {
			fmt.Println("error calling startInstance")
			return
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Starting instance...",
			},
		})
	},
	"stop-instance": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		_, err := stopInstance(optionMap["instance-id"].StringValue())
		if err != nil {
			fmt.Println("error calling stopInstance")
			return
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Stopping instance...",
			},
		})
	},
}
