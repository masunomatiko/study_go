ch=$CH
end=$END
for i in $(seq -w 1 $end); do
    mkdir -p "ch$ch/ex$i"
    touch "ch$ch/ex$i/ex$ch$i.go"
done
