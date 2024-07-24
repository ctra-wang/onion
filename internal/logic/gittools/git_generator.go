package gittools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ctra-wang/onion/internal/logic/model"
	"github.com/google/go-github/v45/github"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func GitGenerator(repositoryParams model.RepositoryParams) {
	switch repositoryParams.Platform {
	// github
	case 1:
		// 1. Set up OAuth2 authentication.
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: repositoryParams.ApiToken},
		)
		tc := oauth2.NewClient(ctx, ts)

		// 2. Create a GitHub client.
		client := github.NewClient(tc)

		// 3. Create a new repository.
		repo := &github.Repository{
			Name:    github.String(repositoryParams.RepoName),
			Private: github.Bool(repositoryParams.RepoPermission), // true for private, false for public
		}

		_, _, err := client.Repositories.Create(ctx, repositoryParams.OrgName, repo)
		if err != nil {
			if _, ok := err.(*github.RateLimitError); ok {
				fmt.Println("Github Rate limit exceeded.")
			} else if ghErr, ok := err.(*github.ErrorResponse); ok {
				if ghErr.Response.StatusCode == 403 {
					fmt.Println("Github Access forbidden: You might not have the necessary permissions or the token might be invalid.")
				} else {
					fmt.Printf("GitHub API error: %s\n", ghErr.Message)
				}
			} else {
				fmt.Printf("Github Unexpected error: %s\n", err)
			}
			return
		}

		fmt.Println("Github Repository created successfully!")
		// Gitee
	case 2:
		data := map[string]interface{}{
			"name":        repositoryParams.RepoName,
			"private":     repositoryParams.RepoPermission, // true for private, false for public
			"description": "",
		}
		jsonData, _ := json.Marshal(data)
		req, err := http.NewRequest("POST", fmt.Sprintf("https://gitee.com/api/v5/orgs/%s/repos", repositoryParams.OrgName), bytes.NewBuffer(jsonData))

		if err != nil {
			fmt.Println("Gitee Error creating Gitee request:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("token %s", repositoryParams.ApiToken))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusCreated {
			fmt.Println("Github Repository created successfully!")
		} else {
			fmt.Println("Failed to create repository, status code:", resp.StatusCode)
		}

		// Gitea
	case 3:
		data := map[string]interface{}{
			"name":    repositoryParams.RepoName,
			"private": repositoryParams.RepoPermission, // true for private, false for public
		}

		jsonData, _ := json.Marshal(data)
		req, err := http.NewRequest("POST", "https://gitea.example.com/api/v1/user/repos", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Gitea Error sending request:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "token "+repositoryParams.ApiToken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			fmt.Println("Gitea Failed to create repository")
			return
		}

		fmt.Println("Gitea Repository created successfully!")
		// Gitea
	case 4:
		client, err := gitlab.NewClient(repositoryParams.ApiToken)
		if err != nil {
			log.Fatal(err)
		}

		visibility := gitlab.PublicVisibility
		if repositoryParams.RepoPermission {
			visibility = gitlab.PrivateVisibility
		}

		// 设置项目选项
		opt := &gitlab.CreateProjectOptions{
			Name:       gitlab.String(repositoryParams.RepoName),
			Visibility: gitlab.Visibility(visibility), // 如果这个字段被弃用，使用 VisibilityValue
		}

		// 创建新项目
		_, _, err = client.Projects.CreateProject(opt)
		if err != nil {
			log.Fatal(err)
		}
	}

}
