package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func pwd() {
	// pwd выводит путь до текущей директории
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("ошибка в команде pwd: ", err)
		os.Exit(1)
	}
	fmt.Println(pwd)
}

func cd(args []string) {
	// cd Выполняет переход в указанную директорию
	if len(args) < 2 {
		log.Println("путь не указан")
		return
	}
	// Chdir изменяет текущий рабочий каталог на каталог с именем. Если возникает ошибка, то она будет типа *PathError
	err := os.Chdir(args[1])
	if err != nil {
		log.Println("ошибка в команде cd: ", err)
	}
}

func echo(args []string) {
	// echo используется для печати введенного текста в терминале
	fmt.Println(strings.Join(args[1:], " "))
}

func kill(args []string) {
	// kill завершает процесс по его PID
	if len(args) < 2 {
		fmt.Println("kill: PID не указан")
		return
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		log.Println("ошибка в команде kill: ", err)
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Println("ошибка в команде kill: ", err)
		return
	}
	err = process.Kill()
	if err != nil {
		log.Println("ошибка в команде kill: ", err)
		return
	}
}

func psCommand(args []string) {
	// ps ввыводит список запущенных процессов
	if len(args) != 1 {
		log.Println("ps не ожидает дополнительных аргументов")
		return
	}

	proceslist, err := ps.Processes()
	if err != nil {
		log.Println("Ошибка в поиске запущенных процессов: ", err)
		return
	}
	for _, val := range proceslist {
		fmt.Println("PID процесса: ", val.Pid(), " Название процесса: ", val.Executable())

	}
}

func execCommand(args []string) {
	// Выполнение внешней команды, основная горутина ожидает выполнения
	// exec.Command создает новый объект команды. args[0] содержит имя команды, которую нужно выполнить (например ls).
	cmd := exec.Command(args[0], args[1:]...)
	// Указывает, что стандартный вывод команды должен быть направлен на стандартный вывод текущего процесса (т.е. на консоль).
	cmd.Stdout = os.Stdout
	// Указывает, что стандартный вывод ошибок команды должен быть направлен на стандартный вывод ошибок текущего процесса.
	cmd.Stderr = os.Stderr
	// Указывает, что стандартный ввод команды должен быть направлен на стандартный ввод текущего процесса.
	cmd.Stdin = os.Stdin
	// cmd.Run() выполняет команду и ждет её завершения.
	err := cmd.Run()
	if err != nil {
		log.Println("Ошибка при выполнении команды: ", err)
	}
}

func forkExecCommand(args []string) {
	// Создание нового процесса и выполнение команды, основная горутина не ждет выполнения.
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Start() // Асинхронный запуск команды
	if err != nil {
		log.Println("Ошибка при выполнении команды: ", err)
		return
	}

	// Ожидание завершения команды в отдельной горутине
	go func() {
		err := cmd.Wait()
		if err != nil {
			log.Println("Ошибка при завершении команды: ", err)
		}
	}()
}

func handleCommand(command string) {
	// В зависимости от введенной команды вызывает функцию ее выполнения.
	// Разбиваем прочитанную строку в слайс по пробелам
	args := strings.Fields(command)

	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		cd(args)
	case "pwd":
		pwd()
	case "echo":
		echo(args)
	case "kill":
		kill(args)
	case "ps":
		psCommand(args)
	case "fork-exec":
		forkExecCommand(args)
	default:
		execCommand(args)
	}
}

func main() {
	// Читаем данные с консоли
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите комманду > ")
		command, err := reader.ReadString('\n')
		// err != nil тогда и только тогда, когда возвращаемые данные не заканчиваются на разделитель
		if err != nil {
			// Если ошибка во время чтения, то программа завершается с кодом 1
			if err == io.EOF {
				os.Exit(1)
			}
			log.Println(err)
		}
		command = strings.TrimSpace(command)
		handleCommand(command)
	}
}
