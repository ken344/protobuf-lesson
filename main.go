package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"protobuf-lesson/pb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Projects:    map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		BirthDate: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	//// Serializeを使うと、Employeeをバイト列に変換できる
	//binData, err := proto.Marshal(employee)
	//if err != nil {
	//	log.Fatal("Can't serialize", err)
	//}
	//
	//if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
	//	log.Fatal("Can't write to file", err)
	//}
	//
	//in, err := ioutil.ReadFile("test.bin")
	//if err != nil {
	//	log.Fatalln("Can't read from file", err)
	//}
	//
	//readEmployee := &pb.Employee{}
	//
	//// Unmarshalを使うと、バイト列をEmployeeに変換できる
	//err = proto.Unmarshal(in, readEmployee)
	//if err != nil {
	//	log.Fatalln("Failed to parse employee", err)
	//}
	//
	//fmt.Println(readEmployee)

	// Marshalerを使うと、EmployeeをJSONに変換できる
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Failed to marshal", err)
	}

	// Unmarshalerを使うと、JSONをEmployeeに変換できる
	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Failed to unmarshal", err)
	}

	fmt.Println(readEmployee)

}
