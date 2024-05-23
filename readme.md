Command to create ECR and Github Repositories:
go run main.go -repo = <reponame>   

Command to test:
Move to the directory of respective test case files for ECR and Github
go test

Note:
Make sure that reponame is in correct format
format: (?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*

Valid Repository Names:

valid-repo
valid_repo
valid.repo
valid/repo/name
valid-repo123
123
valid123

Invalid Repository Names:

-invalid-repo (Leading hyphen)
invalid-repo- (Trailing hyphen)
invalid--repo (Consecutive hyphens)
/invalid-repo (Leading slash)
invalid-repo/ (Trailing slash)
InvalidRepo (Uppercase letters)
répö-nämé (Non-ASCII characters)
invalid repo (Contains spaces)
invalid@repo (Contains special characters)
