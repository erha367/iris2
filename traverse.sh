#!/bin/bash
# 递归处理文件夹中的图片
function traverse_directory {
    local dir="$1"
    local file
    for file in "$dir"/*; do
        if [[ -f "$file" ]]; then
             iris -color=blue -url="$file"
             echo "$file"
        elif [[ -d "$file" ]]; then
            traverse_directory "$file"
        fi
    done
}

traverse_directory "/Volumes/Untitled/elitebabes"
