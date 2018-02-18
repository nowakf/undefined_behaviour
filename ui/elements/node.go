package elements

type Node struct {
	element  UiElement
	parent   *Node
	children []*Node
}

//root returns a Node without a parent, or associate element
func Root() *Node {
	n := new(Node)
	n.children = make([]*Node, 0)
	return n
}

func newNode(parent *Node, associate UiElement) *Node {
	n := new(Node)
	n.element = associate
	n.parent = parent
	n.parent.AddChild(n)
	n.children = make([]*Node, 0)
	return n
}

func (n *Node) GetParent() *Node {
	//need some error handling here
	if n.parent == nil {
		println("Node parent is nil")
		return nil
	}
	return n.parent
}

func (n *Node) GetChildren() []*Node {
	return n.children
}

func (n *Node) GetElement() UiElement {
	if n.element == nil {
		println("Node element is nil")
		return nil
	}
	return n.element
}

func (n *Node) AddChild(child *Node) {
	if child == n {
		panic("you tried to make a Node its own child")
	}
	n.children = append(n.children, child)
}
