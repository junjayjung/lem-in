#!/bin/sh
echo "---------------------"
echo "TEST : example00"
go run . example00.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : example01"
go run . example01.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : example02"
go run . example02.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : example03"
go run . example03.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : example04"
go run . example04.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : example05"
go run . example05.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : example06"
go run . example06.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "## RUNNING TEST WITH: example07 ##"
go run . example07.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : badexample00"
go run . badexample00.txt
echo "Press ENTER to run next test..."

read -n 1 -s

echo "---------------------"
echo "TEST : badexample01"
go run . badexample01.txt
echo "Press ENTER to run next test..."

read -n 1 -s




