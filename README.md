## Binary file "avg"

2 ways to use it:
- -f command should create a new csv file with average appended in case everything is correct.<br>
```console
foo@bar:~$ ./avg -f scores.csv
//avg_scores.csv
```

- -s command should return the average for the input<br>
```console
foo@bar:~$ ./avg -s 10 8.5 10 9 10
//9.5
```
