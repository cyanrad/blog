#!/bin/sh

for file in ./content/markdown/*.md; do
    # Construct the input file path
    input_file="$file"
    
    # Construct the output file path by replacing the directory and extension
    base_name=$(basename "$file" .md)
    output_file="content/posts/${base_name}.html"
    
    # Run pandoc to convert the markdown file to HTML
    ./deps/pandoc -i "$input_file" -o "$output_file"
done