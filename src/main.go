package main

import (
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"google.golang.org/api/iterator"
	"bufio"
	"os"
	"fmt"
	"log"
	"context"
	"mr4dd/mdparse"
)

func main() {
	API_KEY := getAPIKey()

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-flash")
	
	
	model.SafetySettings = []*genai.SafetySetting{
	  {
	    Category:  genai.HarmCategoryHarassment,
	    Threshold: genai.HarmBlockNone,
	  },
	  {
	    Category:  genai.HarmCategoryHateSpeech,
	    Threshold: genai.HarmBlockNone,
	  },
	}


	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
		    Parts: []genai.Part{
		      genai.Text("you're a virtual assistant helping users from their terminal."),
		    },
		    Role: "user",
		  },
		  &genai.Content{
		    Parts: []genai.Part{
		      genai.Text("Hello, how may I assiste you today"),
		    },
		    Role: "model",
		  },
	}
	loop(cs, ctx)
}

func getAPIKey() (string) {
	homeDir := os.Getenv("HOME")
	configDir := homeDir + "/.config/jolt"
	configFile := configDir + "/config"
	
	file, err := os.Open(configFile)
	if err != nil {
		file, err := os.Create(configFile)
		if err != nil {
			err := os.Mkdir(configDir, 0755)
			if err != nil{
				panic(err)
			}
		}
		defer file.Close()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	key := scanner.Text()
	return key
}

func loop(cs *genai.ChatSession, ctx context.Context) {

	for {	
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Prompt: ")
		prompt, _ := reader.ReadString('\n')

		iter := cs.SendMessageStream(ctx, genai.Text(prompt))
		for {
			resp, err := iter.Next()
			if err == iterator.Done {
				fmt.Print("\n")
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			parseAndPrint(resp)
		}
	}
}

func parseAndPrint(resp *genai.GenerateContentResponse){
	for _, cand := range resp.Candidates{
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				stringPart := toString(part)
				mdParsed := mdparse.Parse(stringPart)
				fmt.Print(mdParsed)
				os.Stdout.Sync()
			}	
		}
	}


}

func toString(part genai.Part) (string){
	
	stringC := fmt.Sprint(part)
	return stringC
}


