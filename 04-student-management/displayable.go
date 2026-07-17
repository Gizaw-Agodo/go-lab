package main 

type Displayable interface {
	display()
}

type Updatable interface {
	updateName(name string) error 
}

func DisplayInformation(item Displayable) {
	item.display()
}

func ChangeName(item Updatable, newName string) error {
	return item.updateName(newName)
}