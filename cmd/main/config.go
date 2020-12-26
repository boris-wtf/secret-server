package main

import "github.com/spf13/viper"

func initConfig() error {
	if err := viper.BindEnv("max_note_length"); err != nil {
		return err
	}
	viper.SetDefault("max_note_length", 4096)
	return nil
}
