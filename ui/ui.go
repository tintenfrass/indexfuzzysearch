package ui

import (
	"indexfuzzysearch/config"
	"indexfuzzysearch/search"

	"github.com/gen2brain/iup-go/iup"
)

var boxes map[string]iup.Ihandle

func BuildAndRun() {
	//Control
	iup.Open()
	defer iup.Close()

	//Config
	minSlider := iup.Val("HORIZONTAL").SetAttribute("SIZE", "200x15").SetAttribute("TITLE", "min").SetAttribute("MIN", config.YearMin).SetAttribute("MAX", config.YearMax).SetAttribute("VALUE", config.Config.Year["min"])
	minLabel := iup.Label("").SetHandle("min").SetAttribute("TITLE", config.Config.Year["min"])
	maxSlider := iup.Val("HORIZONTAL").SetAttribute("SIZE", "200x15").SetAttribute("TITLE", "max").SetAttribute("MIN", config.YearMin).SetAttribute("MAX", config.YearMax).SetAttribute("VALUE", config.Config.Year["max"])
	maxLabel := iup.Label("").SetHandle("max").SetAttribute("TITLE", config.Config.Year["max"])
	instantSearch := iup.Toggle("Suchen beim Tippen").SetAttribute("VALUE", getOnOffInstantSearch("instantSearch")).SetAttribute("key", "instantSearch")

	//Buttons
	exitButton := iup.Button("Exit").SetAttribute("SIZE", "50x15")
	searchButton := iup.Button("Search").SetAttribute("SIZE", "50x15")
	selectAllButton := iup.Button("Select All").SetAttribute("SIZE", "45x10")
	selectNoneButton := iup.Button("Select None").SetAttribute("SIZE", "45x10")

	//Search and Find
	searchField := iup.Text().SetAttribute("SIZE", "150x").SetHandle("searchField")
	results := iup.MultiLine().SetAttribute("SIZE", "400x420").SetAttribute("READONLY", "YES").SetHandle("output")

	configs := iup.GridBox(
		iup.Label("Min:").SetAttribute("SIZE", "30x15"),
		minSlider,
		minLabel,
		iup.Label("Max:").SetAttribute("SIZE", "30x15"),
		maxSlider,
		maxLabel,
		instantSearch.SetAttribute("SIZE", "x15"),
	).SetAttribute("NUMDIV", 3)
	//Map
	boxes = map[string]iup.Ihandle{
		"Lommatzsch":          iup.Toggle("Lommatzsch"),
		"Zehren":              iup.Toggle("Zehren"),
		"Leuben":              iup.Toggle("Leuben"),
		"Meißen St. Afra":     iup.Toggle(utf82ui("Meißen St. Afra")),
		"Planitz":             iup.Toggle("Planitz"),
		"Meißen Frauenkirche": iup.Toggle(utf82ui("Meißen Frauenkirche")),
		"Ziegenhain":          iup.Toggle("Ziegenhain"),
		"Krögis":              iup.Toggle(utf82ui("Krögis")),
		"Naustadt":            iup.Toggle("Naustadt"),
		"Constappel":          iup.Toggle("Constappel"),
		"Miltitz":             iup.Toggle("Miltitz"),
		"Heynitz":             iup.Toggle("Heynitz"),
		"Taubenheim":          iup.Toggle("Taubenheim"),
		"Röhrsdorf":           iup.Toggle(utf82ui("Röhrsdorf")),
		"Weistropp":           iup.Toggle("Weistropp"),
		"Burkhardswalde":      iup.Toggle("Burkhardswalde"),
		"Tanneberg":           iup.Toggle("Tanneberg"),
		"Limbach":             iup.Toggle("Limbach"),
		"Wilsdruff":           iup.Toggle("Wilsdruff"),
		"Unkersdorf":          iup.Toggle("Unkersdorf"),
		"Briesnitz":           iup.Toggle("Briesnitz"),
		"Neukirchen":          iup.Toggle("Neukirchen"),
		"Blankenstein":        iup.Toggle("Blankenstein"),
		"Grumbach":            iup.Toggle("Grumbach"),
		"Kesselsdorf":         iup.Toggle("Kesselsdorf"),
		"Dresden Annenkirche": iup.Toggle("Dresden Annenkirche"),
		"Mohorn":              iup.Toggle("Mohorn"),
		"Herzogswalde":        iup.Toggle("Herzogswalde"),
		"Pesterwitz":          iup.Toggle("Pesterwitz"),
		"Fördergersdorf":      iup.Toggle(utf82ui("Fördergersdorf")),
		"Tharandt":            iup.Toggle("Tharandt"),
		"Döhlen":              iup.Toggle(utf82ui("Döhlen")),
		"Plauen":              iup.Toggle("Plauen"),
		"Somsdorf":            iup.Toggle("Somsdorf"),
	}

	for key, box := range boxes {
		box.SetAttribute("VALUE", getOnOffChurches(key)).SetAttribute("key", key)
		box.SetAttribute("TITLE", utf82ui(key)+"\r\n("+search.GetMinMax(key)+")")
	}

	grid := iup.GridBox(
		boxes["Lommatzsch"],
		boxes["Zehren"].SetAttributes("SIZE=70x0"),
		iup.Space().SetAttributes("SIZE=75x0"),
		iup.Space().SetAttributes("SIZE=80x0"),
		iup.Space().SetAttributes("SIZE=65x0"),
		iup.Space().SetAttributes("SIZE=65x0"),
		iup.Space().SetAttributes("SIZE=90x0"),
		boxes["Leuben"],
		iup.Space(),
		boxes["Meißen St. Afra"],
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		boxes["Planitz"],
		iup.Space(),
		boxes["Meißen Frauenkirche"],
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		boxes["Ziegenhain"],
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space().SetAttributes("SIZE=0x15"),
		boxes["Krögis"],
		iup.Space(),
		boxes["Naustadt"],
		boxes["Constappel"],
		iup.Space(),
		iup.Space(),
		iup.Space().SetAttributes("SIZE=0x15"),
		boxes["Miltitz"],
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space().SetAttributes("SIZE=0x15"),
		boxes["Heynitz"],
		boxes["Taubenheim"],
		boxes["Röhrsdorf"],
		boxes["Weistropp"],
		iup.Space(),
		iup.Space(),
		iup.Space().SetAttributes("SIZE=0x15"),
		boxes["Burkhardswalde"],
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space().SetAttributes("SIZE=0x15"),
		boxes["Tanneberg"],
		boxes["Limbach"],
		boxes["Wilsdruff"],
		boxes["Unkersdorf"],
		boxes["Briesnitz"],
		iup.Space(),
		iup.Space().SetAttributes("SIZE=0x15"),
		boxes["Neukirchen"],
		boxes["Blankenstein"],
		boxes["Grumbach"],
		boxes["Kesselsdorf"],
		iup.Space(),
		boxes["Dresden Annenkirche"],
		iup.Space().SetAttributes("SIZE=0x15"),
		iup.Space(),
		boxes["Mohorn"],
		boxes["Herzogswalde"],
		iup.Space(),
		boxes["Pesterwitz"],
		iup.Space(),
		iup.Space().SetAttributes("SIZE=0x15"),
		iup.Space(),
		iup.Space(),
		boxes["Fördergersdorf"],
		boxes["Tharandt"],
		boxes["Döhlen"],
		boxes["Plauen"],
		iup.Space().SetAttributes("SIZE=0x15"),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		boxes["Somsdorf"],
		iup.Space(),
		iup.Space(),
		selectAllButton,
		selectNoneButton,
	).SetAttribute("NUMDIV", 7)

	configFrame := iup.Frame(configs).SetAttribute("TITLE", "Einstellungen").SetAttributes("SIZE=500x70")
	mapFrame := iup.Frame(grid).SetAttribute("TITLE", "Kirchengemeinden")

	//Infotext
	infotext := iup.Label("").SetAttribute("TITLE", utf82ui("Bsp:\n"+
		"vorname nachname  ->  Suche nach Vor- und Nachnamen\r\n"+
		"name ?  ->  Suche nach Vornamen\r\n"+
		"? name  ->  Suche nach Nachnamen\r\n"+
		"name  ->  Suche unabhängig von Vor- oder Nachname\r\n"+
		"\r\n"+
		"Es wird Groß- und kleinschreibung unterschieden.\r\n",
	))

	content := iup.Hbox(
		iup.Vbox(
			configFrame,
			mapFrame,
			searchField,
			searchButton,
			infotext,
			exitButton,
		).SetAttributes("MARGIN=10x10, GAP=8"),
		iup.Vbox(
			results,
		),
	)

	dlg := iup.Dialog(content).SetAttributes(`TITLE="Fuzzy Search"`)
	dlg.SetHandle("dlg").SetAttributes("SIZE=950x450")

	//Callbacks
	iup.SetCallback(exitButton, "ACTION", iup.ActionFunc(exit))
	iup.SetCallback(searchButton, "ACTION", iup.ActionFunc(searchName))
	iup.SetCallback(searchField, "VALUECHANGED_CB", iup.ValueChangedFunc(searchInstant))
	iup.SetCallback(selectAllButton, "ACTION", iup.ActionFunc(selectAll))
	iup.SetCallback(selectNoneButton, "ACTION", iup.ActionFunc(selectNone))
	for _, box := range boxes {
		iup.SetCallback(box, "ACTION", iup.ActionFunc(toogleChurches))
	}
	iup.SetCallback(minSlider, "VALUECHANGED_CB", iup.ValueChangedFunc(valueChanged))
	iup.SetCallback(maxSlider, "VALUECHANGED_CB", iup.ValueChangedFunc(valueChanged))
	iup.SetCallback(instantSearch, "ACTION", iup.ActionFunc(toogleInstantSearch))

	//Run
	iup.Map(dlg)
	iup.Show(dlg)
	iup.MainLoop()
}

func getOnOffChurches(key string) string {
	if config.Config.Churches[key] {
		return "ON"
	}
	return "OFF"
}

func getOnOffInstantSearch(key string) string {
	if config.Config.InstantSearch {
		return "ON"
	}
	return "OFF"
}
