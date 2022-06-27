package controls

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joaquinalvarezdev/entropy/models"
	"github.com/joaquinalvarezdev/entropy/utils"
)

func AnalyzeFile(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(0)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	blockSize, _ := strconv.Atoi(r.FormValue("blocksize"))
	if blockSize == 0 {
		// defaults to 1024 if not indicated.
		blockSize = 1024
	}
	reader := bufio.NewReader(file)
	buf := make([]byte, 0, blockSize)
	var entropyResponse models.EntropyResponse
	lowEntropy := 0
	mediumEntropy := 0
	highEntropy := 0

	for {
		n, err := io.ReadFull(reader, buf[:cap(buf)])
		fmt.Println(n)
		buf = buf[:n]
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != io.ErrUnexpectedEOF {
				fmt.Fprintln(os.Stderr, err)
				break
			}
		}
		entropy, _ := utils.CalculateEntropy(buf, n)
		fmt.Println("Entropy", entropy)
		entropyResponse.EntropyDetail = append(entropyResponse.EntropyDetail, entropy)

		switch {
		case entropy < 2:
			lowEntropy += 1
		case entropy > 7:
			highEntropy += 1
		default:
			mediumEntropy += 1
		}

	}

	summaryMap := map[string]int{
		"lowEntropy":    lowEntropy,
		"mediumEntropy": mediumEntropy,
		"highEntropy":   highEntropy,
	}

	entropyResponse.Summary = summaryMap

	json.NewEncoder(w).Encode(entropyResponse)
}
