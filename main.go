package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	repoOwner          = "Ajmal-MP"
	repoName           = "swarmcd"
	lastCommitFilePath = "./lastcommit.txt"
)

func main() {
	// getting last commit from file

	for {
		binaryLastcommitInFile, err := os.ReadFile(lastCommitFilePath)
		if err != nil {
			panic(err)
		}
		lastcommitInFile := string(binaryLastcommitInFile)

		// checking last commit available or not
		if lastcommitInFile != "" {
			//geting last commit id
			lastCommitGithb := githubLastcommit()

			if lastcommitInFile[:10] == lastCommitGithb[:10] {
				fmt.Println("both commit are same Nothing to do wait one for minute")
				time.Sleep(time.Minute)
			} else {
				//writing the latest commit to lastcommit.txt file
				lastCommitFilePath, error := os.Create(lastCommitFilePath)
				_, error = io.WriteString(lastCommitFilePath, lastCommitGithb)
				if error != nil {
					panic(error)
				}
				fmt.Println("deplyed success fully")
			}
		} else {
			panic("No data in lastcommit.txt file")
		}
	}
}

func githubLastcommit() string {
	remoteURL := "https://github.com/" + repoOwner + "/" + repoName + ".git"
	cmd := exec.Command("git", "ls-remote", remoteURL, "HEAD")
	byteoutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	output := string(byteoutput)
	output = strings.Replace(output, "HEAD", "", -1)
	return output
}
