package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

type GithubRelease struct {
	Assets []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func main() {
	app := &cli.App{
		Name:                 "werft",
		Usage:                "setting up Semantic MediaWiki without the hassle",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "spawn",
				Usage: "Download and set up Semantic MediaWiki files",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "Specify the extraction path",
						Value:   "",
					},
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Specify the output directory name",
						Value:   "werft-smw-image",
					},
				},
				Action: runSpawn,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runSpawn(c *cli.Context) error {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Working: "
	s.Start()

	// Replace with your GitHub repo details
	owner := "uniwue-zpd"
	repo := "werft"

	extractPath := c.String("path")
	if extractPath == "" {
		var err error
		extractPath, err = os.Getwd()
		if err != nil {
			s.Stop()
			return fmt.Errorf("error getting current working directory: %v", err)
		}
	}

	outputDir := c.String("output")
	fullOutputPath := filepath.Join(extractPath, outputDir)

	s.Suffix = " Fetching latest release..."
	customZipURL, err := getCustomZipURL(owner, repo)
	if err != nil {
		s.Stop()
		return fmt.Errorf("error getting custom.zip URL: %v", err)
	}

	s.Suffix = " Downloading custom.zip..."
	zipPath := filepath.Join(extractPath, "custom.zip")
	err = downloadFile(customZipURL, zipPath)
	if err != nil {
		s.Stop()
		return fmt.Errorf("error downloading custom.zip: %v", err)
	}

	s.Suffix = " Extracting ZIP file..."
	err = unzip(zipPath, fullOutputPath)
	if err != nil {
		s.Stop()
		return fmt.Errorf("error extracting ZIP file: %v", err)
	}

	s.Suffix = " Renaming files..."
	renameFiles(fullOutputPath)

	s.Suffix = " Cleaning up..."
	err = cleanup(fullOutputPath)
	if err != nil {
		return err
	}

	err = os.Remove(zipPath)
	if err != nil {
		return err
	}

	s.Stop()
	fmt.Printf("Setup completed successfully in %s!\n", fullOutputPath)
	return nil
}

func getCustomZipURL(owner, repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var release GithubRelease
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return "", err
	}

	for _, asset := range release.Assets {
		if asset.Name == "custom.zip" {
			return asset.BrowserDownloadURL, nil
		}
	}

	return "", fmt.Errorf("custom.zip not found in the latest release")
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	os.MkdirAll(dest, os.ModePerm)

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func renameFiles(dir string) {
	renamePairs := map[string]string{
		"Dockerfile.template":          "Dockerfile",
		"docker-compose.template.yaml": "docker-compose.yaml",
		"template.env":                 ".env",
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		for oldName, newName := range renamePairs {
			if strings.HasSuffix(path, oldName) {
				newPath := filepath.Join(filepath.Dir(path), newName)
				if err := os.Rename(path, newPath); err != nil {
					fmt.Printf("Error renaming %s to %s: %v\n", path, newPath, err)
				}
				break
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through directory: %v\n", err)
	}
}

func cleanup(path string) error {
	customDir := filepath.Join(path, "custom")

	if _, err := os.Stat(customDir); os.IsNotExist(err) {
		return fmt.Errorf("directory 'custom' does not exist in the path: %s", path)
	}

	err := filepath.Walk(customDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(customDir, srcPath)
		if err != nil {
			return err
		}

		destPath := filepath.Join(path, relativePath)

		if info.IsDir() {
			if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
				return err
			}
		} else {
			if err := moveFile(srcPath, destPath); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	if err := os.RemoveAll(customDir); err != nil {
		return fmt.Errorf("failed to remove 'custom' directory: %v", err)
	}

	return nil
}

func moveFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return err
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return err
	}

	if err := os.Remove(src); err != nil {
		return err
	}

	return nil
}
