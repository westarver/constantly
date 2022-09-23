package constantly

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/widget"
)

const (
	AssocShortText = "As.."
	AssocLongText  = "            Associated Data             "
	ValueShortText = "Value"
	ValueLongText  = "                 Value                  "
	Value          = "Value"
	Assoc          = "Assoc"
	SliderValue    = 0.70
	Prefix         = "Prefix"
	BaseID         = "Base Identifier"
	Suffix         = "Suffix"
	Type           = "      Type       "
)

func firstColumn(rows int) *fyne.Container {
	c := container.NewVBox(
		widget.NewLabel("#"),
	)

	for i := 1; i <= rows; i++ {
		c.Add(widget.NewLabel(fmt.Sprintf("%d", i)))
	}

	return c
}

func lastColumn() *fyne.Container {
	sel := widget.NewSelectEntry([]string{"int", "uint", "byte", "int16", "uint16", "int32", "uint32", "int64", "uint64", "float32", "float64", "string"})
	sel.SetPlaceHolder("int")
	sel.OnChanged = func(s string) { AppData.dirty = true }
	AppData.userEntries.underlying = sel
	ent := widget.NewEntry()
	ent.TextStyle.Monospace = true
	ent.SetPlaceHolder("int")
	ent.OnChanged = func(s string) {
		AppData.userEntries.columns[Type].entries[0].SetText(s)
		AppData.dirty = true
	}
	AppData.userEntries.consType = ent
	lbl := widget.NewLabelWithStyle("type", fyne.TextAlignLeading, fyne.TextStyle{Bold: false, Monospace: true})
	split := container.NewHSplit(
		ent,
		sel,
	)
	split.SetOffset(SliderValue)
	h := container.New(layout.NewFormLayout(),
		lbl,
		split,
	)

	chk1 := widget.NewCheck("Generate func to return string", func(b bool) {
		AppData.userEntries.genStr.SetChecked(b)
		AppData.dirty = true
	})
	chk2 := widget.NewCheck("Generate func to map assoc data", func(b bool) {
		AppData.userEntries.genAssoc.SetChecked(b)
		AppData.dirty = true
	})
	chk3 := widget.NewCheck("Generate comment from assoc data", func(b bool) {
		AppData.userEntries.genComment.SetChecked(b)
		AppData.dirty = true
	})
	AppData.userEntries.genStr = chk1
	AppData.userEntries.genAssoc = chk2
	AppData.userEntries.genComment = chk3

	v := container.NewVBox(
		chk1,
		chk2,
		chk3,
		widget.NewButton("Generate Source File", func() { writeConstants() }),
	)
	return container.NewVBox(h, v)
}

func newColumn(label string, rows int) *fyne.Container {
	c := container.New(layout.NewVBoxLayout())

	switch label {
	case Prefix:
		btn := newContextMenuButton(label, prefixPopup())
		c.Add(btn)

	case BaseID:
		btn := newContextMenuButton(label, baseIDPopup())
		btn.OnTapped = func() {
			AppData.rows = populatedCells()
		}
		c.Add(btn)

	case Suffix:
		btn := newContextMenuButton(label, suffixPopup())
		c.Add(btn)

	case Type:
		btn := newContextMenuButton(label, typePopup())
		btn.OnTapped = func() {
			cons := AppData.userEntries.consType.Text
			if len(cons) == 0 {
				cons = AppData.userEntries.consType.PlaceHolder
			}
			AppData.userEntries.columns[label].entries[0].SetText(cons)
		}
		c.Add(btn)

	case Value:
		btn := newContextMenuButton(Value, valuePopup())
		btn.OnTapped = func() {
			if AppData.userEntries.columns[Value].entries[0].Text == "" {
				AppData.userEntries.columns[Value].entries[0].SetText("iota")
			}
		}
		c.Add(btn)
		AppData.userEntries.valueBtn = btn

	case Assoc:
		btn := newContextMenuButton(AssocLongText, assocPopup())
		c.Add(btn)
		AppData.userEntries.assocBtn = btn
	}

	for i := 0; i < AppData.rows; i++ {
		w := widget.NewEntry()
		if label == Type {
			if i > 0 {
				w.Disable()
			}
		}
		if label == BaseID {
			w.OnChanged = func(s string) {
				AppData.rows = populatedCells()
				AppData.dirty = true
			}
		} else {
			w.OnChanged = func(s string) { AppData.dirty = true }
		}
		c.Add(w)
		AppData.userEntries.columns[label].entries[i] = w
	}

	return c
}

func columnGrid() *fyne.Container {
	rightCol := lastColumn()
	last2 := container.New(
		layout.NewFormLayout(),
		newColumn(Value, AppData.rows),
		newColumn(Assoc, AppData.rows),
	)
	AppData.val = last2.Objects[0].(*fyne.Container).Objects[0].(*contextMenuButton)
	AppData.assoc = last2.Objects[1].(*fyne.Container).Objects[0].(*contextMenuButton)

	table := container.NewMax(container.NewHBox(
		firstColumn(AppData.rows),
		container.New(
			layout.NewFormLayout(),
			newColumn(Prefix, AppData.rows),
			newColumn(BaseID, AppData.rows),
		),
		container.New(
			layout.NewFormLayout(),
			newColumn(Suffix, AppData.rows),
			newColumn(Type, AppData.rows),
		),
		last2,
		rightCol,
	))

	return container.NewMax(container.NewScroll(table))
}
