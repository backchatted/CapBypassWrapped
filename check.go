package capbypass

import "errors"

var (
	SupportCaptchaTask = []string{
		"FunCaptchaTask",
		"FunCaptchaProxylessTask",

		"ReCaptchaV2Task",
	}
)

func checkTask(params CapBypassPayload) error {
	captchaTask := params.Task.Type
	exists := false
	for _, task := range SupportCaptchaTask {
		if task == captchaTask {
			exists = true
		}
	}
	if !exists {
		return errors.New("unSupported task " + captchaTask + "you need to pay attention to capitalization, current support types fellow\\n")
	}

	return nil
}
