package main

type ThroneInheritance struct {
	// king
	king string
	// child
	child map[string][]string
	// deadList
	deadList map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	this := ThroneInheritance{
		king:     kingName,
		child:    make(map[string][]string),
		deadList: make(map[string]bool),
	}
	return this
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.child[parentName] = append(this.child[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.deadList[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var dfs func(name string) []string
	dfs = func(name string) []string {
		res := make([]string, 0)

		if !this.deadList[name] {
			res = append(res, name)
		}

		for _, child := range this.child[name] {
			res = append(res, dfs(child)...)
		}
		return res
	}

	return dfs(this.king)
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
