package cmd

import (
	"flag"
	"fmt"
	"net/http"

	"archipelago.dev/chat/chat"
	"archipelago.dev/chat/config"
	"github.com/rs/cors"
)

func ServeCmd() *flag.FlagSet {
	c := flag.NewFlagSet("serve", flag.ExitOnError)
	return c
}

func ServeHandler(c *flag.FlagSet) error {
	handler := chat.NewHandler()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./web/dist")))
	mux.HandleFunc("/ws", handler.HandleWebSocket)

	fmt.Printf("Serving on http://localhost%s\n", config.Port)
	err := http.ListenAndServe(config.Port, cors.Default().Handler(mux))
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}

	return nil
}
