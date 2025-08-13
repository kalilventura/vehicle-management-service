package infrastructure

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"
)

// toSnakeCase converts a given string from CamelCase to UPPER_SNAKE_CASE.
// For example: "DbHost" becomes "DB_HOST".
func toSnakeCase(str string) string {
	var result strings.Builder
	for i, r := range str {
		// If the character is an uppercase letter, and it's not the first character,
		// and the previous character was not an uppercase letter, add an underscore.
		if unicode.IsUpper(r) && i > 0 {
			// Check if the previous char was also uppercase to handle acronyms like "DB"
			prev := rune(str[i-1])
			if !unicode.IsUpper(prev) {
				result.WriteRune('_')
			}
		}
		result.WriteRune(unicode.ToUpper(r))
	}
	return result.String()
}

// SetEnvFromStruct iterates over the fields of the provided struct and sets
// an environment variable for each one. The environment variable key is the
// struct field name converted to UPPER_SNAKE_CASE.
func SetEnvFromStruct(s interface{}) error {
	// Use reflection to inspect the struct.
	v := reflect.ValueOf(s)

	// Ensure the input is a struct.
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("input is not a struct, but a %s", v.Kind())
	}

	// Get the type of the struct to access field names.
	t := v.Type()

	fmt.Println("Setting environment variables...")

	// Iterate over all the fields of the struct.
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// We expect all fields to be strings for this specific struct.
		// For a more generic function, you might need to handle other types.
		if value.Kind() != reflect.String {
			continue
		}

		// Convert the field name to snake case for the env var key.
		envKey := toSnakeCase(field.Name)
		envValue := value.String()

		// Set the environment variable.
		err := os.Setenv(envKey, envValue)
		if err != nil {
			// If setting the variable fails, return an error immediately.
			return fmt.Errorf("failed to set environment variable %s: %w", envKey, err)
		}
		fmt.Printf(" - Set %s=%s\n", envKey, envValue)
	}

	return nil
}

// UnsetEnvFromStruct iterates over the fields of the provided struct and unsets
// the corresponding environment variable for each one.
func UnsetEnvFromStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("input is not a struct, but a %s", v.Kind())
	}
	t := v.Type()

	fmt.Println("\nUnsetting environment variables...")

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		envKey := toSnakeCase(field.Name)

		// Unset the environment variable.
		if err := os.Unsetenv(envKey); err != nil {
			return fmt.Errorf("failed to unset environment variable %s: %w", envKey, err)
		}
		fmt.Printf(" - Unset %s\n", envKey)
	}

	return nil
}
