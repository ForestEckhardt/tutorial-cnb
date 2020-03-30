package node

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/cloudfoundry/packit"
)

func Build() packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		file, err := os.Open(filepath.Join(context.CNBPath, "buildpack.toml"))
		if err != nil {
			return packit.BuildResult{}, err
		}

		var m struct {
			Metadata struct {
				Dependencies []struct {
					URI string `toml:"uri"`
				} `toml:"dependencies"`
			} `toml:"metadata"`
		}
		_, err = toml.DecodeReader(file, &m)
		if err != nil {
			return packit.BuildResult{}, err
		}

		uri := m.Metadata.Dependencies[0].URI
		fmt.Printf("URI -> %s\n", uri)

		nodeLayer, err := context.Layers.Get("node", packit.LaunchLayer)
		if err != nil {
			return packit.BuildResult{}, err
		}

		err = nodeLayer.Reset()
		if err != nil {
			return packit.BuildResult{}, err
		}

		downloadDir, err := ioutil.TempDir("", "downloadDir")
		if err != nil {
			return packit.BuildResult{}, err
		}
		defer os.RemoveAll(downloadDir)

		fmt.Println("Downloading dependency...")
		err = exec.Command("curl",
			uri,
			"-o", filepath.Join(downloadDir, "node.tar.xz"),
		).Run()
		if err != nil {
			return packit.BuildResult{}, err
		}

		fmt.Println("Untaring dependency...")
		err = exec.Command("tar",
			"-xf",
			filepath.Join(downloadDir, "node.tar.xz"),
			"--strip-components=1",
			"-C", nodeLayer.Path,
		).Run()
		if err != nil {
			return packit.BuildResult{}, err
		}

		return packit.BuildResult{
			Plan: context.Plan,
			Layers: []packit.Layer{
				nodeLayer,
			},
		}, nil
	}
}
