package main

import (
	"fmt"
    "strings"
	"a21hc3NpZ25tZW50/helper"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)

func Login(id string, name string) string {

	if id == "" || name == ""{
		return ("ID or Name is undefined!")
	} 
	
	if len(id) != 5 {
		return ("ID must be 5 characters long!")
	}

	
		for _, studentList := range Students {
			parts := strings.Split(studentList, "_")
			if len(parts) == 3 && parts[0] == id && parts[1] == name {
				return fmt.Sprintf("Login berhasil: %s", parts[1],)
			}
		}
	return "Login gagal: data mahasiswa tidak ditemukan"// TODO: replace this
}


func Register(id string, name string, major string) string {

	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	} 

	if len(id) != 5 {
		return ("ID must be 5 characters long!")
	}

	for _, studentList := range Students {
		parts := strings.Split(studentList, "_")
		if len(parts) == 3 && parts[0] == id {
			return ("Registrasi gagal: id sudah digunakan")
		}
	}

	Students = append(Students, fmt.Sprintf("%s_%s_%s", id, name, major))
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major) // TODO: replace this
}

func GetStudyProgram(code string) string {
	code = strings.TrimSpace(code)

	if code == "" {
		return("Code is undefined")
	}
	
	if program, ok := StudentStudyPrograms[code]; ok {
		return program
	}
	return "Kode program studi tidak ditemukan" // TODO: replace this
}

func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	for i, studentList := range Students {
		parts := strings.Split(studentList, "_")
		if len(parts) == 3 && parts[1] == nama {
			fn(programStudi, &Students[i])
			return "Program studi mahasiswa berhasil diubah."
		}
	}
	return "Mahasiswa tidak ditemukan." // TODO: replace this
}


func UpdateStudyProgram(programStudi string, students *string) {
	parts := strings.Split(*students, "_")
	if len(parts) == 3 {
		parts[2] = programStudi
		*students = strings.Join(parts, "_")
	}
	// TODO: answer here
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		for i, student := range Students {
			fmt.Println(i+1, student)
		}

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Change student study program")
		fmt.Println("5. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			var nama, programStudi string
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan program studi baru: ")
			fmt.Scanln(&programStudi)

			fmt.Println(ModifyStudent(programStudi, nama, UpdateStudyProgram))
			helper.Delay(5)
		case "5":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
