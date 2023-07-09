package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
	"net/http"
)

func Test(c *gin.Context) {

	token := "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"

	owner := "sskail"
	repoName := "scholar_mindmap"

	client := getGitHubClient(token)
	fileContent, directoryContent, err := getRepoContents(client, owner, repoName, "/")
	if err != nil {
		fmt.Printf("Error getting repository info: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting repository info"})
		return
	}

	if fileContent == nil && directoryContent == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "File content is empty"})
		return
	}

	// Prepare data to return as a slice of maps with file names and their content
	var data []map[string]string

	if fileContent != nil {
		content, err := fileContent.GetContent()
		if err != nil {
			fmt.Printf("Error getting content: %s\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting content"})
			return
		}

		data = append(data, map[string]string{
			"name":    fileContent.GetName(),
			"content": content,
		})
	}

	if directoryContent != nil {
		for _, file := range directoryContent {
			if file.GetType() == "file" {
				content, err := file.GetContent()
				if err != nil {
					fmt.Printf("Error getting content: %s\n", err)
					continue
				}

				data = append(data, map[string]string{
					"name":    file.GetName(),
					"content": content,
				})
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func getGitHubClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func getRepoInfo(client *github.Client, owner, repoName string) (*github.Repository, error) {
	ctx := context.Background()
	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	return repo, err
}

func getRepoContents(client *github.Client, owner, repoName, path string) (*github.RepositoryContent, []*github.RepositoryContent, error) {
	ctx := context.Background()
	opts := &github.RepositoryContentGetOptions{}

	fileContent, directoryContent, _, err := client.Repositories.GetContents(ctx, owner, repoName, path, opts)

	return fileContent, directoryContent, err
}

func GetCommitHistory(c *gin.Context) {
	token := "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"
	owner := "sskail"
	repoName := "scholar_mindmap"

	client := getGitHubClient(token)

	commits, err := getCommitHistory(client, owner, repoName)
	if err != nil {
		fmt.Printf("Error getting commit history: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting commit history"})
		return
	}

	var commitList []map[string]string
	for _, commit := range commits {
		commitList = append(commitList, map[string]string{
			"sha":     commit.GetSHA(),
			"message": commit.GetCommit().GetMessage(),
			"author":  commit.GetCommit().GetAuthor().GetName(),
			"date":    commit.GetCommit().GetAuthor().GetDate().String(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": commitList})
}

func getCommitHistory(client *github.Client, owner, repoName string) ([]*github.RepositoryCommit, error) {
	ctx := context.Background()
	opts := &github.CommitsListOptions{}

	commits, _, err := client.Repositories.ListCommits(ctx, owner, repoName, opts)
	return commits, err
}
