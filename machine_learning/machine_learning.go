package machine_learning

//The learning game - Machine Learning #1
//Growing up you would have learnt a lot of things like not to stand in fire, to drink food and eat water and not to jump off very tall things But Machines have it difficult they cannot learn for themselves we have to tell them what to do, why don't we give them a chance to learn it for themselves?
//
//Task
//Your task is to finish the Machine object. What the machine object must do is learn from its mistakes! The Machine will be given a command and a number you will return a random action. After the command has returned you will be given a response (true/false) if the response is true then you have done good, if the response is false then the action was a bad one. You must program the machine to learn to apply an action to a given command using the reponse given. Note: It must take no more than 20 times to teach an action to a command also different commands can have the same action.
//
//Info
//In the preloaded section there is a constant called ACTIONS it is a function that returns the 5 possible actions.
//In Java, this a constant Actions.FUNCTIONS of type List<Function<Integer, Integer>>.
//In C++, the actions can be accessed by get_action(i)(unsigned int num) where i chooses the function (and therefore can range from 0 to 4) and num is its argument.
//In python ACTIONS() returns a list of lambdas.
//In Golang Actions() retruns a function slice []func(int) int




//type Machine struct {
//	Cmd int
//	Mapping map[int]int
//}
//
////Function called to get your machine initialised
//func NewMachine() Machine {
//	return Machine{Mapping:make(map[int]int)}
//}
//
//func (m *Machine) Command(cmd int, num int) int {
//	m.Cmd = cmd
//	return Actions()[m.Mapping[cmd]](num)
//}
//
//func (m *Machine) Response(res bool) {
//	if !res {
//		if m.Mapping[m.Cmd] == 4 {
//			m.Mapping[m.Cmd] = 0
//		} else {
//			m.Mapping[m.Cmd] += 1
//		}
//	}
//}



type Machine struct {
	cmd map[int]int
	lrn map[int]bool
	action int
	fnc []func(int) int
}

//Function called to get your machine initialised
func NewMachine() Machine {
	m := Machine{}
	m.cmd = make(map[int]int)
	m.lrn = make(map[int]bool)
	m.fnc = Actions()
	return m
}

func (m *Machine) Command(cmd int, num int) int {

	m.action = cmd
	_, ok := m.lrn[cmd]

	if !ok {
		m.lrn[cmd] = false
		m.cmd[cmd] = 0
	} else if !m.lrn[cmd] {
		m.cmd[cmd]++
		if m.cmd[cmd] == len(m.fnc) { m.cmd[cmd] = 0 } // possible infinite loop if response always fails
	}

	i := m.cmd[cmd]
	return m.fnc[i](num)
}

func (m *Machine) Response(res bool) {
	if res {
		m.lrn[m.action] = true
	} else if m.lrn[m.action] {
		m.lrn[m.action] = false
	}
}
