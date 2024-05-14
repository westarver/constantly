package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/westarver/constantly/bridge"
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
	UserEntries().underlying = sel
	ent := widget.NewEntry()
	ent.TextStyle.Monospace = true
	t := sel.Text
	if len(t) == 0 {
		t = sel.PlaceHolder
	}
	typeWrap := title(t)
	ent.SetText(typeWrap)
	UserEntries().consType = ent
	ent.OnChanged = func(s string) {
		CellAt(Type, 0).SetText(s)
		CellAt(Type, 0).Entry().Refresh()
		SetDirty(true)
	}

	sel.OnChanged = func(s string) {
		SetDirty(true)
		UserEntries().consType.SetText(title(s))
	}

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
		SetDirty(true)
	})
	chk2 := widget.NewCheck("Generate method to map assoc data", func(b bool) {
		SetDirty(true)
	})
	chk3 := widget.NewCheck("Generate comment from assoc data", func(b bool) {
		SetDirty(true)
	})
	chk4 := widget.NewCheck("Generate method to return value", func(b bool) {
		SetDirty(true)
	})
	chk5 := widget.NewCheck("Generate Marshal/UnMarshal methods", func(b bool) {
		SetDirty(true)
	})

	UserEntries().genStr = chk1
	UserEntries().genAssoc = chk2
	UserEntries().genComment = chk3
	UserEntries().genValue = chk4
	UserEntries().genMarshal = chk5

	v := container.NewVBox(
		chk1,
		chk2,
		chk3,
		chk4,
		chk5,
		widget.NewButton("Generate Source File", func() { bridge.WriteConstants() }),
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
		c.Add(btn)

	case Suffix:
		btn := newContextMenuButton(label, suffixPopup())
		c.Add(btn)

	case Type:
		btn := newContextMenuButton(label, typePopup())
		btn.OnTapped = func() {
			typ := ConstType()
			if len(typ) == 0 {
				typ = UnderlyingType()
			}
			SetCellText(label, 0, typ)
		}
		c.Add(btn)

	case Value:
		btn := newContextMenuButton(Value, valuePopup())
		btn.OnTapped = func() {
			if CellAt(Value, 0).Text() == "" {
				SetCellText(label, 0, "iota")
			}
		}
		c.Add(btn)
		ApplicationData.val = btn

	case Assoc:
		btn := newContextMenuButton(AssocLongText, assocPopup())
		c.Add(btn)
		ApplicationData.assoc = btn
	}

	for i := 0; i < ApplicationData.rows; i++ {
		w := widget.NewEntry()
		if label == Type {
			if i > 0 {
				w.Disable()
			}
		}
		if label == BaseID {
			w.OnChanged = func(s string) {
				SetLastRow()
				ApplicationData.dirty = true
			}
		} else {
			w.OnChanged = func(s string) {
				ApplicationData.dirty = true
			}
		}

		CellAt(label, i).col = label
		CellAt(label, i).row = i
		//w.Text = fmt.Sprintf("%s  %d", label, i)
		CellAt(label, i).entry = w
		c.Add(w)
	}
	return c
}

func columnGrid() *fyne.Container {
	rightCol := lastColumn()
	last2 := container.New(
		layout.NewFormLayout(),
		newColumn(Value),
		newColumn(Assoc),
	)

	table := container.NewMax(container.NewHBox(
		firstColumn(ApplicationData.rows),
		container.New(
			layout.NewFormLayout(),
			newColumn(Prefix),
			newColumn(BaseID),
		),
		container.New(
			layout.NewFormLayout(),
			newColumn(Suffix),
			newColumn(Type),
		),
		last2,
		rightCol,
	))

	return container.NewMax(container.NewScroll(table))
}

func title(s string) string {
	tcaser := cases.Title(language.AmericanEnglish)
	return tcaser.String(s)
}
