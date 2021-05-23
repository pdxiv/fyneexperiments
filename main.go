package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var tableData = [][]string{
	{"2021-05-21 13:06:01", "fatal", "top right"},
	{"2021-05-22 13:19:11", "warning", "bottom right"},
	{"2021-05-22 13:19:31", "error", "bottom right"},
}

var textForTreeUID = map[string]string{
	"node_0":        "foo",
	"node_1":        "bar",
	"node_2":        "baz",
	"node_3":        "floop",
	"node_4":        "beep",
	"node_5":        "flarb",
	"subnode_0":     "subnode of foo",
	"subnode_1":     "subnode of bar",
	"subnode_2":     "subnode of baz",
	"subnode_3":     "subnode of floop",
	"subnode_4":     "subnode of beep",
	"subnode_5_XXX": "flarb - because you're worth it",
	"subnode_5_YYY": "flarb - oh my, this is tasty flarb",
	"subnode_5_ZZZ": "flarb - enough is never enough",
}

var treeUIDMapping = map[string][]string{
	"":              {"node_0", "node_1", "node_2", "node_3", "node_4", "node_5"},
	"node_0":        {"subnode_0"},
	"node_1":        {"subnode_1"},
	"node_2":        {"subnode_2"},
	"node_3":        {"subnode_3"},
	"node_4":        {"subnode_4"},
	"node_5":        {"subnode_5_XXX", "subnode_5_YYY"},
	"subnode_5_XXX": {"subnode_5_ZZZ"},
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	horizontalSplitter := makeSplittyBoiWithTreeAndTable()

	myWindow.SetContent(horizontalSplitter)
	myWindow.ShowAndRun()
}

func makeSplittyBoiWithTreeAndTable() *container.Split {
	left := makeTree()

	tableThing := makeTable()

	right := container.NewVSplit(
		widget.NewLabel("Top right side of splitter"),
		tableThing,
	)

	horizontalSplitter := container.NewHSplit(left, right)
	return horizontalSplitter
}

func makeTable() *widget.Table {
	tableThing := widget.NewTable(
		func() (int, int) {
			return len(tableData), len(tableData[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tableData[i.Row][i.Col])
		})
	return tableThing
}

func makeTree() *widget.Tree {
	childUIDs := func(uid widget.TreeNodeID) (c []widget.TreeNodeID) {
		return treeUIDMapping[uid]
	}

	createNode := func(branch bool) (o fyne.CanvasObject) {
		return widget.NewLabel("")
	}

	// It's a branch if uid exists, and has sub-values
	isBranch := func(uid widget.TreeNodeID) (ok bool) {
		if _, ok := treeUIDMapping[uid]; ok {
			if len(treeUIDMapping[uid]) > 0 {
				return true
			}
		}
		return false
	}

	updateNode := func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		node.(*widget.Label).SetText(textForTreeUID[uid])
	}

	return widget.NewTree(childUIDs, isBranch, createNode, updateNode)
}
