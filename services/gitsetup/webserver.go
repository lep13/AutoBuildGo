package gitsetup

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/lep13/AutoBuildGo/services/ecr"
)

type RepoRequest struct {
	RepoName    string `json:"repo_name"`
	Description string `json:"description"`
}

func HandleWebServer() {
	http.HandleFunc("/create-repo", CreateRepoHandler)
	log.Println("Server is starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func CreateRepoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateRepoHandler invoked")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RepoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if req.RepoName == "" {
		http.Error(w, "Repository name is required", http.StatusBadRequest)
		return
	}

	description := req.Description
	if description == "" {
		description = "Created from a template via automated setup"
	}

	// Create ECR client
	ecrClient, err := ecr.CreateECRClient()
	if err != nil {
		http.Error(w, "Failed to create ECR client: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create ECR Repository
	if err := ecr.CreateRepo(req.RepoName, ecrClient); err != nil {
		http.Error(w, "Failed to create ECR repository: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Ensure environment is loaded
	LoadEnv()

	// Create Git Repository
	config := DefaultRepoConfig(req.RepoName, description)
	gitClient := NewGitClient() // Create an instance of GitClient

	if err := gitClient.CreateGitRepository(config); err != nil {
		http.Error(w, "Failed to create Git repository: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 20 second time delay
	time.Sleep(20 * time.Second)

	// Clone the repo, update go.mod, and push changes
	if err := CloneAndPushRepo(req.RepoName); err != nil {
		http.Error(w, "Failed to clone and push repository: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ECR and Git repositories created successfully"))
}
