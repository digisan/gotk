#!/bin/bash

rm -rf ./slice/int/auto.go ./slice/intstr/auto.go ./slice/obj/auto.go ./slice/str/auto.go ./slice/strint/auto.go 

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f
for f in $(find ./ -name '*.log' -or -name '*.txt'); do rm $f; done
find ./ -type d -empty -delete