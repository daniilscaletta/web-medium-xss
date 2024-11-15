package handlers

import (
	"bufio"
	"example/v3/bot/middleware"
	"fmt"
	"net"
)

func HandlerConnection(conn net.Conn) {

	defer conn.Close()

	welcomeMessage := "Welcome to the bot! Please type your message.\n"
	_, err := conn.Write([]byte(welcomeMessage))
	if err != nil {
		fmt.Println("Error sending welcomeMessage:", err)
	}

	// challenge := utils.GenerateHashCashChallenges()
	// challenge = fmt.Sprintf("Your challenge on HashCash: %s\n", challenge)
	// conn.Write([]byte(challenge))

	reader := bufio.NewReader(conn)
	// solution, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading message:", err)
	// 	return
	// }

	// if utils.CheckHashCashChallenge(challenge, strings.TrimSpace(solution)) {
	// 	conn.Write([]byte("Correct! You solved the challenge.\n"))
	conn.Write([]byte("Enter the path for appointment\n"))
	conn.Write([]byte("Example: /MDA1S0wS0wDA0OjU0Rvcmlub2xhcmluZ29b2c\n"))
	// } else {
	// 	conn.Write([]byte("Incorrect! Try again.\n"))
	// 	os.Exit(1)
	// }

	path, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading path:", err)
		return
	}

	middleware.SendRequest(path)

}
