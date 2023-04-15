package cmd

import (
	"fmt"
	"time"

	"github.com/shlima/fortune/internal/pkg/bruteforce"
	"github.com/shlima/fortune/internal/pkg/key"
	"github.com/urfave/cli/v2"
)

func BruteforceTest(c *cli.Context) error {
	address := c.Args().First()
	if address == "" {
		return fmt.Errorf("please provide address in an argument")
	}

	force := bruteforce.New(
		NewIndex(c),
		NewKeyGen(c).SetTesting(address),
		c.Int(FlagWorkers.Name),
	)

	return terror(c, force)
}

func Bruteforce(c *cli.Context) error {
	force := bruteforce.New(
		NewIndex(c),
		NewKeyGen(c),
		c.Int(FlagWorkers.Name),
	)

	return terror(c, force)
}

func terror(c *cli.Context, force *bruteforce.Executor) error {
	fmt.Printf("loaded: %d addresses\n", force.DataLength())
	fmt.Printf("test passed: %v\n", force.Get(c.String(FlagTestAddress.Name)))

	go heartbit(c, force)
	go telegram(c, force)

	force.Run(onFound(c))
	return nil
}

func heartbit(c *cli.Context, force *bruteforce.Executor) {
	for range time.Tick(time.Second * time.Duration(c.Int(FlagHeartBeatSec.Name))) {
		fmt.Println(force.Heartbeat().ToString())
	}
}

func telegram(c *cli.Context, force *bruteforce.Executor) {
	bot := NewTelegram(c)
	for range time.Tick(time.Second * time.Duration(c.Int(FlagTelegramPingSec.Name))) {
		if err := bot.SendHeartBeat(force.Heartbeat()); err != nil {
			fmt.Printf("failed to send to telegram: %s\n", err)
		}
	}
}

func onFound(c *cli.Context) bruteforce.FoundFn {
	return func(chain key.Chain) {
		fmt.Printf(">>>> FOUND <<<< %s\n", chain.ToString())
		err := NewTelegram(c).KeyFound(chain)
		fmt.Printf("Send to telegram result: %s\n", err)
		panic(chain.ToString())
	}
}