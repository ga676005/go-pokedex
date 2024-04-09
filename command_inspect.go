package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("please provide the name")
	}

	name := params[0]
	data, exist := cfg.store.Get(name)
	if !exist {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %v\n", data.Height)
	fmt.Printf("Weight: %v\n", data.Weight)
	fmt.Println("Stats:")
	for _, v := range data.Stats {
		fmt.Printf("  -%s: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range data.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}

	return nil
}
