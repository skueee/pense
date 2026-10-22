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

type ScoredCommand struct {
	Command configfile.Command
	Score   float64
}

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

func Search(request string, commands []configfile.Command) []configfile.Command {
	var searchResults []ScoredCommand

    for _, cmd := range commands {
    	nameScore := SearchAlgo(request, cmd.Name) * 2
        noteScore := SearchAlgo(request, cmd.Note) * 0.8
        keywordsScore := SearchInKeywords(request, cmd) * 2

        program, args, _ := strings.Cut(cmd.Command, " ")
        programScore := SearchAlgo(request, program) * 1.5
        argsScore := SearchAlgo(request, args) * 0.5

        totalScore := nameScore + noteScore + keywordsScore + programScore + argsScore

        if totalScore > 5 {
            searchResults = append(searchResults, ScoredCommand{
                Command: cmd,
                Score:   totalScore,
            })
        }
    }

    sort.Slice(searchResults, func(i, j int) bool {
        return searchResults[i].Score > searchResults[j].Score
    })

    var results []configfile.Command
    for _, i := range searchResults {
        results = append(results, i.Command)
    }

    return results
}
