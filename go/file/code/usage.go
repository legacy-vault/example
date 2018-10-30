// usage.go.

package main

import (
	"fmt"

	"github.com/legacy-vault/library/go/file"
)

func main() {

	var files []string
	var folders []string
	var items []string
	var path string

	// Prepare Data.
	path = "/something/here" // <- Fill it as you wish.
	fmt.Println("Current Directory:")
	fmt.Println(path)

	// List Files without Sub-Levels.
	fmt.Println("Files (1 Level):")
	files = file.ListFiles(path, false)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Files with Sub-Levels.
	fmt.Println("Files (All Levels):")
	files = file.ListFiles(path, true)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Directories without Sub-Levels.
	fmt.Println("Folders (1 Level):")
	folders = file.ListFolders(path, false)
	for _, folder := range folders {
		fmt.Println(folder)
	}

	// List Directories with Sub-Levels.
	fmt.Println("Folders (All Levels):")
	folders = file.ListFolders(path, true)
	for _, folder := range folders {
		fmt.Println(folder)
	}

	// List Files & Folders without Sub-Levels.
	fmt.Println("Files & Folders (1 Level):")
	items = file.ListFilesAndFolders(path, false)
	for _, item := range items {
		fmt.Println(item)
	}

	// List Files & Folders with Sub-Levels.
	fmt.Println("Files & Folders (All Levels):")
	items = file.ListFilesAndFolders(path, true)
	for _, item := range items {
		fmt.Println(item)
	}

	// List Edge Directories.
	fmt.Println("Edge Folders:")
	folders = file.ListEdgeFolders(path)
	for _, folder := range folders {
		fmt.Println(folder)
	}

	// List Edge Files.
	fmt.Println("Edge Files:")
	files = file.ListEdgeFiles(path)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Files, Level <= 2.
	fmt.Println("Files (Level <= 2):")
	files = file.ListFilesLR(
		path,
		file.LevelRestrictionsCriterionLessOrEqual,
		2,
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Files, Level = 2.
	fmt.Println("Files (Level = 2):")
	files = file.ListFilesLR(
		path,
		file.LevelRestrictionsCriterionEqual,
		2,
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Files, Level >= 2.
	fmt.Println("Files (Level >= 2):")
	files = file.ListFilesLR(
		path,
		file.LevelRestrictionsCriterionGreaterOrEqual,
		2,
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Folders, Level <= 2.
	fmt.Println("Folders (Level <= 2):")
	files = file.ListFoldersLR(
		path,
		file.LevelRestrictionsCriterionLessOrEqual,
		2,
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Folders, Level = 2.
	fmt.Println("Folders (Level = 2):")
	files = file.ListFoldersLR(
		path,
		file.LevelRestrictionsCriterionEqual,
		2,
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Folders, Level >= 2.
	fmt.Println("Folders (Level >= 2):")
	files = file.ListFoldersLR(
		path,
		file.LevelRestrictionsCriterionGreaterOrEqual,
		2,
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Files & Folders, Level <= 2.
	fmt.Println("Files & Folders (Level <= 2):")
	files = file.ListFilesAndFoldersLR(
		path,
		file.LevelRestrictionsCriterionLessOrEqual,
		2,
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Files & Folders, Level = 2.
	fmt.Println("Files & Folders (Level = 2):")
	files = file.ListFilesAndFoldersLR(
		path,
		file.LevelRestrictionsCriterionEqual,
		2,
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Files & Folders, Level >= 2.
	fmt.Println("Files & Folders (Level >= 2):")
	files = file.ListFilesAndFoldersLR(
		path,
		file.LevelRestrictionsCriterionGreaterOrEqual,
		2,
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Files with allowed Extensions.
	fmt.Println("Files (All Levels, Allowed Extensions List):")
	files = file.ListFilesExtAllowed(
		path,
		true,
		[]string{"txt", "torrent"},
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}

	// List Files without forbidden Extensions.
	fmt.Println("Files (All Levels, Forbidden Extensions List):")
	files = file.ListFilesExtForbidden(
		path,
		true,
		[]string{"txt"},
	)
	for _, aFile := range files {
		fmt.Println(aFile)
	}
}
