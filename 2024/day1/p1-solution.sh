
abs_diff () {
    lnum=$1
    rnum=$2
    echo $(($1-$2)) | sed 's/-//'
}

consume () {
    # consume two fifos, print their contents to stdout in two columns
    llfifo=$1
    rlfifo=$2
    ofifo=$3
    while IFS= read -u 3 -r li && read -u 4 -r ri
    do
        printf "%s\n" $(abs_diff $li $ri)
    done 3<"$llfifo" 4<"$rlfifo"
}


if [ ! -f /tmp/aoc24-day1-input.txt ]; then
    echo "ERR -- get the input by tangling the org-mode file!"
fi
if [ ! -f llfifo ]; then
    mkfifo llfifo
fi
if [ ! -f rlfifo ]; then
    mkfifo rlfifo
fi
if [ ! -f ofifo ]; then
    mkfifo ofifo
fi
# read left col
awk '{print $1}' /tmp/aoc24-day1-input.txt | sort -n | tee llfifo &
# read right col
awk '{print $2}' /tmp/aoc24-day1-input.txt | sort -n | tee rlfifo &

# consume both fifos into third
consume ./llfifo ./rlfifo | tee ofifo &
# sum the results in ofifo
awk '{sum+=$1} END{print sum}' ofifo

unlink llfifo 
unlink rlfifo 
unlink ofifo
