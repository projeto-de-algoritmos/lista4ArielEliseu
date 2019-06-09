package control

type control_I interface {
	checkstatus() (bool, map[string]string)
}
