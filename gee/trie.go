package gee

import "strings"

type node struct {
	pattern  string  // 路由全名，例如/p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node //子节点
	isWild   bool    // 是否精准匹配
}

// 获取第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) (child *node) {

	for _, v := range n.children {
		if part == v.part || v.isWild {
			child = v
			break
		}
	}

	return
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) (nodes []*node) {

	for _, v := range n.children {
		if part == v.part || v.isWild {
			nodes = append(nodes, v)
		}
	}

	return
}

func (n *node) insert(pattern string, parts []string, height int) {

	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) (result *node) {

	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern != "" {
			result = n
		}
	} else {
		part := parts[height]
		children := n.matchChildren(part)

		for _, child := range children {
			result = child.search(parts, height+1)
			if result != nil {
				break
			}
		}
	}

	return
}
