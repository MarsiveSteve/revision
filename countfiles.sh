#! /bin/bash
file_count=$(find . \( -type f -o -type d \) | wc -l)

echo $((file_count));