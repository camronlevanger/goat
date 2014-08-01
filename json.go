package goat

import "net/http"

// WriteError writes a string as JSON encoded error
//
// Example output:
// {
//   "error": "User not Found"
// }
func WriteError(w http.ResponseWriter, e string) {
	WriteJSON(w, map[string]string{
		"error": e,
	})
}

// WriteJSON writes the given interface as JSON or returns an error
func WriteJSON(w http.ReponseWriter, v interface{}) err {
	b, err := MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	w.Write(b)
	return nil
}