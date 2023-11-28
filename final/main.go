// package main

// import (
// 	"fmt"
// 	// "bufio"
// 	// "context"
// 	// "errors"
// 	// "fmt"
// 	// "io"
// 	"log"
// 	"os"
// 	// "strings"

// 	gpt3 "github.com/sashabaranov/go-openai"
// 	"github.com/spf13/cobra"
// )

// // var generateCmd = &cobra.Command{
// // 	Use:   "generate",
// // 	Short: "Generate text using GPT-3",
// // 	Long:  "Generate text using GPT-3 with the given prompt",
// // 	Args:  cobra.MinimumNArgs(1),
// // 	Run: func(cmd *cobra.Command, args []string) {
// // 	// 取得用戶輸入的提示
// // 	prompt := args[0]
	
// // 	// 使用 GPT-3 模型生成文本
// // 	generatedText, err := client.Generate(prompt)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	// 輸出生成的文本
// // 	fmt.Println(generatedText)
// // },

// // }

// func main() {
// 	// log.SetOutput(new(NullWriter))
// 	apiKey := os.Getenv("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")
// 	client := gpt3.NewClient(apiKey)

// 	// 創建一個新的 Cobra 命令行應用程序
// 	var cmd = &cobra.Command{
// 		Use:   "gpt3-console-client",
// 		Short: "A simple GPT-3 console client",
// 		Long:  "A simple GPT-3 console client built with Cobra and Go-gpt3",
// 	}
// 	var generateCmd = &cobra.Command{
// 		Use:   "generate",
// 		Short: "Generate text using GPT-3",
// 		Long:  "Generate text using GPT-3 with the given prompt",
// 		Args:  cobra.MinimumNArgs(1),
// 		Run: func(cmd *cobra.Command, args []string) {
// 		// 取得用戶輸入的提示
// 		prompt := args[0]
		
// 		// 使用 GPT-3 模型生成文本
// 		generatedText, err := client.Generate(prompt)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
	
// 		// 輸出生成的文本
// 		fmt.Println(generatedText)
// 	},
	
// 	}
// 	// 定義子命令
// 	cmd.AddCommand(generateCmd)

// 	// 執行命令
// 	cmd.Execute()
// }

// package main

// import (
// 	"bufio"
// 	"context"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"

// 	gpt3 "github.com/sashabaranov/go-openai"
// 	"github.com/spf13/cobra"
// )

// func GetResponse(client *gpt3.Client, ctx context.Context, question string) {
// 	req := gpt3.CompletionRequest{
// 		Model:     gpt3.GPT3TextDavinci001,
// 		MaxTokens: 300,
// 		Prompt:    question,
// 		Stream:    true,
// 	}

// 	resp, err := client.CreateCompletionStream(ctx, req)
// 	if err != nil {
// 		fmt.Println("createcompletionstream error")
// 		fmt.Errorf("CreateCompletionStream returned error: %v", err)
// 	}
// 	defer resp.Close()
	
// 	counter := 0
// 	for {
// 		data, err := resp.Recv()
// 		if err != nil {
// 			if errors.Is(err, io.EOF) {
// 				fmt.Println("getresponse error")
// 				break
// 			}
// 			fmt.Errorf("Stream error: %v", err)
// 			} else {
// 				counter++
// 				fmt.Print(data.Choices[0].Text)
				
// 			}
// 		}
// 		if counter == 0 {
// 		fmt.Println("check")
// 		fmt.Errorf("Stream did not return any responses")
// 	}
// 	fmt.Println("")
// }

// type NullWriter int

// func (NullWriter) Write([]byte) (int, error) { return 0, nil }

// func main() {
// 	log.SetOutput(new(NullWriter))
// 	apiKey := os.Getenv("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")
// 	// if apiKey == "" {
// 	// 	panic("Missing API KEY")
// 	// }

// 	ctx := context.Background()
// 	client := gpt3.NewClient(apiKey)
// 	rootCmd := &cobra.Command{
// 		Use:   "chatgpt",
// 		Short: "Chat with ChatGPT in console.",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			scanner := bufio.NewScanner(os.Stdin)
// 			quit := false

// 			for !quit {
// 				fmt.Print("Input your question (type `quit` to exit): ")

// 				if !scanner.Scan() {
// 					break
// 				}

// 				question := scanner.Text()
// 				// questionParam := validateQuestion(question)
// 				switch question {
// 				case "quit":
// 					quit = true
// 				default:
// 					fmt.Println(question)
// 					GetResponse(client, ctx, question)
// 				}
// 			}
// 		},
// 	}

// 	log.Fatal(rootCmd.Execute())
// }

// func validateQuestion(question string) string {
// 	quest := strings.Trim(question, " ")
// 	keywords := []string{"", "loop", "break", "continue", "cls", "exit", "block"}
// 	for _, x := range keywords {
// 		if quest == x {
// 			return ""
// 		}
// 	}
// 	return quest
// }

package main

import (
    "fmt"
    "log"
	"context"
	"time"
    "os"
    "github.com/sashabaranov/go-openai"
)

func main() {
    // 获取OpenAI API key
    apiKey := os.Getenv("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")

    // 创建OpenAI客户端
    client := openai.NewClient(apiKey)
	req := openai.CompletionRequest{
		Model:       "text-davinci-001", // 模型名称或 ID
		Prompt:      "Once upon a time", // 输入的文本提示
		MaxTokens:   50,                 // 生成的最大标记数量
		// 可以添加其他参数，比如Temperature、Stop等
	}
	parentCtx := context.Background()                             // 创建一个顶层父 Context
	ctxWithTimeout, cancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer cancel()
	resp, err := client.CreateCompletionStream(ctxWithTimeout ,req)
    // 调用OpenAI API
    // resp, err := client.Completion.Create(
    //     "The quick brown fox jumps over the lazy dog",
    //     "The quick brown fox jumps over the",
    //     10,
    // )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp)
}
