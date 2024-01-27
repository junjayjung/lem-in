#!/bin/sh

run_test() {
    echo "---------------------"
    echo "Testing with $1"
    echo -e "\n"
    go run . $1
    echo "---------------------"
    echo -e "\n"
    echo "Press ENTER to proceed;"
    echo "Or press CTRL+C to exit."
    read -n 1 -s
}

run_test "examples/example00.txt"
run_test "examples/example01.txt"
run_test "examples/example02.txt"
run_test "examples/example03.txt"
run_test "examples/example04.txt"
run_test "examples/example05.txt"
run_test "badexamples/example00.txt"
run_test "badexamples/example01.txt"
run_test "examples/example06.txt"
run_test "examples/example07.txt"