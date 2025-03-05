package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/IAmRasputin/artifacts-go/internal/client"
	"github.com/IAmRasputin/artifacts-go/internal/config"
)

var baseUrl string = "https://api.artifactsmmo.com/"

func ClientTest() error {
	genclient, err := client.NewClient(baseUrl)

	if err != nil {
		os.Exit(69)
	}

	ctx := context.Background()

	characterInfo, err := genclient.GetMyCharactersMyCharactersGet(ctx,
		func(ctx context.Context, req *http.Request) error {
			tokenGetter := config.NewDefaultTokenGetter()
			// Named such to avoid shadowing err above
			tok, tokenErr := tokenGetter.GetToken()

			if tokenErr != nil {
				return err
			}

			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tok))

			return nil
		})

	if err != nil {
		return err
	}

	bodyBytes, err := io.ReadAll(characterInfo.Body)
	if err != nil {
		return err
	}

	var myCharacters client.MyCharactersListSchema
	json.Unmarshal(bodyBytes, &myCharacters)

	formattedBytes, err := json.MarshalIndent(myCharacters, "", "    ")

	fmt.Println(string(formattedBytes))

	return nil
}
