#!/bin/sh
echo "---------------------"
echo "TEST : examples/example00"
go run . examples/example00.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : examples/example01"
go run . examples/example01.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : examples/example02"
go run . examples/example02.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : examples/example03"
go run . examples/example03.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : examples/example04"
go run . examples/example04.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : examples/example05"
go run . examples/example05.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : examples/example06"
go run . examples/example06.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "## RUNNING TEST WITH: examples/example07 ##"
go run . examples/example07.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : badexamples/example00"
go run . badexamples/example00.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : badexamples/example01"
go run . badexamples/example01.txt
echo "Press ENTER to run next test..."

read -n 1 -s




