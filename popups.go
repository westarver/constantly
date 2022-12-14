package constantly

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/mgutz/str"
	fynewidgets "github.com/westarver/fyne-widgets"
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
	return fyne.NewMenu(Value, menuItemClear(Value), menuItemCopy(Value), menuItemXiota(Value), menuItemXalpha(Value), menuItemXassoc(Value), menuItemCatIota(Value), menuItemCatAlpha(Value), menuItemCatAssoc(Value), menuItemToggleValueSize(Value), menuItemCaseXform(Value))
}

func assocPopup() *fyne.Menu {
	return fyne.NewMenu(Assoc, menuItemClear(Assoc), menuItemCopy(Assoc), menuItemXiota(Assoc), menuItemXalpha(Assoc), menuItemCatIota(Assoc), menuItemCatAlpha(Assoc), menuItemCaseXform(Assoc))
}

func menuItemClear(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Clear", func() {
		rows := AppData.rows
		for i := 0; i < rows; i++ {
			AppData.userEntries.columns[label].entries[i].SetText("")
		}
	})
}

func menuItemCopy(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Copy down", func() {
		for i := 0; i < AppData.rows; i++ {
			AppData.userEntries.columns[label].entries[i].SetText(AppData.userEntries.columns[label].entries[0].Text)
		}
	})
}

func menuItemCopyToLimit(label string) *fyne.MenuItem {
	num := fynewidgets.NewNumericalEntry()
	fi := widget.NewFormItem("Copy down to:", num)
	var fis = []*widget.FormItem{fi}
	return fyne.NewMenuItem(
		"Copy to limit",
		func() {
			var limit int
			num.OnChanged = func(n string) { limit, _ = strconv.Atoi(n) }
			dlg := dialog.NewForm("Enter limit", "Ok", "Cancel", fis,
				func(b bool) {
					if b {
						for i := 0; i < limit; i++ {
							AppData.userEntries.columns[label].entries[i].SetText(AppData.userEntries.columns[label].entries[0].Text)
						}
					}
				},
				AppData.MainWindow())
			dlg.Show()
		})
}

func menuItemCaseXform(label string) *fyne.MenuItem {
	mi := fyne.NewMenuItem("Change case", nil)
	mi.ChildMenu = fyne.NewMenu(
		"Case transforms",
		fyne.NewMenuItem("To upper",
			func() {
				for i := 0; i < AppData.rows; i++ {
					AppData.userEntries.columns[label].entries[i].SetText(strings.ToUpper(AppData.userEntries.columns[label].entries[i].Text))
				}
			}),
		fyne.NewMenuItem("To lower",
			func() {
				for i := 0; i < AppData.rows; i++ {
					AppData.userEntries.columns[label].entries[i].SetText(strings.ToLower(AppData.userEntries.columns[label].entries[i].Text))
				}
			}),
		fyne.NewMenuItem("To title",
			func() {
				for i := 0; i < AppData.rows; i++ {
					tcaser := cases.Title(language.AmericanEnglish)
					AppData.userEntries.columns[label].entries[i].SetText(tcaser.String(AppData.userEntries.columns[label].entries[i].Text))
				}
			}),
		fyne.NewMenuItem("To snake",
			func() {
				for i := 0; i < AppData.rows; i++ {
					AppData.userEntries.columns[label].entries[i].SetText(str.Underscore(str.Camelize(AppData.userEntries.columns[label].entries[i].Text)))
				}
			}),
		fyne.NewMenuItem("To camel",
			func() {
				for i := 0; i < AppData.rows; i++ {
					AppData.userEntries.columns[label].entries[i].SetText(str.Camelize(AppData.userEntries.columns[label].entries[i].Text))
				}
			}),
	)

	return mi
}

func menuItemXiota(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Copy text down + iota", func() {
		start := AppData.userEntries.columns[label].entries[0].Text
		for i := 0; i < AppData.rows; i++ {
			AppData.userEntries.columns[label].entries[i].SetText(start + fmt.Sprintf("%d", i))
		}
	})
}

func menuItemCatIota(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Concatenate text + iota", func() {
		for i := 0; i < AppData.rows; i++ {
			AppData.userEntries.columns[label].entries[i].SetText(AppData.userEntries.columns[label].entries[i].Text + fmt.Sprintf("%d", i))
		}
	})
}

func menuItemXalpha(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Copy text down + alpha++", func() {
		start := AppData.userEntries.columns[label].entries[0].Text
		for i := 0; i < AppData.rows; i++ {
			AppData.userEntries.columns[label].entries[i].SetText(start + string(rune('A'+i)))
		}
	})
}

func menuItemCatAlpha(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Concatenate text + alpha++", func() {
		for i := 0; i < AppData.rows; i++ {
			AppData.userEntries.columns[label].entries[i].SetText(AppData.userEntries.columns[label].entries[i].Text + string(rune('A'+i)))
		}
	})
}

func menuItemXassoc(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Copy text + assoc data", func() {
		start := AppData.userEntries.columns[label].entries[0].Text
		for i := 0; i < AppData.rows; i++ {
			AppData.userEntries.columns[label].entries[i].SetText(start + AppData.userEntries.columns[Assoc].entries[i].Text)
		}
	})
}

func menuItemCatAssoc(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Concatenate text + assoc data", func() {
		for i := 0; i < AppData.rows; i++ {
			AppData.userEntries.columns[label].entries[i].SetText(AppData.userEntries.columns[label].entries[i].Text + AppData.userEntries.columns[Assoc].entries[i].Text)
		}
	})
}

func menuItemToggleValueSize(label string) *fyne.MenuItem {
	return fyne.NewMenuItem("Toggle size", func() {
		AppData.userEntries.valueshort = !AppData.userEntries.valueshort
		if AppData.userEntries.valueshort {
			AppData.userEntries.valueBtn.Text = ValueShortText
			AppData.userEntries.assocBtn.Text = AssocLongText
		} else {
			AppData.userEntries.valueBtn.Text = ValueLongText
			AppData.userEntries.assocBtn.Text = AssocShortText
		}
		AppData.val.Refresh()
		AppData.assoc.Refresh()
		//fyne.CurrentApp().Driver().CanvasForObject(AppData.userEntries.valueBtn).Content().Refresh()
	})
}
