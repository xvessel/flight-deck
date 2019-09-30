package component

var CMD_CREATE string = "CREATE"
var CMD_READY string = "READY"
var CMD_UPDATE_CHECK string = "UPDATE_CHECK"
var CMD_UPDATE string = "UPDATE"
var CMD_DELETE string = "DELETE"

type ComponentIf interface {
	Run(cmdstr string, env []string, namespace string, id string) (error, map[string]string)
	GetSpec() Spec
}
