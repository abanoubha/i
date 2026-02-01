package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	GithubRepo  = "abanoubha/i"
	InstallName = "i"
	DefaultDir  = "/usr/local/bin"
)

type Release struct {
	Assets []struct {
		Name        string `json:"name"`
		DownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func installLatestVersion() {
	assetName, err := detectAsset()
	if err != nil {
		fail("Detection failed: %v", err)
	}
	fmt.Printf("Detected system: %s/%s. Looking for asset: %s\n", runtime.GOOS, runtime.GOARCH, assetName)

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", GithubRepo)
	downloadURL, err := getDownloadURL(apiURL, assetName)
	if err != nil {
		fail("Failed to find download URL: %v", err)
	}

	fmt.Printf("Downloading: %s\n", downloadURL)
	tmpFile, err := os.CreateTemp("", "i-installer-*")
	if err != nil {
		fail("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // clean up on exit
	defer tmpFile.Close()

	if err := downloadFile(downloadURL, tmpFile); err != nil {
		fail("Download failed: %v", err)
	}
	tmpFile.Close() // Close 'explicitly' before moving/copying

	installDir := os.Getenv("INSTALL_DIR")
	if installDir == "" {
		installDir = DefaultDir
	}
	targetPath := filepath.Join(installDir, InstallName)

	fmt.Printf("Installing to %s...\n", targetPath)
	if err := installBinary(tmpFile.Name(), targetPath); err != nil {
		fail("Installation failed: %v", err)
	}

	fmt.Printf("[info] successfully installed '%s' to '%s'\n", InstallName, targetPath)
}

func detectAsset() (string, error) {
	switch runtime.GOOS {
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			return "i-linux-x64", nil
		case "arm64":
			return "i-linux-arm64", nil
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			return "i-macos-intel-x64", nil
		case "arm64":
			return "i-macos-apple-silicon-arm64", nil
		}
	}
	return "", fmt.Errorf("unsupported platform: %s/%s", runtime.GOOS, runtime.GOARCH)
}

func getDownloadURL(apiURL, assetTarget string) (string, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "i-installer-go")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}

	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, assetTarget) {
			return asset.DownloadURL, nil
		}
	}

	return "", fmt.Errorf("asset '%s' not found in latest release", assetTarget)
}

func downloadFile(url string, dest *os.File) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(dest, resp.Body)
	return err
}

func installBinary(srcPath, destPath string) error {
	dir := filepath.Dir(destPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		if os.IsPermission(err) {
			if err := runSudo("mkdir", "-p", dir); err != nil {
				return fmt.Errorf("sudo mkdir failed: %w", err)
			}
		} else {
			return err
		}
	}

	err := copyFile(srcPath, destPath)
	if err == nil {
		return os.Chmod(destPath, 0755)
	}

	if os.IsPermission(err) {
		fmt.Println("Permission denied. Attempting installation with sudo...")
		if err := runSudo("cp", srcPath, destPath); err != nil {
			return fmt.Errorf("sudo cp failed: %w", err)
		}
		if err := runSudo("chmod", "755", destPath); err != nil {
			return fmt.Errorf("sudo chmod failed: %w", err)
		}
		return nil
	}

	return err
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func runSudo(args ...string) error {
	cmd := exec.Command("sudo", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
	os.Exit(1)
}
