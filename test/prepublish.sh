if [ $(gofmt -l . | wc -l) -eq 0 ]
then
    exit 0
else
    exit 1
fi
