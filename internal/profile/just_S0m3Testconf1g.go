package profile

import (
	"github.com/EXBO-Community/stalcraft-jvm-optimization/internal/config"
	"github.com/EXBO-Community/stalcraft-jvm-optimization/internal/sysinfo"
)

// Profile to test new featiure, in fact same as v1.1.2/default
func just_S0m3Testconf1g_generateDefault(sys sysinfo.Info) config.Config {
	return generateV112Default(sys)
}

// Same as default
func just_S0m3Testconf1g_generateLight(sys sysinfo.Info) config.Config {
	return generateV112Default(sys)
}
