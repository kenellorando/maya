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
		Description: "Says hello to Maya.",
	},
	{
		Name:        "describe-instances",
		Description: "Get a list of all instances manageable by Maya.",
	},
	{
		Name:        "describe-instance-status",
		Description: "Get an instance's reachability and system health status.",
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
				Content: "Hello, I'm Maya. 👋",
			},
		})
	},
	"describe-instances": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		describeInstancesOutput, err := describeInstances()
		if err != nil {
			fmt.Println(err)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Could not describe instances.",
				},
			})
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
	"describe-instance-status": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		describeInstanceStatusOutput, err := describeInstanceStatus(optionMap["instance-id"].StringValue())
		if err != nil {
			fmt.Println("error calling describeInstanceStatus")
			return
		}
		if len(describeInstanceStatusOutput.InstanceStatuses) != 1 {
			fmt.Println("No instance found matching the provided ID")
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "No status available for `" + optionMap["instance-id"].StringValue() + "`.",
				},
			})
			return
		}
		instanceStatus := describeInstanceStatusOutput.InstanceStatuses[0]
		var results string

		results += "Status for `" + *instanceStatus.InstanceId + "`:"
		results += "```"
		results += "State: " + *instanceStatus.InstanceState.Name
		results += "\n"
		results += "Instance Reachability: " + *instanceStatus.InstanceStatus.Status
		results += "\n"
		if instanceStatus.InstanceStatus.Details[0].ImpairedSince != nil {
			results += "\tImpaired: " + instanceStatus.InstanceStatus.Details[0].ImpairedSince.GoString()
			results += "\n"
		}
		results += "System Health: " + *instanceStatus.SystemStatus.Status
		if instanceStatus.SystemStatus.Details[0].ImpairedSince != nil {
			results += "\n"
			results += "\tImpaired: " + instanceStatus.SystemStatus.Details[0].ImpairedSince.GoString()
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
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Could not start `" + optionMap["instance-id"].StringValue() + "`.",
				},
			})
			return
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Starting instance `" + optionMap["instance-id"].StringValue() + "`...",
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
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Could not stop `" + optionMap["instance-id"].StringValue() + "`.",
				},
			})
			return
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Stopping instance `" + optionMap["instance-id"].StringValue() + "`...",
			},
		})
	},
}
