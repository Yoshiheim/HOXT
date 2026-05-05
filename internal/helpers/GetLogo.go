package helpers

import (
	"fmt"
	"hoxt/data"
	"time"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func UpdateLogo() {
	c, err := data.LoadDynamicConfig("./data/textconf.json")
	if err != nil {
		fmt.Println(err)
	}
	if c.Logo.Hide == true {
		return
	}
	flags := aic_package.DefaultFlags()

	flags.Dimensions = []int{c.Logo.Heigth, c.Logo.Width}
	if c.Logo.CharMap == "" {
		flags.CustomMap = " .:-~+=*%&)@"
	} else {
		flags.CustomMap = c.Logo.CharMap
	}
	flags.Negative = true
	flags.Full = false
	flags.Dither = true
	asciiArt, err := aic_package.Convert(c.Logo.Path, flags)
	if err != nil {
		fmt.Printf("-- %s --\n", err)
		data.Logo = []byte("")
	} else {
		fmt.Printf("[CONVERT IMAGE TO ASCIII - OK]\n")
		data.Logo = []byte(asciiArt)
	}
}

func UpdateLogoTick() {
	UpdateLogo()
	Dest, err := ParseCustomDuration(data.Configs.UpdateLogoTick)
	if err != nil {
		return
	}
	go func() {
		tick := time.NewTicker(Dest)
		for range tick.C {
			UpdateLogo()
			fmt.Println("[LOGO UPDATED]")
		}
	}()
}
