package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const cfgStack = "cfg/base.yaml,cfg/template_breadcrumb.yaml,cfg/rmpp.base.yaml,cfg/rmpp.breadcrumb.lined.default.ampm.dailycal.reflectextra.yaml"

func main() {
	nowYear := time.Now().Year()
	year := flag.Int("year", nowYear, "planner year to build")
	name := flag.String("name", "Planner", "name for the copied pdf (without extension)")
	output := flag.String("output", "compiled", "directory to copy the finished pdf into")
	flag.Parse()

	if *name == "" {
		fmt.Fprintln(os.Stderr, "name cannot be empty")
		os.Exit(1)
	}

	if err := os.MkdirAll(*output, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "create output dir: %v\n", err)
		os.Exit(1)
	}

	tempBase := "rmpp_app_build"
	cmd := exec.Command("./single.sh")
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("PLANNER_YEAR=%d", *year),
		"PASSES=2",
		fmt.Sprintf("CFG=%s", cfgStack),
		fmt.Sprintf("NAME=%s", tempBase),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "build planner: %v\n", err)
		os.Exit(1)
	}

	src := tempBase + ".pdf"
	if _, err := os.Stat(src); err != nil {
		fmt.Fprintf(os.Stderr, "expected %s: %v\n", src, err)
		os.Exit(1)
	}

	dest := filepath.Join(*output, *name+".pdf")

	if err := copyFile(src, dest); err != nil {
		fmt.Fprintf(os.Stderr, "copy file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Planner saved to %s\n", dest)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	return out.Close()
}
