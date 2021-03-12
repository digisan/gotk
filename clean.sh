#!/bin/bash

rm -rf ./slice/ti/auto.go 
rm -rf ./slice/ti32/auto.go 
rm -rf ./slice/tu8/auto.go 
rm -rf ./slice/tis/auto.go 
rm -rf ./slice/to/auto.go 
rm -rf ./slice/ts/auto.go 
rm -rf ./slice/tsi/auto.go

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f
for f in $(find ./ -name '*.log' -or -name '*.txt'); do rm $f; done
find ./ -type d -empty -delete