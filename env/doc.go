// Package env provides functionality for loading environment variables.
//
// The variables are primarily sourced from optional .env files located
// within the project directory. These .env files offer a convenient method
// for setting environment variables for both development and production.
//
// The package features an AutoLoad function. AutoLoad is designed to
// automatically load environment variables from a default .env file.
//
// Additionally, AutoLoad checks for the existence of .env.<ENV>.local and
// .env.local files, loading variables from them if they are found. This
// functionality allows for environment-specific variables to be set,
// thereby enhancing the flexibility of the environment configuration.
package env
