/*
 Operations related to searching in commands
*/

package search

import (
	"sort"
	"strings"

	"github.com/agnivade/levenshtein"

	"github.com/skueee/pense/pkg/configfile"
)

// A struct to save a command with a score
type ScoredCommand struct {
	Command configfile.Command
	Score   float64
}

// Takes two strings and returns a similarity score
func SearchAlgo(request string, target string) float64 {

	if request == target {
		return 1.0
	}

	dist := levenshtein.ComputeDistance(request,target)
	maxLen := max(len(request), len(target))
    if maxLen == 0 {
    	return 0
    }

  	similarity := 1.0 - (float64(dist) / float64(maxLen))

    if similarity > 0.8 {
        return similarity
    }

   return 0
}

// Takes all the keywords from a command and returns the best score
func SearchInKeywords(request string, command configfile.Command) float64 {
	var bestMatch float64
	for i := 0; i < len(command.Keywords); i++ {
		result := SearchAlgo(request, command.Keywords[i])
		if result > bestMatch {
			bestMatch = result
		}
	}
	return bestMatch
}

// Main function, takes a command and performs all the searches then sort the commands
func Search(request string, commands []configfile.Command) []configfile.Command {
	var searchResults []ScoredCommand

    for _, cmd := range commands {
    	nameScore := SearchAlgo(request, cmd.Name) * 2
        noteScore := SearchAlgo(request, cmd.Note) * 0.8
        keywordsScore := SearchInKeywords(request, cmd) * 2

        // Cuts program and args, because the user is more likely to search for the program than the args
        program, args, _ := strings.Cut(cmd.Command, " ")
        programScore := SearchAlgo(request, program) * 1.5
        argsScore := SearchAlgo(request, args) * 0.5

        totalScore := nameScore + noteScore + keywordsScore + programScore + argsScore

        if totalScore > 0.75 {
            searchResults = append(searchResults, ScoredCommand{
                Command: cmd,
                Score:   totalScore,
            })
        }
    }

    sort.Slice(searchResults, func(i, j int) bool {
        return searchResults[i].Score > searchResults[j].Score
    })

    // Retransforms all the ScoreCommand into Command
    var results []configfile.Command
    for _, i := range searchResults {
        results = append(results, i.Command)
    }

    return results
}
