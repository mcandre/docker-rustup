// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "target"

// Default references the default build task.
var Default = Test

// ImageNamespace places Docker images in an ownership hierarchy.
var ImageNamespace = "mcandre"

// Image labels the Docker artifact.
var Image = "rustup"

// X8664GnuTag distinguishes the 64-bit x86 GNU image variant.
var X8664GnuTag = "x86_64-gnu"
// I686GnuTag distinguishes the 32-bit x86 GNU image variant.
var I686GnuTag = "i686-gnu"
// X8664MuslTag distinguishes the 64-bit x86 Musl image variant.
var X8664MuslTag = "x86_64-musl"
// I686MuslTag distinguishes the 32-bit x86 Musl image variant.
var I686MuslTag = "i686-musl"

// ImageTags collects all available tags.
var ImageTags = []string{
	X8664GnuTag,
	I686GnuTag,
	X8664MuslTag,
	I686MuslTag,
}

// X8664Gnu generates a 64-bit x86 GNU Docker image.
func X8664Gnu() error {
	cmd := exec.Command(
		"docker",
		"build",
		"-t",
		fmt.Sprintf("%s/%s:%s", ImageNamespace, Image, X8664GnuTag),
		".",
	)
	cmd.Dir = X8664GnuTag
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// TestX8664Gnu performs a basic check of the 64-bit x86 GNU image.
func TestX8664Gnu() error {
	mg.Deps(X8664Gnu)

	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	cwdAbs, err := filepath.Abs(cwd)

	if err != nil {
		return err
	}

	cmd := exec.Command(
		"docker",
		"run",
		"-v",
		fmt.Sprintf("%s:/src", cwdAbs),
		fmt.Sprintf("%s/%s:%s", ImageNamespace, Image, X8664GnuTag),
		"sh",
		"-c",
		"cd /src && cargo build --release",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// X8664Musl generates a 64-bit x86 musl Docker image.
func X8664Musl() error {
	cmd := exec.Command(
		"docker",
		"build",
		"-t",
		fmt.Sprintf("%s/%s:%s", ImageNamespace, Image, X8664MuslTag),
		".",
	)
	cmd.Dir = X8664MuslTag
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// TestX8664Musl performs a basic check of the 64-bit x86 musl image.
func TestX8664Musl() error {
	mg.Deps(X8664Musl)

	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	cwdAbs, err := filepath.Abs(cwd)

	if err != nil {
		return err
	}

	cmd := exec.Command(
		"docker",
		"run",
		"-v",
		fmt.Sprintf("%s:/src", cwdAbs),
		fmt.Sprintf("%s/%s:%s", ImageNamespace, Image, X8664MuslTag),
		"sh",
		"-c",
		"cd /src && cargo build --release",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// I686Gnu generates a 32-bit x86 GNU Docker image.
func I686Gnu() error {
	cmd := exec.Command(
		"docker",
		"build",
		"-t",
		fmt.Sprintf("%s/%s:%s", ImageNamespace, Image, I686GnuTag),
		".",
	)
	cmd.Dir = I686GnuTag
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// TestI686Gnu performs a basic check of the 32-bit x86 GNU image.
func TestI686Gnu() error {
	mg.Deps(I686Gnu)

	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	cwdAbs, err := filepath.Abs(cwd)

	if err != nil {
		return err
	}

	cmd := exec.Command(
		"docker",
		"run",
		"-v",
		fmt.Sprintf("%s:/src", cwdAbs),
		fmt.Sprintf("%s/%s:%s", ImageNamespace, Image, I686GnuTag),
		"sh",
		"-c",
		"cd /src && cargo build --release",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// I686Musl generates a 32-bit x86 musl Docker image.
func I686Musl() error {
	cmd := exec.Command(
		"docker",
		"build",
		"-t",
		fmt.Sprintf("%s/%s:%s", ImageNamespace, Image, I686MuslTag),
		".",
	)
	cmd.Dir = I686MuslTag
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// TestI686Musl performs a basic check of the 32-bit x86 musl image.
func TestI686Musl() error {
	mg.Deps(I686Musl)

	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	cwdAbs, err := filepath.Abs(cwd)

	if err != nil {
		return err
	}

	cmd := exec.Command(
		"docker",
		"run",
		"-v",
		fmt.Sprintf("%s:/src", cwdAbs),
		fmt.Sprintf("%s/%s:%s", ImageNamespace, Image, I686MuslTag),
		"sh",
		"-c",
		"cd /src && cargo build --release",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// Test executes the integration test suite.
func Test() error {
	mg.Deps(TestX8664Gnu)
	mg.Deps(TestI686Gnu)
	mg.Deps(TestX8664Musl)
	mg.Deps(TestI686Musl)
	return nil
}

// NPMBin attempts to locate the directory for NPM binaries.
func NPMBin() (*string, error) {
	cmd := exec.Command("npm", "bin")
	cmd.Stderr = os.Stderr

	outBuf, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	out := strings.TrimSpace(string(outBuf))

	return &out, nil
}

// DockerfileBarLint runs projectatomic/dockerfile_lint.
func DockerfileBarLint() error {
	npmBinDir, err := NPMBin()

	if err != nil {
		return err
	}

	dockerfiles, err := filepath.Glob("*/Dockerfile")

	if err != nil {
		return err
	}

	for _, dockerfile := range(dockerfiles) {
		cmd := exec.Command(
			path.Join(*npmBinDir, "dockerfile_lint"),
			"-f",
			dockerfile,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// Hadolint runs hadolint.
func Hadolint() error {
	dockerfiles, err := filepath.Glob("*/Dockerfile")

	if err != nil {
		return err
	}

	for _, dockerfile := range(dockerfiles) {
		cmd := exec.Command(
			"hadolint",
			dockerfile,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// DockerfileLint runs replicatedhq/dockerfilelint.
func DockerfileLint() error {
	npmBinDir, err := NPMBin()

	if err != nil {
		return err
	}

	dockerfiles, err := filepath.Glob("*/Dockerfile")

	if err != nil {
		return err
	}

	for _, dockerfile := range(dockerfiles) {
		cmd := exec.Command(
			path.Join(*npmBinDir, "dockerfilelint"),
			dockerfile,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// DockerLint runs RedCoolBeans/dockerlint.
func DockerLint() error {
	npmBinDir, err := NPMBin()

	if err != nil {
		return err
	}

	dockerfiles, err := filepath.Glob("*/Dockerfile")

	if err != nil {
		return err
	}

	for _, dockerfile := range(dockerfiles) {
		cmd := exec.Command(
			path.Join(*npmBinDir, "dockerlint"),
			"-f",
			dockerfile,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return nil
		}
	}

	return nil
}

// DockerfileUtils runs rcjsuen/dockerfile-utils.
func DockerfileUtils() error {
	npmBinDir, err := NPMBin()

	if err != nil {
		return err
	}

	dockerfiles, err := filepath.Glob("*/Dockerfile")

	if err != nil {
		return err
	}

	for _, dockerfile := range(dockerfiles) {
		cmd := exec.Command(
			path.Join(*npmBinDir, "dockerfile-utils"),
			"lint",
			dockerfile,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(DockerfileBarLint)
	mg.Deps(Hadolint)
	mg.Deps(DockerfileLint)
	mg.Deps(DockerLint)
	mg.Deps(DockerfileUtils)
	return nil
}

// Publish uploads images to Docker repository.
func Publish() error {
	for _, imageTag := range(ImageTags) {
		cmd := exec.Command(
			"docker",
			"push",
			fmt.Sprintf("%s/%s:%s", ImageNamespace, Image, imageTag),
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// Clean deletes any leftover build artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
