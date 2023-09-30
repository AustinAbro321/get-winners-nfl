## Pre-reqs

In order to use this you will need to have go install. Follow this guide to install go https://go.dev/doc/install

git clone or download and unzip the code

Running the script  - 
Once you have go installed open a terminal. Make sure you have get_winners.go downloaded and in your current directory. type the below command in your terminal. 

```
go run get_winnners.go (path_to_excel_file)
```

Example: 
```
go run get_winners.go /home/austin/Downloads/test.xlsx
```

A new excel file called calculatedWins.xlsx will be in the directory you called the script from

## Assumptions 

This script makes a few assumptions
- The first row in the third excel column needs to have the number for the week in it. Example "Week 12"
- Assumes it's the current football season. However if you want to use historical data you can add an additional arguement year when calling the script. Example
```
go run get_winnners.go /home/austin/Downloads/test.xlsx 2022
```
- The third excel column, week 3, should have one of the following as the values - case doesn't matter
  - falcons
  - bills
  - bears
  - bengals
  - browns
  - cowboys
  - broncos
  - lions
  - packers
  - titans
  - colts
  - chiefs
  - raiders
  - rams
  - dolphins
  - vikings
  - patriots
  - saints
  - giants
  - jets
  - eagles
  - cardinals
  - steelers
  - chargers
  - 49ers
  - seahawks
  - buccaneers
  - commanders
  - panthers
  - jaguars
  - ravnes
  - texans