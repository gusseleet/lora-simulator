package main

import (
	"lora_simulator/logging"
	"lora_simulator/lora"
	"lora_simulator/simulator"

	"flag"
	"time"
	"fmt"
)


func parseFlags() []string {

	c := flag.String("config", "remote.toml", "Path to configfile")
	d := flag.String("payload", "{'allInfo': 'test'}", "Json object of the payload")
	flag.Parse()

	return []string{*c, *d}

}

func main(){




	logging.Log().Info("Starting LORA Simulator")

	args := parseFlags();
	LConfig, err := lora.ReadConfig(lora.LoraConfig{}, args[0])

	if err != nil {
		logging.Log().Error("Error with config", err)
	}

	appSKey, nwsKey, devAddr := lora.GenerateAddr(&LConfig)
	g_config := simulator.GeneralConfig{DevAddr:devAddr, AppSKey:appSKey, NwkSKey:nwsKey, FPort: uint8(1)}

	tick := 0

	ticker := time.NewTicker(20 * time.Second)
	quit := make(chan struct{})
	go func(){
		for {
			select {
			case <-ticker.C:
				// do stuff
				err = simulator.SendMessageLora(&LConfig, &g_config, args[1], tick)

				if err != nil {
					logging.Log().Error("Error1", err)
				} else {
					logging.Log().Info("Message sent")
					tick = tick + 1
				}

			case <- quit:
				ticker.Stop()
				return

			}
		}

	}()



	for {
		a := 1
		a -= 1

	}
	fmt.Println()

}