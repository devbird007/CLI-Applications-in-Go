package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"fmt"
	"github.com/devbird007/pScan/"
)

func setup(t *testing.T, hosts []string, initList