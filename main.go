package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alex-schaaf/recap-challenge/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	handleErrFatal(err)
	if os.Getenv("API_KEY") == "" {
		log.Fatal("Required `API_KEY` environment variable not set")
	}

	flags, err := parseFlags()
	handleErrFatal(err)

	apiResponse, err := utils.FetchSymbolData(flags.symbol, flags.start, flags.end)
	handleErrFatal(err)

	var adjClose []float64
	for _, d := range apiResponse.Datatable.Data {
		adjClose = append(adjClose, d.Adj_close)
	}

	simpleReturn := utils.SimpleReturn(adjClose)
	maxDrawdown := utils.MaximumDrawdown(adjClose)

	message := formatMessage(flags, simpleReturn, maxDrawdown)

	if flags.isSendEmail {
		err := utils.SendEmail(fmt.Sprintf("Update on Symbol %s", flags.symbol), message)
		handleErrFatal(err)
	} else {
		fmt.Println(message)
	}
}

func formatMessage(flags Flags, simpleReturn, maxDrawdown float64) string {
	header := fmt.Sprintf("%s (%s to %s)\n\n", flags.symbol, flags.start, flags.end)
	simpleReturnStatement := fmt.Sprintf("%-20s %8.2f %%\n", "Simple Return", simpleReturn)
	maxDrawdownStatement := fmt.Sprintf("%-20s %8.2f %%", "Maximum Drawdown", maxDrawdown)

	return header + simpleReturnStatement + maxDrawdownStatement
}

type Flags struct {
	symbol      string
	start       string
	end         string
	isSendEmail bool
}

func parseFlags() (Flags, error) {
	symbolPtr := flag.String("symbol", "", "The stock symbol for which to calculate the simple return")
	startPtr := flag.String("start", "", "The start date (YYYY-MM-DD)")
	endPtr := flag.String("end", "", "The end date (YYYY-MM-DD)")
	isSendEmailPtr := flag.Bool("email", false, "Send result per email")

	flag.Parse()

	symbol := *symbolPtr
	start := *startPtr
	end := *endPtr

	if symbol == "" {
		log.Fatal("No symbol provided")
	}
	if start == "" {
		log.Fatal("No start date provided")
	}
	if end == "" {
		log.Fatal("No end date provided")
	}
	if *isSendEmailPtr && (os.Getenv("EMAIL_SENDER") == "" ||
		os.Getenv("EMAIL_RECEIVER") == "" ||
		os.Getenv("AWS_ACCESS_KEY") == "" ||
		os.Getenv("AWS_SECRET_KEY") == "") {
		log.Fatal("Environment variables `EMAIL_SENDER`, `EMAIL_SENDER`, `AWS_ACCESS_KEY` and `AWS_SECRET_KEY` must be available to send emails ")
	}

	return Flags{symbol, start, end, *isSendEmailPtr}, nil
}

func handleErrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
