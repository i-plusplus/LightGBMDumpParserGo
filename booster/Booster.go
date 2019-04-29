package booster

type Booster struct {
	features [][]string
	value float64
	isLeaf bool
	index int
	left *Booster
	right *Booster
	condition *Condition
}

func (b *Booster) SetFeatures(f [][]string){
	b.features = f
}

func (b *Booster) SetValue(f float64){
	b.value = f
}

func (b *Booster) SetIsLeaf(f bool){
	b.isLeaf = f
}

func (b *Booster) SetIndex(f int){
	b.index = f
}

func (b *Booster) SetLeft(f *Booster){
	b.left = f
}

func (b *Booster) SetRight(f *Booster){
	b.right = f
}

func (b *Booster) SetCondition(f Condition){
	b.condition = &f
}





func (b *Booster) GetLeaf(input *map[string]string) int{
	if(b.isLeaf){
		return b.index
	}
	if d, err := (*b.condition).IsLeft(input); err==nil && d == true{
		return b.left.GetLeaf(input)
	}
	return b.right.GetLeaf(input);
}

func (b Booster) GetValue(input *map[string]string) (float64, error){
	if(b.isLeaf){
		return b.value, nil
	}
	r, err := (*b.condition).IsLeft(input)
	if err != nil{
		return -1, err
	}
	if r {
		return b.left.GetValue(input)
	}
	return b.right.GetValue(input);
}

