# Assigning team members to an issue 

It is not possible to assign a team to an issue in GitHub, the alternative is to assign the team members itself to an issue with `gh-cli`:

```bash
members=$(gh api orgs/ORG/teams/TEAM-NAME/members --jq 'map(.login) | join(",")')
gh issue edit NUMBER --repo ORG/REPO --add-assignee $members
```
