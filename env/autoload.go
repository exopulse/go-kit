package env

import "os"

// AutoLoad loads the default .env file and the .env.local file if it exists.
// The default .env file is expected to be embedded in the binary, and provided
// as the defaultEnvContent argument.
//
// It also loads the .env.<ENV>.local file if the ENV environment variable is set.
// If the ENV environment variable is not set, it defaults to "dev".
func AutoLoad(defaultEnvContent string) error {
	return autoLoad(
		defaultEnvContent,
		resolveSelector(os.Getenv("ENV")),
		NewLoader(),
	)
}

func autoLoad(defaultEnvContent, selector string, loader *Loader) error {
	// embedded
	if err := loader.Apply(defaultEnvContent); err != nil {
		return err
	}

	files := []string{
		".env." + selector + ".local",
		".env.local",
	}

	for _, file := range files {
		if err := loader.LoadOptional(file); err != nil {
			return err
		}
	}

	return nil
}

func resolveSelector(selector string) string {
	const defaultSelector = "dev"

	if selector == "" {
		return defaultSelector
	}

	return selector
}
