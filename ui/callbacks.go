package ui

import (
	"fmt"
	"indexfuzzysearch/config"
	"indexfuzzysearch/search"
	"strconv"
	"strings"

	"github.com/gen2brain/iup-go/iup"
)

func exit(ih iup.Ihandle) int {
	config.Save()
	return iup.CLOSE
}

func searchName(ih iup.Ihandle) int {
	val := iup.GetHandle("searchField").GetAttribute("VALUE")

	fmt.Println("Search:", ui2utf8(val))

	output := search.FindMarriage(strings.TrimSpace((ui2utf8(val))))

	iup.GetHandle("output").SetAttribute("VALUE", utf82ui(output))

	return iup.DEFAULT
}

func toogleChurches(ih iup.Ihandle) int {
	config.Config.Churches[ih.GetAttribute("key")] = ih.GetAttribute("VALUE") == "ON"
	config.Save()
	return iup.DEFAULT
}

func toogleInstantSearch(ih iup.Ihandle) int {
	config.Config.InstantSearch = ih.GetAttribute("VALUE") == "ON"
	config.Save()
	return iup.DEFAULT
}

func selectAll(ih iup.Ihandle) int {
	for _, box := range boxes {
		box.SetAttribute("VALUE", "ON")
		toogleChurches(box)
	}
	return iup.DEFAULT
}

func selectNone(ih iup.Ihandle) int {
	for _, box := range boxes {
		box.SetAttribute("VALUE", "OFF")
		toogleChurches(box)
	}
	return iup.DEFAULT
}

func valueChanged(ih iup.Ihandle) int {
	title := ih.GetAttribute("TITLE")
	t := ih.GetAttribute("ORIENTATION")
	switch string(t[0]) {
	case "H":
		value := ih.GetAttribute("VALUE")
		floatValue, _ := strconv.ParseFloat(value, 32)
		iup.GetHandle(title).SetAttribute("TITLE", int(floatValue))
		config.Config.Year[title] = int(floatValue)
	}
	config.Save()
	return iup.DEFAULT
}

func searchInstant(ih iup.Ihandle) int {
	if !config.Config.InstantSearch {
		return iup.DEFAULT
	}
	return searchName(ih)
}
