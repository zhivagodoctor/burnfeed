package artifact

// Each artifact is a saved file on IPFS
// it is (first) a JSON formatted text.
type Artifact struct {
	CID string
}

// ArtifactRepository defines the interface for artifact data storage.
// Basically responsible for CRUD operations.
type ArtifactRepository interface {
	CreateArtifact(artifactType, artifactJsonString string) (string, error)
	GetArtifactByID(ID string) (*Artifact, error)
	GetUserArtifacts(wallet string) (string, error)
}

// ArtifactService defines the interface for business logic.
type ArtifactService interface {
	CreateArtifact(artifactType, artifactJsonString string) (string, error)
	GetArtifactByID(ID string) (*Artifact, error)
	GetUserArtifacts(wallet string) (string, error)
}
