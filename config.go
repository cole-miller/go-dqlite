package dqlite

// ConfigMultiThread sets the threading mode of SQLite to Multi-thread.
//
// By default go-dqlite configures SQLite to Single-thread mode, because the
// dqlite engine itself is single-threaded, and enabling Multi-thread or
// Serialized modes would incur in a performance penality.
//
// If your Go process also uses SQLite directly (e.g. using the
// github.com/mattn/go-sqlite3 bindings) you might need to switch to
// Multi-thread mode in order to be thread-safe.
//
// IMPORTANT: It's possible to successfully change SQLite's threading mode only
// if no SQLite APIs have been invoked yet (e.g. no database has been opened
// yet). Therefore you'll typically want to call ConfigMultiThread() very early
// in your process setup. Alternatively you can set the GO_DQLITE_MULTITHREAD
// environment variable to 1 at process startup, in order to prevent go-dqlite
// from setting Single-thread mode at all.
func ConfigMultiThread() error {
	return nil
}

func init() {
}
