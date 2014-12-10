#!/bin/sh

# make the code aways recompile
echo "" >> mylib.go
head -$((`cat mylib.go | wc -l`-1)) mylib.go > $$ && mv $$ mylib.go

# record the install program
go install -work -x 2> step2.sh

# fix CGO_LDFLAGS format
sed -i -e 's/"-g" "-O2"/"-g -O2"/g' step2.sh

# fix go tool command
sed -i -e 's/pack r/go tool pack r/g' step2.sh

# the hack
head -5 step2.sh > $$ && \
echo 'sed -i -e "s/runtime·cgocall/runtime·asmcgocall/g" `find $WORK/labs22/mylib/_obj/ -name *.c`' >> $$ && \
tail -n$((`cat step2.sh | wc -l`-5)) step2.sh >> $$ && \
mv $$ step2.sh

sh step2.sh
rm step2.*