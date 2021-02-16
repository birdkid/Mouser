package actions

import (
	"os"
	"strings"

	"github.com/go-vgo/robotgo"
)

// NewAppBranch creates an app-based actions branch.
func NewAppBranch(
	branches map[string]Action,
	fallback Action,
) Action {
	return NewAppBranchCustom(
		branches,
		fallback,
		getAppRobotgo,
		os.PathSeparator,
	)
}

// NewAppBranchCustom creates an app-based actions branch based on a custom
// app-detection callback.
func NewAppBranchCustom(
	branches map[string]Action,
	fallback Action,
	getApp func() string,
	pathSeparator rune,
) Action {
	return func() {
		app := getApp()
		action := getAppMatch(branches, fallback, app, pathSeparator)
		if action != nil {
			action()
		}
	}
}

func getAppMatch(
	branches map[string]Action,
	fallback Action,
	app string,
	pathSeparator rune,
) Action {
	if action, ok := branches[app]; ok {
		return action
	}

	if app != "" {
		pathSep := string(pathSeparator)
		for appPrefix := range branches {
			if appPrefix != "" && strings.HasPrefix(app, appPrefix+pathSep) {
				// Remember action assocaition for next time.
				action := branches[appPrefix]
				branches[app] = action
				return action
			}
		}
	}

	return fallback
}

func getAppRobotgo() string {
	pid := robotgo.GetPID()
	if pid == 0 {
		return ""
	}
	appPath, err := robotgo.FindPath(pid)
	if err != nil {
		return ""
	}
	return appPath
}
