package gemini

import (
	"encoding/json"
	"fmt"
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
	"github.com/google/generative-ai-go/genai"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"os"
	"strconv"
	"strings"
)

type Gemini struct {
	model *genai.GenerativeModel
}

func New() *Gemini {
	key := os.Getenv("GEMINI_API_KEY")
	typeModel := os.Getenv("GEMINI_TYPE_MODEL")
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(key))
	if err != nil {
		panic(err)
	}

	model := client.GenerativeModel(typeModel)

	return &Gemini{model: model}
}

func (g *Gemini) PredictFoodNutritions(ctx context.Context, picture []byte) (dto.GenerateNutritionsResponse, error) {
	content, err := g.model.GenerateContent(ctx,
		genai.Text("I am providing an image of a food item. Please analyze the image and provide a JSON response containing the following details: the name of the food, the number of calories it contains, the amount of carbohydrates (in grams), the amount of protein (in grams), whether the food is recommended for pregnant women/children, and a very short description of the food with indonesia language. if its not food just set food_name to \\\"notfood\\\""),
		genai.Text("Format the JSON response exactly like this example: {\"food_name\": \"Example Food Name\", \"calories\": 123, \"carbohydrates\": 45, \"protein\": 67, \"recommended_for_pregnant_children\": true, \"description\": \"A short description of the food\"}"),
		genai.Text("Give me only the JSON output in one-line, without anything else"),
		genai.ImageData("jpeg", picture),
	)
	if err != nil {
		return dto.GenerateNutritionsResponse{}, err
	}

	part := content.Candidates[0].Content.Parts[0]
	byteJson, err := json.Marshal(part)
	if err != nil {
		return dto.GenerateNutritionsResponse{}, err
	}

	strJson, err := strconv.Unquote(string(byteJson))
	if err != nil {
		return dto.GenerateNutritionsResponse{}, err
	}

	var res dto.GenerateNutritionsResponse
	if err := jsoniter.Unmarshal([]byte(strJson), &res); err != nil {
		return dto.GenerateNutritionsResponse{}, err
	}

	return res, nil
}

func (g *Gemini) GenerateChatResponse(ctx context.Context, message string) (string, error) {
	content, err := g.model.GenerateContent(ctx,
		genai.Text("You are a health assistant whose aim is to educate users of applications aimed at pregnant women and mothers with toddlers. Answer messages from users directly without asking the user to answer back. Here the user will send you a message or question. Your job is to answer the question clearly. Next, this is a message from the user."),
		genai.Text(message),
		genai.Text(".Answer it in indonesian language"),
	)
	if err != nil {
		return "", err
	}

	var formatted strings.Builder
	if content != nil && content.Candidates != nil {
		for _, cand := range content.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					formatted.WriteString(fmt.Sprintf("%v", part))
				}
			}
		}
	}

	return formatted.String(), err
}
