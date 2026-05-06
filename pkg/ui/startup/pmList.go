package startup

import (
	"updep/pkg/config"
	packagemanager "updep/pkg/packageManager"

	tea "charm.land/bubbletea/v2"

	"charm.land/bubbles/v2/list"
)

type PmList struct {
	list.Model
	itemDelegate list.ItemDelegate
}

type PmListItem struct {
	pm packagemanager.PackageManager
}

func (i PmListItem) Title() string       { return i.pm.Name() }
func (i PmListItem) Description() string { return "" }
func (i PmListItem) FilterValue() string { return i.pm.Name() }

func NewPmList(title string) PmList {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.BorderForeground(
		config.Theme.Primary,
	)

	list := list.New(nil, delegate, 0, 0)
	list.Title = title
	list.SetShowStatusBar(false)
	list.Styles.Title = list.Styles.Title.Background(config.Theme.Primary).
		Foreground(config.Theme.PrimaryText)

	return PmList{
		Model:        list,
		itemDelegate: delegate,
	}
}

func (l *PmList) SetItems(pms []packagemanager.PackageManager) tea.Cmd {
	items := make([]list.Item, len(pms))
	for i, pm := range pms {
		items[i] = PmListItem{pm: pm}
	}

	cmd := l.Model.SetItems(items)
	l.SetHeight(l.minListHeight())
	return cmd
}

func (l *PmList) SetSize(width, maxHeight int) {
	l.Model.SetSize(width, l.minListHeight())
}

func (l PmList) minListHeight() int {
	itemHeight := l.itemDelegate.Height()
	separatorHeight := l.itemDelegate.Spacing()
	baseListHeight := 6

	return len(
		l.Items(),
	)*(itemHeight+separatorHeight) - separatorHeight + baseListHeight
}
