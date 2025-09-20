#!/usr/bin/env bash

set -euo pipefail

GREEN="\033[32m"
RESET="\033[0m"

files=()

snake_to_title() {
  echo "$1" | sed 's/_/ /g' | sed 's/\b\w/\U&/g'
}

snake_to_pascal() {
  echo "$1" | sed -r 's/(^|_)([a-z])/\U\2/g'
}

print_help() {
  cat <<EOF
Usage: $0 -d <directory> -f <file1,file2,...> -e <extension> [more -d -f -e sets]

Flags:
  -d	Directory name to create
  -f	Comma-separated list of files (without extension)
  -e	File extension (e.g. go, md, txt)
  -h	Show this help message

Examples:
  $0 	-d sorting -f bubble_sort,quick_sort -e go
  $0	-d searching -f binary_search,linear_search -e go \\
			-d utils -f helpers -e go
EOF
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    -d)
      dir=$2
      shift 2
      ;;
    -f)
      IFS=',' read -ra files <<< "$2"
      shift 2
      ;;
    -e)
      ext=$2
      shift 2
      ;;
    -h|--help)
      print_help
      exit 0
      ;;
    *)
      echo "Unknown option: $1"
      print_help
      exit 1
      ;;
  esac
  if [[ -n "${dir:-}" && -n "${ext:-}" && -n "${files+x}" && ${#files[@]} -gt 0 ]]; then
    mkdir -p "$dir"
    
    dir_title=$(snake_to_title "$dir")
    cat > "$dir/README.md" << EOF
# $dir_title Algorithms

This directory contains implementations of $dir_title algorithms.

EOF

    for file in "${files[@]}"; do
      title_case=$(snake_to_title "$file")
      cat >> "$dir/README.md" << EOF
## $title_case

### Properties

- Time Complexity:
- Space Complexity:

### Use-Cases

### How it works

### Flow

\`\`\`mermaid
flowchart TD
\`\`\`

EOF
    done

    for file in "${files[@]}"; do
      if [[ "$ext" == "go" ]]; then
        pascal_case=$(snake_to_pascal "$file")
        
        cat > "$dir/$file.$ext" << EOF
package $dir

// $pascal_case implements the $file algorithm
func $pascal_case() {
	// TODO: Implement $file algorithm
}
EOF

        cat > "$dir/${file}_test.$ext" << EOF
package $dir

import (
	"testing"
)

// Test$pascal_case tests the $pascal_case function
func Test$pascal_case(t *testing.T) {
	// TODO: Add test cases for $file algorithm
}
EOF
      else
        touch "$dir/$file.$ext"
      fi
    done
    echo -e "${GREEN}[+]${RESET} Collection '$dir' created with ${#files[@]} files and README.md."
    unset dir ext
    files=()
  fi
done
