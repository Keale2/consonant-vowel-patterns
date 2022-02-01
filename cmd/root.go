package cmd

import (
	"fmt"
	"os"

	"github.com/Keale2/consonant-vowel-patterns/internal/analysis"
	"github.com/Keale2/consonant-vowel-patterns/internal/files"
	"github.com/Keale2/consonant-vowel-patterns/internal/util"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "consonant-vowel-patterns",
		Short: "Score the words.",
		Long:  `Score the words. Score the words. Score the words. Score the words. Score the words. `,
		RunE: func(cmd *cobra.Command, args []string) error {
			// loop through the slice of strings
			// calculate a score for each word
			// write a file (--output-file) which is the slice of words, sorted by score
			// fmt.Println(scoreWords(w))

			// take a file (--words-file)
			words, err := files.ReadFile(inputFile)
			if err != nil {
				return err
			}

			results := analysis.Analyze(words)
			files.WriteFile(outputFile, util.SortMapByValue(results).String())

			return nil
		},
	}
	inputFile  string
	outputFile string
)

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "The input file")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "pattern-frequency.txt", "The output file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
