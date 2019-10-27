package timetracker

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/google/uuid"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type response struct {
	values url.Values
	err    error
}

func NewConfig(clientID string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: "",
		RedirectURL:  "http://localhost:" + strconv.Itoa(Port),
		Scopes:       []string{calendar.CalendarScope},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}
}

func GetAccessToken(config *oauth2.Config) (*oauth2.Token, error) {
	state := uuid.New().String()

	challengeRaw := randomStringURLSafe(96)
	challengeSha256 := sha256.Sum256([]byte(challengeRaw))
	challengeUrlEncoded := base64url.Encode(challengeSha256[:])

	codeChallenge := oauth2.SetAuthURLParam("code_challenge", challengeUrlEncoded)
	codeChallengeMethod := oauth2.SetAuthURLParam("code_challenge_method", "S256")

	authURL := config.AuthCodeURL(state, oauth2.AccessTypeOffline, codeChallenge, codeChallengeMethod)

	log.Println("open the browser and start the authorization server")
	if err := browser.OpenURL(authURL); err != nil {
		return nil, fmt.Errorf("cannot open a browser to handle the authorization flow: %w", err)
	}

	res := <-callback("127.0.0.1:" + strconv.Itoa(Port))

	if errorCode := res.values.Get("error"); errorCode != "" {
		return nil, fmt.Errorf("the user did not grant the required permissions")
	}

	actual := res.values.Get("state")
	if state != actual {
		return nil, fmt.Errorf("state does not match the original one, something nasty happened")
	}

	code := res.values.Get("code")
	verifier := oauth2.SetAuthURLParam("code_verifier", challengeRaw)
	token, err := config.Exchange(context.Background(), code, verifier)
	if err != nil {
		return nil, fmt.Errorf("cannot exchange the OAuth 2 code for an access token: %w", err)
	}

	return token, nil
}

func callback(address string) chan *response {
	responseCh := make(chan *response)
	shutdownCh := make(chan bool)
	interruptCh := make(chan bool)

	server := &http.Server{Addr: address}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		var msg string
		if r.URL.Query().Get("code") != "" {
			msg = "Alfred TimeTracker authenticated correctly, you can now close this window."
		} else {
			msg = "Something went wrong with the authorization workflow. Please try again."
		}

		_, err := w.Write([]byte(msg))
		if err != nil {
			log.Printf("http.ResponseWriter write failed: %v", err)
		}

		interruptCh <- true
		responseCh <- &response{values: r.URL.Query()}
		shutdownCh <- true
	})

	// run the server
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("error running the authorization server: %s\n", err)
		}
	}()

	// shutdown the server after a timeout
	go func() {
		select {
		case <-interruptCh:
		case <-time.After(10 * time.Minute):
			responseCh <- &response{err: fmt.Errorf("timeout to complete the authorization flow expired")}
			shutdownCh <- false
		}
	}()

	// shutdown the server gracefully
	go func() {
		done := <-shutdownCh

		if done {
			log.Println("authorization flow done, shutting down the authorization server")
		} else {
			log.Println("timeout to done the authorization flow expired, shutting down the HTTP server")
		}

		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("authorization server could not shutdown gracefully: %v", err)
		}
	}()

	return responseCh
}
