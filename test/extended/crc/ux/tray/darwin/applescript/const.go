package applescript

const (
	scriptsRelativePath string = "darwin/applescript/scripts"

	checkTrayIconIsVisible string = "checkTrayIconIsVisible.applescript"
	clickTrayMenuItem      string = "clickTrayMenuItem.applescript"
	setPullSecretFile      string = "setPullSecretFile.applescript"
	getTrayFieldlValue     string = "getTrayFieldlValue.applescript"
	installTray            string = "installTray.applescript"
	getOCLoginCommand      string = "getOCLoginCommand.applescript"
	runOCLoginCommand      string = "runOCLoginCommand.applescript"

	bundleIdentifier string = "com.redhat.codeready.containers"
	appPath          string = "/Applications/CodeReady Containers.app"
)

// TODO check if move to tray package
const (
	actionStart  string = "start"
	actionStop   string = "stop"
	actionDelete string = "delete"
	actionQuit   string = "quit"

	fieldState string = "state"

	stateRunning string = "Running"
	stateStopped string = "Stopped"

	userKubeadmin string = "kubeadmin"
	userDeveloper string = "developer"
)

var (
	elements = []Element{
		{
			Name:         actionStart,
			AXIdentifier: "start"},
		{
			Name:         actionStop,
			AXIdentifier: "stop"},
		{
			Name:         actionDelete,
			AXIdentifier: "delete"},
		{
			Name:         actionQuit,
			AXIdentifier: "quit"},
		{
			Name:         fieldState,
			AXIdentifier: "cluster_status"},
		{
			Name:         userKubeadmin,
			AXIdentifier: "kubeadmin_login"},
		{
			Name:         userDeveloper,
			AXIdentifier: "developer_login"},
	}
)

const (
	uxCheckAccessibilityDuration = "2s"
	uxCheckAccessibilityRetry    = 10
)
