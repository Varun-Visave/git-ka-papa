package repo

import "strings"

type TreeNode struct {
	isFile   bool
	hash     string
	children map[string]*TreeNode
}

func buildTree(entries map[string]string) *TreeNode {
	root := &TreeNode{isFile: false, children: map[string]*TreeNode{}}

	for path, hash := range entries {
		parts := strings.Split(path, "/")
		current := root

		for _, part := range parts[:len(parts)-1]{
		if current.children[part] == nil{
			current.children[part] = &TreeNode{isFile: false, children: map[string]*TreeNode{}}
		}
		current = current.children[part]
	}
	fileName := parts[len(parts)-1]
	current.children[fileName] = &TreeNode{isFile: true, hash: hash}
	}

	return root
}
