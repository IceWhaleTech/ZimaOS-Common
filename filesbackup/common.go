package filesbackup

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const ServiceName = "files-backup"

func DefaultMetadataPath(dataPath string) string {
	return filepath.Join(dataPath, ServiceName)
}

func GetAllBackups[M any](metadataPath string) (map[string][]M, error) {
	// walk thru metadataPath and load each file starting with "backup_" in JSON format as a FolderBackup into a map, and return
	// the map.

	allBackups := map[string][]M{}

	if err := filepath.WalkDir(metadataPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasPrefix(d.Name(), "backup_") {
			return nil
		}

		backup, err := LoadMetadata[M](path)
		if err != nil {
			return fs.SkipDir
		}

		// clientID is the directory name of path
		clientID := filepath.Base(filepath.Dir(path))

		allBackups[clientID] = append(allBackups[clientID], *backup)

		return nil
	}); err != nil {
		return nil, err
	}

	return allBackups, nil
}

func LoadMetadata[Backup any](metadataFilePath string) (*Backup, error) {
	metadataFile, err := os.Open(metadataFilePath)
	if err != nil {
		return nil, err
	}
	defer metadataFile.Close()

	decoder := json.NewDecoder(metadataFile)
	var backup Backup
	if err := decoder.Decode(&backup); err != nil {
		return nil, err
	}

	return &backup, nil
}
