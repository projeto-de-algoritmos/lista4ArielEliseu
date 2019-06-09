package control

type control struct {
	status bool
	log    map[string]bool
	issues map[string]string
}
