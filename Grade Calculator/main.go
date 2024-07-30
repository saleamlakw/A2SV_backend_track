package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//calculate average
func calcualteAverage(result map[string]float64,numberOfSubjects uint) float64{ 
	var total float64
	for _,grade :=range result{
		total+=grade
	}
	return total/float64(numberOfSubjects)
}

//get grade 
func  getGrade(value float64)string{
	switch {
	case value >= 90:
		return "A+"
	case value >= 85:
		return "A"
	case value >= 80:
		return "B+"
	case value >= 75:
		return "B-"
	case value >= 70:
		return "B"
	case value >= 60:
		return "C"
	case value >= 50:
		return "D"
	default:
		return "F"
	}
}
//accept and validate number of subjects
func acceptNoOfSubjects()uint{
	reader:=bufio.NewReader(os.Stdin)
	for{
		fmt.Println("Enter the number of subjects you are taking :")
		value,_:=reader.ReadString('\n')
		trimed:=strings.TrimSpace(value)

		numberOfSubjects,err:=strconv.ParseUint(trimed,10,2)

		if err!=nil || numberOfSubjects<=0{
			fmt.Println("Please enter the correct number of subjects you have taken. The number must be a positive integer.")
			continue
		}
		return uint(numberOfSubjects)
		}
	}

//accept and validate grade
func acceptGrade() float64{
	reader:=bufio.NewReader(os.Stdin)
	for{
		fmt.Println("Enter the grade you got :")
		value,_:=reader.ReadString('\n')
		trimed:=strings.TrimSpace(value)
		grade,err:=strconv.ParseFloat(trimed,64)
		if err!=nil || grade<0 || grade>100{
			fmt.Println("Please enter a valid grade between 1 and 100.")
			continue
		}
		return float64(grade)
}

}
func main(){
	var firstName string
	var lastName string

	//accept user name and number of subjects
	fmt.Println("Enter your First Name :")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your Last Name :")
	fmt.Scanln(&lastName)

	
	numberOfSubjects:=acceptNoOfSubjects()
	

	// accept all subjects name and grade 
	var subjectName string
	var grade float64


	result:=make(map [string]float64)
	for i :=0;i<int(numberOfSubjects);i++{
		fmt.Printf("Enter the name of subject %v :\n",i+1)
		fmt.Scanln(&subjectName)
		//accept and validate grade 
		grade=acceptGrade()
		result[subjectName]=grade

	}
	average:=calcualteAverage(result,numberOfSubjects)
	// header
    header := fmt.Sprintf("%-20s\n","Grade Report")
    line:=strings.Repeat("_", 20)
    fmt.Println(line)
    fmt.Println(header)
	fmt.Println(line)
    // Print subjects and grades
	fmt.Println(strings.Repeat("_", 20))
	fmt.Printf("|%-10s | %5s |\n","Subject","Grade")
	fmt.Println(strings.Repeat("_", 20))
	for subject,grade := range result{
		fmt.Printf("|%-10s | %5s |\n",subject,getGrade(grade))
	}
	
	//print avearge
	fmt.Println(strings.Repeat("_", 20))
	fmt.Printf("|%-10s : %5.2f |\n","Average",average)
	fmt.Println(strings.Repeat("_", 20))
}