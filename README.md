# goworld
This repository provides means to access C language interface from Go Programming Language.
It's ideal to have plain go language but scenario where some large legacy code porting can be huge task. 
In such scenario wrapping legacy c or c++ code under the hood of a go language is an optimal solution.

BUILD goworld

git clone https://github.com/mohankri/goworld

cd go/cppwrap/cppsrc
make
cd ..
make

The above will build dynamic C libaries

cd ../../go/src/main
export GOPATH=$HOME/goworld/go
go build

export LD_LIBRARY_PATH=$GOPATH/cppwrap

run the program
./main 
