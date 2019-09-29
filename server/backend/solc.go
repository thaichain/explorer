package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gochain-io/explorer/server/utils"
)

var versionRegexp = regexp.MustCompile(`([0-9]+)\.([0-9]+)\.([0-9]+)`)

// Contract contains information about a compiled contract, alongside its code and runtime code.
type Contract struct {
	Code        string       `json:"code"`
	RuntimeCode string       `json:"runtime-code"`
	Info        ContractInfo `json:"info"`
}

// ContractInfo contains information about a compiled contract, including access
// to the ABI definition, source mapping, user and developer docs, and metadata.
//
// Depending on the source, language version, compiler version, and compiler
// options will provide information about how the contract was compiled.
type ContractInfo struct {
	Source          string          `json:"source"`
	Language        string          `json:"language"`
	LanguageVersion string          `json:"languageVersion"`
	CompilerVersion string          `json:"compilerVersion"`
	CompilerOptions string          `json:"compilerOptions"`
	SrcMap          string          `json:"srcMap"`
	SrcMapRuntime   string          `json:"srcMapRuntime"`
	AbiDefinition   []utils.AbiItem `json:"abiDefinition"`
	UserDoc         interface{}     `json:"userDoc"`
	DeveloperDoc    interface{}     `json:"developerDoc"`
	Metadata        string          `json:"metadata"`
}

// Solidity contains information about the solidity compiler.
type Solidity struct {
	Path, Version string
	Optimization  bool
}

// --combined-output format
type solcOutput struct {
	Contracts map[string]struct {
		BinRuntime                                  string `json:"bin-runtime"`
		SrcMapRuntime                               string `json:"srcmap-runtime"`
		Bin, SrcMap, Abi, Devdoc, Userdoc, Metadata string
	}
	Version string
}

func (s *Solidity) makeArgs() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	args := []string{
		"run", "-i", "-v", "/home" + dir + ":/workdir", "-w", "/workdir", "ethereum/solc:" + s.Version,
		"--combined-json",
		"bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata",
	}
	if s.Optimization {
		args = append(args, "--optimize")
	}
	return args, nil
}

// CompileSolidityString builds and returns all the contracts contained within a source string.
func CompileSolidityString(ctx context.Context, compilerVersion, source string, optimization bool) (map[string]*Contract, error) {
	if len(source) == 0 {
		return nil, errors.New("solc: empty source string")
	}
	s := &Solidity{Path: "docker", Version: compilerVersion, Optimization: optimization}
	args, err := s.makeArgs()
	if err != nil {
		return nil, err
	}
	args = append(args, "--")
	cmd := exec.CommandContext(ctx, s.Path, append(args, "-")...)
	cmd.Stdin = strings.NewReader(source)
	return s.run(cmd, source)
}

func (s *Solidity) run(cmd *exec.Cmd, source string) (map[string]*Contract, error) {
	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("solc: %v\n%s", err, stderr.Bytes())
	}
	args, err := s.makeArgs()
	if err != nil {
		return nil, err
	}
	return ParseCombinedJSON(stdout.Bytes(), source, s.Version, s.Version, strings.Join(args, " "))
}

func ParseCombinedJSON(combinedJSON []byte, source string, languageVersion string, compilerVersion string, compilerOptions string) (map[string]*Contract, error) {
	var output solcOutput
	if err := json.Unmarshal(combinedJSON, &output); err != nil {
		return nil, err
	}

	// Compilation succeeded, assemble and return the contracts.
	contracts := make(map[string]*Contract)
	for name, info := range output.Contracts {
		// Parse the individual compilation results.
		var abi []utils.AbiItem
		if err := json.Unmarshal([]byte(info.Abi), &abi); err != nil {
			return nil, fmt.Errorf("solc: error reading abi definition (%v)", err)
		}
		var userdoc interface{}
		if err := json.Unmarshal([]byte(info.Userdoc), &userdoc); err != nil {
			return nil, fmt.Errorf("solc: error reading user doc: %v", err)
		}
		var devdoc interface{}
		if err := json.Unmarshal([]byte(info.Devdoc), &devdoc); err != nil {
			return nil, fmt.Errorf("solc: error reading dev doc: %v", err)
		}
		contracts[name] = &Contract{
			Code:        "0x" + info.Bin,
			RuntimeCode: "0x" + info.BinRuntime,
			Info: ContractInfo{
				Source:          source,
				Language:        "Solidity",
				LanguageVersion: languageVersion,
				CompilerVersion: compilerVersion,
				CompilerOptions: compilerOptions,
				SrcMap:          info.SrcMap,
				SrcMapRuntime:   info.SrcMapRuntime,
				AbiDefinition:   abi,
				UserDoc:         userdoc,
				DeveloperDoc:    devdoc,
				Metadata:        info.Metadata,
			},
		}
	}
	return contracts, nil
}
