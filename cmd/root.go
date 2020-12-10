package cmd

import (
	"fmt"
	"os"

	"github.com/jthanio/advent2020/internal"
	"github.com/jthanio/advent2020/pkg/day1"
	"github.com/jthanio/advent2020/pkg/day10"
	"github.com/jthanio/advent2020/pkg/day2"
	"github.com/jthanio/advent2020/pkg/day3"
	"github.com/jthanio/advent2020/pkg/day4"
	"github.com/jthanio/advent2020/pkg/day5"
	"github.com/jthanio/advent2020/pkg/day6"
	"github.com/jthanio/advent2020/pkg/day7"
	"github.com/jthanio/advent2020/pkg/day8"
	"github.com/jthanio/advent2020/pkg/day9"
	"github.com/spf13/cobra"
)

var input string

func init() {
	// Add all day commands as subcommands, with configured flags
	for i, c := range dayCommands {
		c.Flags().StringP("input", "i", fmt.Sprintf("inputs/day%d.txt", i+1), "the puzzle input file")
		rootCmd.AddCommand(c)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "advent2020",
	Short: "advent2020 is an application that runs advent of code 2020 solvers.",
}

var dayCommands = []*cobra.Command{
	{
		Use: "day1",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := internal.LoadIntInputFile(cmd.Flags().Lookup("input").Value.String())
			if err != nil {
				return err
			}

			fmt.Printf("The answer to the Day 1 part 1 puzzle: %d\n", day1.SolveDay1Part1(input))
			fmt.Printf("The answer to the Day 1 part 2 puzzle: %d\n", day1.SolveDay1Part2(input))
			return nil
		},
	},
	{
		Use: "day2",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := internal.LoadStringInputFile(cmd.Flags().Lookup("input").Value.String())
			if err != nil {
				return err
			}

			a1, err := day2.SolveDay2Part1(input)
			if err != nil {
				return err
			}

			a2, err := day2.SolveDay2Part2(input)
			if err != nil {
				return err
			}

			fmt.Printf("The answer to the Day 2 part 1 puzzle: %d\n", a1)
			fmt.Printf("The answer to the Day 2 part 2 puzzle: %d\n", a2)
			return nil
		},
	},
	{
		Use: "day3",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := internal.LoadStringInputFile(cmd.Flags().Lookup("input").Value.String())
			if err != nil {
				return err
			}

			fmt.Printf("The answer to the Day 3 part 1 puzzle: %d\n", day3.SolveDay3Part1(input))
			fmt.Printf("The answer to the Day 3 part 2 puzzle: %d\n", day3.SolveDay3Part2(input))
			return nil
		},
	},
	{
		Use: "day4",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := internal.LoadStringInputFile(cmd.Flags().Lookup("input").Value.String())
			if err != nil {
				return err
			}

			fmt.Printf("The answer to the Day 4 part 1 puzzle: %d\n", day4.SolveDay4Part1(input))
			fmt.Printf("The answer to the Day 4 part 2 puzzle: %d\n", day4.SolveDay4Part2(input))
			return nil
		},
	},
	{
		Use: "day5",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := internal.LoadStringInputFile(cmd.Flags().Lookup("input").Value.String())
			if err != nil {
				return err
			}

			a1, err := day5.SolveDay5Part1(input)
			if err != nil {
				return err
			}

			a2, err := day5.SolveDay5Part2(input)
			if err != nil {
				return err
			}

			fmt.Printf("The answer to the Day 5 part 1 puzzle: %d\n", a1)
			fmt.Printf("The answer to the Day 5 part 2 puzzle: %d\n", a2)
			return nil
		},
	},
	{
		Use: "day6",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := internal.LoadStringInputFile(cmd.Flags().Lookup("input").Value.String())
			if err != nil {
				return err
			}

			fmt.Printf("The answer to the Day 6 part 1 puzzle: %d\n", day6.SolveDay6Part1(input))
			fmt.Printf("The answer to the Day 6 part 2 puzzle: %d\n", day6.SolveDay6Part2(input))
			return nil
		},
	},
	{
		Use: "day7",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := internal.LoadStringInputFile(cmd.Flags().Lookup("input").Value.String())
			if err != nil {
				return err
			}

			a1, err := day7.SolveDay7Part1(input)
			if err != nil {
				return err
			}

			a2, err := day7.SolveDay7Part2(input)
			if err != nil {
				return err
			}

			fmt.Printf("The answer to the Day 7 part 1 puzzle: %d\n", a1)
			fmt.Printf("The answer to the Day 7 part 2 puzzle: %d\n", a2)
			return nil
		},
	},
	{
		Use: "day8",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := internal.LoadStringInputFile(cmd.Flags().Lookup("input").Value.String())
			if err != nil {
				return err
			}

			a1, err := day8.SolveDay8Part1(input)
			if err != nil {
				return err
			}

			a2, err := day8.SolveDay8Part2(input)
			if err != nil {
				return err
			}

			fmt.Printf("The answer to the Day 8 part 1 puzzle: %d\n", a1)
			fmt.Printf("The answer to the Day 8 part 2 puzzle: %d\n", a2)
			return nil
		},
	},
	{
		Use: "day9",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := internal.LoadIntInputFile(cmd.Flags().Lookup("input").Value.String())
			if err != nil {
				return err
			}

			a1, err := day9.SolveDay9Part1(input)
			if err != nil {
				return err
			}

			a2, err := day9.SolveDay9Part2(input)
			if err != nil {
				return err
			}

			fmt.Printf("The answer to the Day 9 part 1 puzzle: %d\n", a1)
			fmt.Printf("The answer to the Day 9 part 2 puzzle: %d\n", a2)
			return nil
		},
	},
	{
		Use: "day10",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := internal.LoadIntInputFile(cmd.Flags().Lookup("input").Value.String())
			if err != nil {
				return err
			}

			a1, err := day10.SolveDay10Part1(input)
			if err != nil {
				return err
			}

			a2, err := day10.SolveDay10Part2(input)
			if err != nil {
				return err
			}

			fmt.Printf("The answer to the Day 10 part 1 puzzle: %d\n", a1)
			fmt.Printf("The answer to the Day 10 part 2 puzzle: %d\n", a2)
			return nil
		},
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
