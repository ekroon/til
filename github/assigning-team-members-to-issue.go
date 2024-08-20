package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func getTeamMembers(org, team string) (string, error) {
    cmd := exec.Command("gh", "api", fmt.Sprintf("orgs/%s/teams/%s/members", org, team), "--jq", "map(.login) | join(\",\")")
    output, err := cmd.Output()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(output)), nil
}

func editIssue(org, repo, issueNumber, assignees string) error {
    cmd := exec.Command("gh", "issue", "edit", issueNumber, "--repo", fmt.Sprintf("%s/%s", org, repo), "--add-assignee", assignees)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func main() {
    org := flag.String("org", "", "GitHub organization")
    team := flag.String("team", "", "GitHub team name")
    repo := flag.String("repo", "", "GitHub repository")
    issueNumber := flag.String("issue", "", "GitHub issue number or link")
    flag.Parse()

    if *org == "" || *team == "" || *repo == "" || *issueNumber == "" {
        flag.Usage()
        os.Exit(1)
    }

    members, err := getTeamMembers(*org, *team)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error getting team members: %v\n", err)
        os.Exit(1)
    }

    if err := editIssue(*org, *repo, *issueNumber, members); err != nil {
        fmt.Fprintf(os.Stderr, "Error editing issue: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Issue edited successfully")
}
