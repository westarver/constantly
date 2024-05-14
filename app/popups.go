package app

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/mgutz/str"
	"github.com/westarver/fynewidgets"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type contextMenuButton struct {
	widget.Button
	menu *fyne.Menu
}

func (b *contextMenuButton) TappedSecondary(e *fyne.PointEvent) {
	widget.ShowPopUpMenuAtPosition(b.menu, fyne.CurrentApp().Driver().CanvasForObject(b), e.AbsolutePosition)
}

func newContextMenuButton(label string, menu *fyne.Menu) *contextMenuButton {
	b := &contextMenuButton{menu: menu}
	b.Text = label

	b.ExtendBaseWidget(b)
	return b
}

func prefixPopup() *fyne.Menu {
	return fyne.NewMenu(Prefix, menuItemClear(Prefix), menuItemCopy(Prefix), menuItemXiota(Prefix), menuItemXalpha(Prefix), menuItemXassoc(Prefix), menuItemCatIota(Prefix), menuItemCatAlpha(Prefix), menuItemCatAssoc(Prefix), menuItemCaseXform(Prefix))
}

func baseIDPopup() *fyne.Menu {
	return fyne.NewMenu(BaseID, menuItemClear(BaseID), menuItemCopy(BaseID), menuItemCopyToLimit(BaseID), menuItemXiota(BaseID), menuItemXalpha(BaseID), menuItemXassoc(BaseID), menuItemCatIota(BaseID), menuItemCatAlpha(BaseID), menuItemCatAssoc(BaseID), menuItemCaseXform(BaseID))
}

func suffixPopup() *fyne.Menu {
	return fyne.NewMenu(Suffix, menuItemClear(Suffix), menuItemCopy(Suffix), menuItemXiota(Suffix), menuItemXalpha(Suffix), menuItemXassoc(Suffix), menuItemCatIota(Suffix), menuItemCatAlpha(Suffix), menuItemCatAssoc(Suffix), menuItemCaseXform(Suffix))
}

func typePopup() *fyne.Menu {
	return fyne.NewMenu(Type, menuItemClear(Type))
}

func valuePopup() *fyne.Menu {
	return fyne.NewMenu(Value, menuItemClear(Value), menuItemCopy(Value), menuItemXiota(Value), menuItemXalpha(Value), menuItemXassoc(Value), menuItemCatIota(Value), menuItemCatAlpha(Value), menuItemCatAssoc(Value), menuItemToggleValueSize(), menuItemCaseXform(Value))
}

func assocPopup() *fyne.Menu {
	return fyne.NewMenu(Assoc, menuItemClear(Assoc), menuItemCopy(Assoc), menuItemXiota(Assoc), menuItemXalpha(Assoc), menuItemCatIota(Assoc), menuItemCatAlpha(Assoc), menuItemCaseXform(Assoc))
}

func menuItemClear(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Clear", func() {
		r := LastRow() + 1
		for i := 0; i < r; i++ {
			SetCellText(label, i, "")
		}
	})
}

func menuItemCopy(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Copy down", func() {
		r := LastRow() + 1
		for i := 0; i < r; i++ {
			SetCellText(label, i, CellText(label, 0))
		}
	})
}

func menuItemCopyToLimit(label string) *fyne.MenuItem {
	num := fynewidgets.NewNumericalEntry()

	fi := widget.NewFormItem("Copy down to:", num)
	var fis = []*widget.FormItem{fi}

	return fyne.NewMenuItem("Copy to limit", func() {
		var limit int
		num.OnChanged = func(n string) {
			limit, _ = strconv.Atoi(n)
			if limit > ApplicationData.rows || limit < 1 {
				limit = ApplicationData.rows
			}
		}

		dlg := dialog.NewForm("Enter limit", "Ok", "Cancel", fis,
			func(b bool) {
				if b {
					for i := 0; i < limit; i++ {
						SetCellText(label, i, CellText(label, 0))
					}
					SetLastRow()
				}
			},
			MainWindow())
		dlg.Show()
	})
}

func menuItemCaseXform(label string) *fyne.MenuItem {
	mi := fyne.NewMenuItem("Change case", nil)
	mi.ChildMenu = fyne.NewMenu(
		"Case transforms",
		fyne.NewMenuItem("To upper",
			func() {
				r := LastRow() + 1
				for i := 0; i < r; i++ {
					SetCellText(label, i, strings.ToUpper(CellText(label, i)))
					fmt.Println(CellText(label, i))
				}
			}),
		fyne.NewMenuItem("To lower",
			func() {
				r := LastRow() + 1
				for i := 0; i < r; i++ {
					SetCellText(label, i, strings.ToLower(CellText(label, i)))
				}
			}),
		fyne.NewMenuItem("To title",
			func() {
				r := LastRow() + 1
				for i := 0; i < r; i++ {
					tcaser := cases.Title(language.AmericanEnglish)
					SetCellText(label, i, tcaser.String(CellText(label, i)))
				}
			}),
		fyne.NewMenuItem("To snake",
			func() {
				r := LastRow() + 1
				for i := 0; i < r; i++ {
					SetCellText(label, i, str.Underscore(str.Camelize(CellText(label, i))))
				}
			}),
		fyne.NewMenuItem("To camel",
			func() {
				r := LastRow() + 1
				for i := 0; i < r; i++ {
					SetCellText(label, i, str.Camelize(CellText(label, i)))
				}
			}),
	)

	return mi
}

func menuItemXiota(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Copy text down + iota", func() {
		start := CellText(label, 0)
		r := LastRow() + 1
		for i := 0; i < r; i++ {
			SetCellText(label, i, start+fmt.Sprintf("%d", i))
		}
	})
}

func menuItemCatIota(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Concatenate text + iota", func() {
		r := LastRow() + 1
		for i := 0; i < r; i++ {
			SetCellText(label, i, CellText(label, i)+fmt.Sprintf("%d", i))
		}
	})
}

func menuItemXalpha(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Copy text down + alpha++", func() {
		start := CellText(label, 0)
		r := LastRow() + 1
		for i := 0; i < r; i++ {
			SetCellText(label, i, start+string(rune('A'+i)))
		}
	})
}

func menuItemCatAlpha(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Concatenate text + alpha++", func() {
		r := LastRow() + 1
		for i := 0; i < r; i++ {
			SetCellText(label, i, CellText(label, i)+string(rune('A'+i)))
		}
	})
}

func menuItemXassoc(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Copy text + assoc data", func() {
		start := CellText(label, 0)
		r := LastRow() + 1
		for i := 0; i < r; i++ {
			CellAt(label, i).SetText(start + CellText(Assoc, i))
		}
	})
}

func menuItemCatAssoc(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Concatenate text + assoc data", func() {
		r := LastRow() + 1
		for i := 0; i < r; i++ {
			CellAt(label, i).SetText(CellText(label, i) + CellText(Assoc, i))
		}
	})
}

func menuItemToggleValueSize() *fyne.MenuItem {
	return fyne.NewMenuItem("Toggle size", func() {
		ApplicationData.valueshort = !ApplicationData.valueshort
		if ApplicationData.valueshort {
			ApplicationData.val.Text = ValueShortText
			ApplicationData.assoc.Text = AssocLongText
		} else {
			ApplicationData.val.Text = ValueLongText
			ApplicationData.assoc.Text = AssocShortText
		}
		ApplicationData.val.Refresh()
		ApplicationData.assoc.Refresh()
	})
}
