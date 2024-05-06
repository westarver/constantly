package app

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
	Prefix         = "Prefix"
	BaseID         = "Base Identifier"
	Suffix         = "Suffix"
	Type           = "      Type       "
	SliderValue    = 0.62
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
	applicationData.userEntries.underlying = sel
	ent := widget.NewEntry()
	ent.TextStyle.Monospace = true
	ent.SetPlaceHolder(sel.Text)
	ent.OnSubmitted = func(s string) {
		applicationData.userEntries.columns[Type].entries[0].SetText(s)
		applicationData.dirty = true
	}

	sel.OnSubmitted = func(s string) {
		applicationData.dirty = true
		applicationData.userEntries.underlying = sel

		// The default typename that wraps the go type
		// selected from the dropdown is <type>+"Wrap"
		// Edit to suit
		typeWrap := sel.Text + "Wrap"
		ent.SetText(typeWrap)
	}

	applicationData.userEntries.consType = ent
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

	chk1 := widget.NewCheck("Generate method to return string", func(b bool) {
		applicationData.userEntries.genStr.SetChecked(b)
		applicationData.dirty = true
	})
	chk2 := widget.NewCheck("Generate method to map assoc data", func(b bool) {
		applicationData.userEntries.genAssoc.SetChecked(b)
		applicationData.dirty = true
	})
	chk3 := widget.NewCheck("Generate comment from assoc data", func(b bool) {
		applicationData.userEntries.genComment.SetChecked(b)
		applicationData.dirty = true
	})
	chk4 := widget.NewCheck("Generate method to return value", func(b bool) {
		applicationData.userEntries.genValue.SetChecked(b)
		applicationData.dirty = true
	})
	chk5 := widget.NewCheck("Generate Marshal/UnMarshal methods", func(b bool) {
		applicationData.userEntries.genMarshal.SetChecked(b)
		applicationData.dirty = true
	})

	applicationData.userEntries.genStr = chk1
	applicationData.userEntries.genAssoc = chk2
	applicationData.userEntries.genComment = chk3
	applicationData.userEntries.genValue = chk4
	applicationData.userEntries.genMarshal = chk5

	v := container.NewVBox(
		chk1,
		chk2,
		chk3,
		chk4,
		chk5,
		widget.NewButton("Generate Source File", func() { writeConstants() }),
	)
	return container.NewVBox(h, v)
}

func newColumn(label string) *fyne.Container {
	c := container.New(layout.NewVBoxLayout())

	switch label {
	case Prefix:
		btn := newContextMenuButton(label, prefixPopup())
		c.Add(btn)

	case BaseID:
		btn := newContextMenuButton(label, baseIDPopup())
		btn.OnTapped = func() {
			applicationData.rows = populatedCells()
		}
		c.Add(btn)

	case Suffix:
		btn := newContextMenuButton(label, suffixPopup())
		c.Add(btn)

	case Type:
		btn := newContextMenuButton(label, typePopup())
		btn.OnTapped = func() {
			cons := applicationData.userEntries.consType.Text
			if len(cons) == 0 {
				cons = applicationData.userEntries.consType.PlaceHolder
			}
			applicationData.userEntries.columns[label].entries[0].SetText(cons)
		}
		c.Add(btn)

	case Value:
		btn := newContextMenuButton(Value, valuePopup())
		btn.OnTapped = func() {
			if applicationData.userEntries.columns[Value].entries[0].Text == "" {
				applicationData.userEntries.columns[Value].entries[0].SetText("iota")
			}
		}
		c.Add(btn)
		applicationData.userEntries.valueBtn = btn

	case Assoc:
		btn := newContextMenuButton(AssocLongText, assocPopup())
		c.Add(btn)
		applicationData.userEntries.assocBtn = btn
	}

	for i := 0; i < defaultRows; i++ {
		w := widget.NewEntry()
		if label == Type {
			if i > 0 {
				w.Disable()
			}
		}
		if label == BaseID {
			w.OnSubmitted = func(s string) {
				for j := 0; j < defaultRows; i++ {
					if j > applicationData.lastRow {
						applicationData.lastRow = j
						fmt.Println(applicationData.lastRow)
						break
					}
				}
				applicationData.dirty = true
			}
		} else {
			w.OnSubmitted = func(s string) { applicationData.dirty = true }
		}
		c.Add(w)
		applicationData.userEntries.columns[label].entries[i] = w
	}

	return c
}

func columnGrid() *fyne.Container {
	rightCol := lastColumn()
	last2 := container.New(
		layout.NewFormLayout(),
		newColumn(Value), // applicationData.rows),
		newColumn(Assoc), // applicationData.rows),
	)
	applicationData.val = last2.Objects[0].(*fyne.Container).Objects[0].(*contextMenuButton)
	applicationData.assoc = last2.Objects[1].(*fyne.Container).Objects[0].(*contextMenuButton)

	table := container.NewMax(container.NewHBox(
		firstColumn(defaultRows),
		container.New(
			layout.NewFormLayout(),
			newColumn(Prefix), // applicationData.rows),
			newColumn(BaseID), // applicationData.rows),
		),
		container.New(
			layout.NewFormLayout(),
			newColumn(Suffix), // applicationData.rows),
			newColumn(Type),   // applicationData.rows),
		),
		last2,
		rightCol,
	))

	return container.NewMax(container.NewScroll(table))
}
