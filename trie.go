package cuten

import "fmt"

type node struct {
	pattern  string
	part     string  // 路由的某部分，例如/user/admin中的 /admin
	children []*node //  子节点
	precise  bool    // 是否精确匹配
}

// 这里应该只有一个child，如果1个以上，会产生路由分歧
func (n *node) matchChild(part string) *node {
	for _, n2 := range n.children {
		if n2.part == part {
			return n2
		}
	}
	return nil
}

// 匹配所以合法的子路由
func (n *node) matchChildren(part string) []*node {
	ret := make([]*node, 0, len(n.children))
	for _, n2 := range n.children {
		if n2.part == part || !n2.precise {
			ret = append(ret, n2)
		}
	}
	return ret
}

/*
	/user/:id
	/user/*
	这种路由会冲突
	/user/*
	/user/cliend/:d
	/user/cliend/*
	这种路由也会冲突
	*路由 层级及其以后不能有路由层级
*/

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		return
	}
	// 查看child里面是否有*
	for _, v := range n.children {
		if v.part[0] == '*' {
			panic(fmt.Sprintf("pattern:%s, part:%s already register !", v.pattern, v.part))
		}
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{pattern: pattern, part: part, precise: !(part[0] == ':' || part[0] == '*')}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height {
		if n.pattern != "" || !n.precise {
			return n
		}
		return nil
	}
	if n.part[0] == '*' {
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		ret := child.search(parts, height+1)
		if ret != nil {
			return ret
		}
	}
	return nil
}
