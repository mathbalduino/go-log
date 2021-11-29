// Matheus Leonel Balduino
// Everywhere, under @mathbalduino
//   @mathbalduino on GitHub
//   @mathbalduino on Instagram
//   @mathbalduino on Twitter
// Live at mathbalduino.com.br
// 2021-11-29 4:39 PM

package main

type treeNode struct {
	log
	parentPtr *treeNode
	childs    []treeNode
}

func buildTree(logs []log) []treeNode {
	parents := make(map[string]*treeNode, len(logs))
	rootNodes := make([]treeNode, 0, len(logs))
	for _, log := range logs {
		timestamp := log.Timestamp
		parent := parents[log.Parent]
		if parent == nil {
			rootNodes = append(rootNodes, treeNode{log, nil, nil})
			parents[timestamp] = &rootNodes[len(rootNodes)-1]
			continue
		}

		parent.childs = append(parent.childs, treeNode{log, parent, nil})
		parents[timestamp] = &parent.childs[len(parent.childs)-1]
	}
	return rootNodes
}
