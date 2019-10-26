package heap

func (class *Class) IsArray() bool {
	return class.Name[0] == '['
}

func (class *Class) IsPrimitiveArray() bool {
	return class.IsArray() && len(class.Name) == 2
}

func (class *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(class.Name)
	return class.bootLoader.LoadClass(componentClassName)
}

func (class *Class) arrayClass() *Class {
	arrayClassName := getArrayClassName(class.Name)
	return class.bootLoader.LoadClass(arrayClassName)
}
