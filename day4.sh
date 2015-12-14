#!/bin/dash
# Proper header for a Bash script.

# Advent of Code, 2015: Day 4
echo "Advent of Code, 2015: Day 4"
INPUTCODE=ckczppom

echo "Starting calculation..."
for COUNTER in $(seq 1 1000000)
  do
  # bash
  # sum=$(md5sum <<< $COUNTER$INPUTCODE | cut -d ' ' -f 1)
  # dash
  sum=$(echo $COUNTER$INPUTCODE | md5sum)

  cutsum=$(echo $sum | cut -c1-2)

  # Output debugging information
  # echo "Input: " $COUNTER$INPUTCODE " | "  "Santa MD5: " $sum " | "  "Cut Santa MD5: " $cutsum

  if [ $cutsum = "00" ]
  then
    echo "Lowest Value with two zeroes: " $COUNTER  "  MD5: " $sum
    echo "Lowest Value with two zeroes: " $COUNTER  "  MD5: " $sum > output.txt
    break
  fi

done


exit
