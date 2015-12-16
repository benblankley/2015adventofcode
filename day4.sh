#!/bin/dash
# Proper header for a Bash script.

# Advent of Code, 2015: Day 4
echo "Advent of Code, 2015: Day 4"
INPUTCODE=ckczppom

echo "Starting calculation..."
for COUNTER in $(seq 1 10000000)
  do
  # bash
  sum=$(md5sum - <<< "$COUNTER$INPUTCODE" | tr -d ' -')
  # dash
  # sum=$(echo -n $COUNTER$INPUTCODE | md5sum -)
  cutsum=$(echo -n "$sum" | cut -c1-5)
  cutsum2=$(echo -n "$sum" | cut -c1-6)

  # Output debugging information
  echo "Input: " $COUNTER$INPUTCODE " | "  "Santa MD5: " $sum " | "  "Cut Santa MD5: " $cutsum

  # Output minimal debug information
  if [ $((COUNTER % 1000)) -eq 0 ]
  then
    echo "Step: " "$COUNTER"
    echo "Step: " "$COUNTER" >> status.txt
  fi


  if [ "$cutsum"="00000" ]
  then
    echo "Lowest Value with five zeroes: " "$COUNTER"  "  MD5: " "$sum"
    echo "Lowest Value with five zeroes: " "$COUNTER"  "  MD5: " "$sum" >> output.txt
    break
  fi

  if [ "$cutsum2"="000000" ]
  then
    echo "Lowest Value with six zeroes: " "$COUNTER"  "  MD5: " "$sum"
    echo "Lowest Value with six zeroes: " "$COUNTER"  "  MD5: " "$sum" >> output.txt
    break
  fi

done


exit
