# extract
func ExtractErrorFromJSON() extracts the contents in the "message" field of a JSON blob and returns it as a Go error value. If error messages are embedded in multiple layers, ExtractErrorFromJSON() returns the highest-level error.
