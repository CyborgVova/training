package command

import "container/list"

type Command interface {
	Execute() string
}

type ToggleOnCommand struct {
	receiver *Reciever
}

func (on *ToggleOnCommand) Execute() string {
	return on.receiver.ToggleOn()
}

type ToggleOffCommand struct {
	receiver *Reciever
}

func (off *ToggleOffCommand) Execute() string {
	return off.receiver.ToggleOff()
}

type Reciever struct{}

func (r *Reciever) ToggleOn() string {
	return "Toggle On"
}

func (r *Reciever) ToggleOff() string {
	return "Toggle Off"
}

type Invoke struct {
	commands list.List
}

func (i *Invoke) AddCommandToInvoke(command Command) {
	i.commands.PushBack(command)
}

func (i *Invoke) RemoveCommandFromFront() {
	if i.commands.Len() > 0 {
		i.commands.Remove(i.commands.Front())
	}
}

func (i *Invoke) RemoveCommandFromBack() {
	if i.commands.Len() > 0 {
		i.commands.Remove(i.commands.Back())
	}
}

func (i *Invoke) Executor() string {
	result := ""
	com := i.commands.Front().Value.(Command)
	for j := 0; j < i.commands.Len(); j++ {
		result += com.Execute() + "\n"
		com = i.commands.Front().Next().Value.(Command)
	}
	return result
}
